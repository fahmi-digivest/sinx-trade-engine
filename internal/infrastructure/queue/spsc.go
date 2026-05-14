// Package queue menyediakan implementasi queue yang thread-safe untuk berbagai
// pola penggunaan concurrent.
//
// Contoh penggunaan dasar SPSCQueue:
//
//	q := queue.NewSPSC[string](16)
//	defer q.Close()
//
//	// Producer goroutine
//	go func() {
//	    for _, msg := range messages {
//	        if err := q.Enqueue(msg); err != nil {
//	            return // queue sudah ditutup
//	        }
//	    }
//	    q.Close()
//	}()
//
//	// Consumer goroutine
//	for {
//	    item, err := q.Dequeue()
//	    if errors.Is(err, queue.ErrQueueClosed) {
//	        break // selesai
//	    }
//	    process(item)
//	}
package queue

import (
	"errors"
	"sync"

	port "github.com/fahmi-digivest/sinx-trade-engine/internal/domain/port"
)

// ErrQueueClosed dikembalikan oleh Enqueue dan Dequeue ketika queue sudah
// ditutup via Close(). Gunakan errors.Is() untuk memeriksanya.
var ErrQueueClosed = errors.New("queue closed")

var _ port.SPSCQueue[any] = (*SPSCQueue[any])(nil)

// SPSCQueue adalah unbounded single-producer/single-consumer queue berbasis
// growable ring buffer. Data disimpan dalam slice agar locality memori tetap
// baik, lalu buffer diperbesar dua kali lipat saat penuh.
//
// Penggunaan yang benar:
//   - Hanya satu goroutine boleh memanggil Enqueue (producer).
//   - Hanya satu goroutine boleh memanggil Dequeue / TryDequeue (consumer).
//   - Close() aman dipanggil dari goroutine mana pun, termasuk lebih dari sekali.
//
// Contoh pola producer-consumer:
//
//	q := queue.NewSPSC[int](32)
//
//	// Producer
//	go func() {
//	    defer q.Close()
//	    for i := 0; i < 100; i++ {
//	        if err := q.Enqueue(i); err != nil {
//	            return
//	        }
//	    }
//	}()
//
//	// Consumer
//	for {
//	    v, err := q.Dequeue()
//	    if errors.Is(err, queue.ErrQueueClosed) {
//	        break
//	    }
//	    fmt.Println(v)
//	}
type SPSCQueue[T any] struct {
	mu     sync.Mutex
	buffer []T
	head   int
	count  int
	closed bool

	notifyCh  chan struct{}
	closeCh   chan struct{}
	closeOnce sync.Once
}

// NewSPSC membuat queue dengan kapasitas awal. Queue akan tumbuh otomatis saat
// penuh, sehingga Enqueue tidak pernah blocking karena kapasitas.
// Minimum kapasitas awal adalah 1.
//
// Pilih capacity yang mendekati throughput rata-rata untuk menghindari
// realokasi berulang; kapasitas yang terlalu kecil menyebabkan grow() sering
// dipanggil.
func NewSPSC[T any](capacity int) *SPSCQueue[T] {
	if capacity < 1 {
		capacity = 1
	}

	return &SPSCQueue[T]{
		buffer:   make([]T, capacity),
		notifyCh: make(chan struct{}, 1),
		closeCh:  make(chan struct{}),
	}
}

// Enqueue menambahkan item ke ujung queue tanpa blocking.
// Buffer akan diperbesar otomatis jika penuh.
//
// Mengembalikan [ErrQueueClosed] jika queue sudah ditutup; item tidak masuk.
// Aman dipanggil hanya dari satu goroutine (producer).
func (q *SPSCQueue[T]) Enqueue(item T) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return ErrQueueClosed
	}

	if q.count == len(q.buffer) {
		q.grow()
	}

	tail := (q.head + q.count) % len(q.buffer)
	q.buffer[tail] = item
	q.count++
	q.notifyConsumer()
	return nil
}

// Dequeue mengambil satu item dari depan queue. Blocking sampai ada item
// tersedia atau queue ditutup.
//
// Mengembalikan [ErrQueueClosed] hanya setelah queue ditutup DAN kosong;
// item yang sudah ada sebelum Close() tetap bisa diambil.
//
// Aman dipanggil hanya dari satu goroutine (consumer).
//
// Untuk non-blocking, gunakan [TryDequeue].
func (q *SPSCQueue[T]) Dequeue() (T, error) {
	for {
		q.mu.Lock()
		if q.count > 0 {
			item := q.dequeueItem()
			q.mu.Unlock()
			return item, nil
		}
		if q.closed {
			q.mu.Unlock()
			var zero T
			return zero, ErrQueueClosed
		}
		q.mu.Unlock()

		select {
		case <-q.notifyCh:
		case <-q.closeCh:
			// Drain sisa item yang di-enqueue sebelum Close dipanggil.
			q.mu.Lock()
			if q.count > 0 {
				item := q.dequeueItem()
				q.mu.Unlock()
				return item, nil
			}
			q.mu.Unlock()
			var zero T
			return zero, ErrQueueClosed
		}
	}
}

// TryDequeue mengambil satu item tanpa blocking.
// Mengembalikan (item, true) jika ada item, atau (zero, false) jika kosong.
//
// Berguna untuk polling atau saat consumer tidak boleh blocking:
//
//	if item, ok := q.TryDequeue(); ok {
//	    process(item)
//	}
func (q *SPSCQueue[T]) TryDequeue() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.count == 0 {
		var zero T
		return zero, false
	}

	return q.dequeueItem(), true
}

// Close menutup queue dan membangunkan consumer yang sedang menunggu.
// Setelah Close, Enqueue akan mengembalikan [ErrQueueClosed].
// Item yang sudah ada di queue sebelum Close tetap bisa di-Dequeue.
//
// Aman dipanggil lebih dari sekali dan dari goroutine mana pun.
func (q *SPSCQueue[T]) Close() {
	q.mu.Lock()
	if q.closed {
		q.mu.Unlock()
		return
	}
	q.closed = true
	q.mu.Unlock()

	q.closeOnce.Do(func() {
		close(q.closeCh)
	})
}

// IsClosed melaporkan apakah queue sudah ditutup.
func (q *SPSCQueue[T]) IsClosed() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.closed
}

// Len mengembalikan jumlah item yang saat ini ada di queue.
func (q *SPSCQueue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.count
}

// Cap mengembalikan kapasitas buffer saat ini.
// Nilai ini bisa bertambah seiring pertumbuhan queue (lihat grow).
func (q *SPSCQueue[T]) Cap() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.buffer)
}

// Snapshot mengembalikan salinan seluruh item di queue pada saat dipanggil,
// diurutkan dari depan (head) ke belakang (tail).
//
// Hasil adalah point-in-time copy — perubahan pada queue setelah Snapshot
// tidak tercermin pada slice yang dikembalikan, begitu pula sebaliknya.
// Mengembalikan nil jika queue kosong.
func (q *SPSCQueue[T]) Snapshot() []T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.count == 0 {
		return nil
	}

	out := make([]T, q.count)
	for i := range q.count {
		out[i] = q.buffer[(q.head+i)%len(q.buffer)]
	}
	return out
}

// dequeueItem mengambil item dari head buffer.
// Harus dipanggil hanya saat mu sudah di-hold dan count > 0.
func (q *SPSCQueue[T]) dequeueItem() T {
	item := q.buffer[q.head]
	var zero T
	q.buffer[q.head] = zero // hindari memory leak untuk reference types
	q.head = (q.head + 1) % len(q.buffer)
	q.count--
	return item
}

// notifyConsumer mengirim sinyal ke consumer tanpa blocking.
// Jika channel sudah terisi (consumer belum sempat baca), sinyal diabaikan —
// consumer akan tetap terbangun karena sinyal sebelumnya masih ada.
func (q *SPSCQueue[T]) notifyConsumer() {
	select {
	case q.notifyCh <- struct{}{}:
	default:
	}
}

// grow menggandakan kapasitas buffer dan menyalin item yang ada ke buffer baru.
// Dipanggil oleh Enqueue saat buffer penuh. Head di-reset ke 0 setelah copy
// untuk menghindari fragmentasi pada ring buffer.
func (q *SPSCQueue[T]) grow() {
	newBuffer := make([]T, len(q.buffer)*2)
	for i := range q.count {
		newBuffer[i] = q.buffer[(q.head+i)%len(q.buffer)]
	}
	q.buffer = newBuffer
	q.head = 0
}

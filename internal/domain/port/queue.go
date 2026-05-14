package port

type SPSCQueue[T any] interface {
	Enqueue(item T) error
	Dequeue() (T, error)
	TryDequeue() (T, bool)
	Close()
}

package message

// ExchangeInfo is an embedded 32-byte pass-thru structure inside Enter/Replace/Accept/Replace messages.
// Layout (offset within the 32-byte field):
//
//	[0..3]  orderSource      4 bytes  Alpha
//	[4]     settlementMethod 1 byte   Alpha
//	[5..31] reserved         27 bytes
type ExchangeInfo struct {
	// OrderSource is a 4-character code describing the origin of the order.
	// Position 1 (application): R, O, D, M, Z
	// Position 2 (feature): E, S, P, C, Z, or space
	// Position 3 (platform): A, B, F, Z, or space
	// Position 4 (AO): G, Z, or space
	OrderSource string // max 4 chars

	// SettlementMethod: '1' = DVP, '2' = DFOP, 0 = not set
	SettlementMethod SettlementMethod

	// Reserved holds the remaining 27 bytes (usually zeroed).
	Reserved [27]byte
}

// Bytes serialises ExchangeInfo into exactly 32 bytes (right-padded with spaces / zero).
func (e ExchangeInfo) Bytes() [32]byte {
	var b [32]byte

	// orderSource: left-justified, right-padded with spaces, 4 bytes
	src := []byte(e.OrderSource)
	for i := 0; i < 4; i++ {
		if i < len(src) {
			b[i] = src[i]
		} else {
			b[i] = ' '
		}
	}

	// settlementMethod
	if e.SettlementMethod != 0 {
		b[4] = byte(e.SettlementMethod)
	} else {
		b[4] = ' '
	}

	// bytes [5..31] remain as zero (reserved)
	return b
}

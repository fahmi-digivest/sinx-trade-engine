package codec

import (
	"encoding/binary"
	"fmt"

	ouchmessage "github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/ouch/message"
)

// Encoder serializes OUCH inbound messages into binary wire payloads.
type Encoder struct{}

// NewEncoder returns a new OUCH encoder.
func NewEncoder() *Encoder { return &Encoder{} }

// Encode serializes a supported OUCH message into its binary wire payload.
func (e *Encoder) Encode(msg any) ([]byte, error) {
	switch m := msg.(type) {
	case *ouchmessage.EnterOrder:
		return encodeEnterOrder(*m)
	case ouchmessage.EnterOrder:
		return encodeEnterOrder(m)
	case *ouchmessage.ReplaceOrder:
		return encodeReplaceOrder(*m)
	case ouchmessage.ReplaceOrder:
		return encodeReplaceOrder(m)
	case *ouchmessage.CancelOrder:
		return encodeCancelOrder(*m)
	case ouchmessage.CancelOrder:
		return encodeCancelOrder(m)
	case *ouchmessage.CancelOrderByID:
		return encodeCancelOrderByID(*m)
	case ouchmessage.CancelOrderByID:
		return encodeCancelOrderByID(m)
	default:
		return nil, fmt.Errorf("ouch/codec: unsupported encode message type %T", msg)
	}
}

func encodeEnterOrder(m ouchmessage.EnterOrder) ([]byte, error) {
	if len(m.ClientAccount) != 15 {
		return nil, fmt.Errorf("clientAccount must be exactly 15 characters")
	}

	buf := make([]byte, 113)
	buf[0] = byte(ouchmessage.MsgTypeEnterOrder)
	binary.BigEndian.PutUint64(buf[1:9], uint64(m.OrderToken))
	binary.BigEndian.PutUint32(buf[9:13], uint32(m.OrderBookID))
	buf[13] = byte(m.Side)
	binary.BigEndian.PutUint64(buf[14:22], uint64(m.Quantity))
	binary.BigEndian.PutUint64(buf[22:30], uint64(m.Price))
	buf[30] = byte(m.TimeInForce)
	buf[31] = byte(m.OpenClose)

	copy(buf[32:47], m.ClientAccount)
	buf[47] = 0

	if err := writeAlpha(buf[48:63], m.CustomerInfo); err != nil {
		return nil, fmt.Errorf("customerInfo: %w", err)
	}

	exchangeInfo := m.ExchangeInfo.Bytes()
	copy(buf[63:95], exchangeInfo[:])

	binary.BigEndian.PutUint64(buf[95:103], uint64(m.DisplayQuantity))
	buf[103] = byte(m.OrderType)
	binary.BigEndian.PutUint16(buf[104:106], uint16(m.TimeInForceData))
	buf[106] = byte(m.OrderCapacity)
	binary.BigEndian.PutUint32(buf[107:111], uint32(m.SelfMatchPreventionKey))
	binary.BigEndian.PutUint16(buf[111:113], uint16(m.Attributes))

	return buf, nil
}

func encodeReplaceOrder(m ouchmessage.ReplaceOrder) ([]byte, error) {
	buf := make([]byte, 112)
	buf[0] = byte(ouchmessage.MsgTypeReplaceOrder)
	binary.BigEndian.PutUint64(buf[1:9], uint64(m.ExistingOrderToken))
	binary.BigEndian.PutUint64(buf[9:17], uint64(m.ReplacementOrderToken))
	binary.BigEndian.PutUint64(buf[17:25], uint64(m.Quantity))
	binary.BigEndian.PutUint64(buf[25:33], uint64(m.Price))
	buf[33] = byte(m.OpenClose)

	if m.ClientAccount != "" {
		if len(m.ClientAccount) > 15 {
			return nil, fmt.Errorf("clientAccount must be at most 15 characters")
		}
		copy(buf[34:49], m.ClientAccount)
		buf[49] = 0
	}

	if err := writeAlpha(buf[50:65], m.CustomerInfo); err != nil {
		return nil, fmt.Errorf("customerInfo: %w", err)
	}

	exchangeInfo := m.ExchangeInfo.Bytes()
	copy(buf[65:97], exchangeInfo[:])

	binary.BigEndian.PutUint64(buf[97:105], uint64(m.DisplayQuantity))
	buf[105] = byte(m.TimeInForce)
	binary.BigEndian.PutUint16(buf[106:108], uint16(m.TimeInForceData))
	binary.BigEndian.PutUint32(buf[108:112], uint32(m.SelfMatchPreventionKey))

	return buf, nil
}

func encodeCancelOrder(m ouchmessage.CancelOrder) ([]byte, error) {
	buf := make([]byte, 9)
	buf[0] = byte(ouchmessage.MsgTypeCancelOrder)
	binary.BigEndian.PutUint64(buf[1:9], uint64(m.OrderToken))
	return buf, nil
}

func encodeCancelOrderByID(m ouchmessage.CancelOrderByID) ([]byte, error) {
	buf := make([]byte, 14)
	buf[0] = byte(ouchmessage.MsgTypeCancelOrderByID)
	binary.BigEndian.PutUint32(buf[1:5], uint32(m.OrderBookID))
	buf[5] = byte(m.Side)
	binary.BigEndian.PutUint64(buf[6:14], uint64(m.OrderID))
	return buf, nil
}

func writeAlpha(buf []byte, s string) error {
	if len(s) > len(buf) {
		return fmt.Errorf("value %q too long for field of length %d", s, len(buf))
	}

	for i := range buf {
		buf[i] = ' '
	}

	copy(buf, s)
	return nil
}

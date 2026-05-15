package codec

import (
	"encoding/binary"
	"fmt"
	"strings"

	ouchmessage "github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/ouch/message"
)

// Decoder parses OUCH binary payloads into concrete message structs.
type Decoder struct{}

// NewDecoder returns a new OUCH decoder.
func NewDecoder() *Decoder { return &Decoder{} }

// Decode parses a single OUCH payload into its concrete message type.
func (d *Decoder) Decode(payload []byte) (any, error) {
	if len(payload) == 0 {
		return nil, fmt.Errorf("ouch/codec: empty payload")
	}

	switch ouchmessage.MessageType(payload[0]) {
	case ouchmessage.MsgTypeOrderAccepted:
		return decodeOrderAccepted(payload)
	case ouchmessage.MsgTypeOrderRejected:
		return decodeOrderRejected(payload)
	case ouchmessage.MsgTypeOrderReplaced:
		return decodeOrderReplaced(payload)
	case ouchmessage.MsgTypeOrderCanceled:
		return decodeOrderCanceled(payload)
	case ouchmessage.MsgTypeOrderExecuted:
		return decodeOrderExecuted(payload)
	default:
		return nil, fmt.Errorf("ouch/codec: unknown message type %q", payload[0])
	}
}

func decodeOrderAccepted(p []byte) (*ouchmessage.OrderAccepted, error) {
	if len(p) < 130 {
		return nil, fmt.Errorf("order_accepted: payload too short (%d < 130)", len(p))
	}

	return &ouchmessage.OrderAccepted{
		OuchMessageType:        int8(p[0]),
		Timestamp:              int64(binary.BigEndian.Uint64(p[1:9])),
		OrderToken:             int64(binary.BigEndian.Uint64(p[9:17])),
		OrderBookID:            int32(binary.BigEndian.Uint32(p[17:21])),
		Side:                   ouchmessage.Side(p[21]),
		OrderID:                int64(binary.BigEndian.Uint64(p[22:30])),
		Quantity:               int64(binary.BigEndian.Uint64(p[30:38])),
		Price:                  int64(binary.BigEndian.Uint64(p[38:46])),
		TimeInForce:            ouchmessage.TimeInForce(p[46]),
		OpenClose:              ouchmessage.OpenClose(p[47]),
		ClientAccount:          decodeClientAccount(p[48:64]),
		OrderState:             ouchmessage.OrderState(p[64]),
		CustomerInfo:           strings.TrimRight(string(p[65:80]), " "),
		ExchangeInfo:           decodeExchangeInfo(p[80:112]),
		DisplayQuantity:        int64(binary.BigEndian.Uint64(p[112:120])),
		OrderType:              ouchmessage.OrderType(p[120]),
		TimeInForceData:        int16(binary.BigEndian.Uint16(p[121:123])),
		OrderCapacity:          ouchmessage.OrderCapacity(p[123]),
		SelfMatchPreventionKey: int32(binary.BigEndian.Uint32(p[124:128])),
		Attributes:             ouchmessage.Attributes(int16(binary.BigEndian.Uint16(p[128:130]))),
	}, nil
}

func decodeOrderRejected(p []byte) (*ouchmessage.OrderRejected, error) {
	if len(p) < 29 {
		return nil, fmt.Errorf("order_rejected: payload too short (%d < 29)", len(p))
	}

	return &ouchmessage.OrderRejected{
		OuchMessageType: int8(p[0]),
		Timestamp:       int64(binary.BigEndian.Uint64(p[1:9])),
		OrderToken:      int64(binary.BigEndian.Uint64(p[9:17])),
		OrderID:         int64(binary.BigEndian.Uint64(p[17:25])),
		RejectCode:      int32(binary.BigEndian.Uint32(p[25:29])),
	}, nil
}

func decodeOrderReplaced(p []byte) (*ouchmessage.OrderReplaced, error) {
	if len(p) < 138 {
		return nil, fmt.Errorf("order_replaced: payload too short (%d < 138)", len(p))
	}

	return &ouchmessage.OrderReplaced{
		OuchMessageType:        int8(p[0]),
		Timestamp:              int64(binary.BigEndian.Uint64(p[1:9])),
		ReplacementOrderToken:  int64(binary.BigEndian.Uint64(p[9:17])),
		PreviousOrderToken:     int64(binary.BigEndian.Uint64(p[17:25])),
		OrderBookID:            int32(binary.BigEndian.Uint32(p[25:29])),
		Side:                   ouchmessage.Side(p[29]),
		OrderID:                int64(binary.BigEndian.Uint64(p[30:38])),
		Quantity:               int64(binary.BigEndian.Uint64(p[38:46])),
		Price:                  int64(binary.BigEndian.Uint64(p[46:54])),
		TimeInForce:            ouchmessage.TimeInForce(p[54]),
		OpenClose:              ouchmessage.OpenClose(p[55]),
		ClientAccount:          decodeClientAccount(p[56:72]),
		OrderState:             ouchmessage.OrderState(p[72]),
		CustomerInfo:           strings.TrimRight(string(p[73:88]), " "),
		ExchangeInfo:           decodeExchangeInfo(p[88:120]),
		DisplayQuantity:        int64(binary.BigEndian.Uint64(p[120:128])),
		OrderType:              ouchmessage.OrderType(p[128]),
		TimeInForceData:        int16(binary.BigEndian.Uint16(p[129:131])),
		OrderCapacity:          ouchmessage.OrderCapacity(p[131]),
		SelfMatchPreventionKey: int32(binary.BigEndian.Uint32(p[132:136])),
		Attributes:             ouchmessage.Attributes(int16(binary.BigEndian.Uint16(p[136:138]))),
	}, nil
}

func decodeOrderCanceled(p []byte) (*ouchmessage.OrderCanceled, error) {
	if len(p) < 31 {
		return nil, fmt.Errorf("order_canceled: payload too short (%d < 31)", len(p))
	}

	return &ouchmessage.OrderCanceled{
		OuchMessageType: int8(p[0]),
		Timestamp:       int64(binary.BigEndian.Uint64(p[1:9])),
		OrderToken:      int64(binary.BigEndian.Uint64(p[9:17])),
		OrderBookID:     int32(binary.BigEndian.Uint32(p[17:21])),
		Side:            ouchmessage.Side(p[21]),
		OrderID:         int64(binary.BigEndian.Uint64(p[22:30])),
		CancelReason:    ouchmessage.CancelReason(p[30]),
	}, nil
}

func decodeOrderExecuted(p []byte) (*ouchmessage.OrderExecuted, error) {
	if len(p) < 50 {
		return nil, fmt.Errorf("order_executed: payload too short (%d < 50)", len(p))
	}

	return &ouchmessage.OrderExecuted{
		OuchMessageType: int8(p[0]),
		Timestamp:       int64(binary.BigEndian.Uint64(p[1:9])),
		OrderToken:      int64(binary.BigEndian.Uint64(p[9:17])),
		OrderBookID:     int32(binary.BigEndian.Uint32(p[17:21])),
		TradeQuantity:   int64(binary.BigEndian.Uint64(p[21:29])),
		TradePrice:      int64(binary.BigEndian.Uint64(p[29:37])),
		MatchID:         int64(binary.BigEndian.Uint64(p[37:45])),
		ComboGroupID:    int32(binary.BigEndian.Uint32(p[45:49])),
		DealSource:      ouchmessage.DealSource(p[49]),
	}, nil
}

func decodeClientAccount(p []byte) string {
	if len(p) == 0 {
		return ""
	}

	end := len(p)
	for i, b := range p {
		if b == 0 {
			end = i
			break
		}
	}

	return strings.TrimRight(string(p[:end]), " ")
}

func decodeExchangeInfo(p []byte) ouchmessage.ExchangeInfo {
	var info ouchmessage.ExchangeInfo
	if len(p) < 32 {
		return info
	}

	info.OrderSource = strings.TrimRight(string(p[0:4]), " ")
	if p[4] != 0 && p[4] != ' ' {
		info.SettlementMethod = ouchmessage.SettlementMethod(p[4])
	}
	copy(info.Reserved[:], p[5:32])
	return info
}

package handler

import (
	"log/slog"

	"github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/message"
	ouchcodec "github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/ouch/codec"
	ouchmessage "github.com/fahmi-digivest/sinx-trade-engine/internal/delivery/tcp/soupbin/ouch/message"
)

// OuchClientHandler logs SoupBin client events for OUCH client runtime services.
type OuchClientHandler struct {
	logger     *slog.Logger
	ouchLogger *slog.Logger
	decoder    *ouchcodec.Decoder
}

func NewOuchClientHandler(logger, ouchLogger *slog.Logger) *OuchClientHandler {
	return &OuchClientHandler{
		logger:     logger,
		ouchLogger: ouchLogger,
		decoder:    ouchcodec.NewDecoder(),
	}
}

func (h *OuchClientHandler) OnLoginAccepted(session string, nextSeq uint64) {
	h.ouchLogger.Info("login accepted", "session", session, "next_seq", nextSeq)
}

func (h *OuchClientHandler) OnSequencedData(seq uint64, msg []byte) {
	decoded, err := h.decoder.Decode(msg)
	if err != nil {
		h.logger.Error("decode ouch sequenced payload failed", "seq", seq, "err", err)
		return
	}

	switch m := decoded.(type) {
	case *ouchmessage.OrderAccepted:
		h.ouchLogger.Info(
			"order accepted",
			"seq", seq,
			"order_token", m.OrderToken,
			"order_book_id", m.OrderBookID,
			"order_id", m.OrderID,
			"side", m.Side.String(),
			"quantity", m.Quantity,
			"price", m.Price,
			"client_account", m.ClientAccount,
			"order_state", m.OrderState,
		)
	case *ouchmessage.OrderRejected:
		h.ouchLogger.Warn(
			"order rejected",
			"seq", seq,
			"order_token", m.OrderToken,
			"order_id", m.OrderID,
			"reject_code", m.RejectCode,
		)
	case *ouchmessage.OrderReplaced:
		h.ouchLogger.Info(
			"order replaced",
			"seq", seq,
			"replacement_order_token", m.ReplacementOrderToken,
			"previous_order_token", m.PreviousOrderToken,
			"order_book_id", m.OrderBookID,
			"order_id", m.OrderID,
			"quantity", m.Quantity,
			"price", m.Price,
		)
	case *ouchmessage.OrderCanceled:
		h.ouchLogger.Info(
			"order canceled",
			"seq", seq,
			"order_token", m.OrderToken,
			"order_book_id", m.OrderBookID,
			"order_id", m.OrderID,
			"cancel_reason", m.CancelReason,
		)
	case *ouchmessage.OrderExecuted:
		h.ouchLogger.Info(
			"order executed",
			"seq", seq,
			"order_token", m.OrderToken,
			"order_book_id", m.OrderBookID,
			"trade_quantity", m.TradeQuantity,
			"trade_price", m.TradePrice,
			"match_id", m.MatchID,
		)
	default:
		h.ouchLogger.Info("sequenced data received", "seq", seq, "payload_len", len(msg))
	}
}

func (h *OuchClientHandler) OnServerHeartbeat() {
	h.ouchLogger.Debug("server heartbeat received")
}

func (h *OuchClientHandler) OnUnsequencedData(msg *message.UnsequencedData) {
	h.ouchLogger.Debug("unsequenced data received", "payload_len", len(msg.Message))
}

func (h *OuchClientHandler) OnEndOfSession() {
	h.logger.Warn("end of session received")
}

func (h *OuchClientHandler) OnError(err error) {
	h.logger.Error("soupbin client error", "err", err)
}

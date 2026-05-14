package handler

import "log/slog"

// OuchClientHandler logs SoupBin client events for OUCH client runtime services.
type OuchClientHandler struct {
	logger     *slog.Logger
	ouchLogger *slog.Logger
}

func NewOuchClientHandler(logger, ouchLogger *slog.Logger) *OuchClientHandler {
	return &OuchClientHandler{
		logger:     logger,
		ouchLogger: ouchLogger,
	}
}

func (h *OuchClientHandler) OnLoginAccepted(session string, nextSeq uint64) {
	h.logger.Info("login accepted", "session", session, "next_seq", nextSeq)
}

func (h *OuchClientHandler) OnSequencedData(seq uint64, msg []byte) {
	h.ouchLogger.Debug("sequenced data received", "seq", seq, "payload_len", len(msg))
}

func (h *OuchClientHandler) OnEndOfSession() {
	h.logger.Warn("end of session received")
}

func (h *OuchClientHandler) OnError(err error) {
	h.logger.Error("soupbin client error", "err", err)
}

package handler

import (
	"mini_chat/internal/configs"
	"mini_chat/pkg/logger"
	"net"
)

type Handler struct {
	cfg *configs.ServerConfig
	log *logger.Logger
}

func NewHandler(cfg *configs.ServerConfig, log *logger.Logger) *Handler {
	return &Handler{
		cfg: cfg,
		log: log,
	}
}

/*
Writing in file in Unix use POSIX IO syscalls,
so they are safe as the underlying syscalls
*/
func (h *Handler) InitRoutes() {
	l, err := net.Listen(h.cfg.ListenType, h.cfg.ListenAddr)
	if err != nil {
		h.log.Error(err.Error())
		return
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			h.log.Error(err.Error())
		}

		go h.handleConn(conn)
	}
}

func (h *Handler) handleConn(c net.Conn) {
	_, err := c.Write([]byte("Hello world"))

	if err != nil {
		h.log.Error(err.Error())
		return
	}
}

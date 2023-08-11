package handler

import (
	"fmt"
	"mini_chat/internal/config"
	"mini_chat/pkg/logger"
	"net"
)

type Handler struct {
	cfg *config.ServerConfig
	log *logger.Logger
}

func NewHandler(cfg *config.ServerConfig, log *logger.Logger) *Handler {
	return &Handler{
		cfg: cfg,
		log: log,
	}
}

/*
Writing in file in Unix use POSIX IO syscalls,
so they are safe as the underlying syscalls.
*/
func (h *Handler) InitRoutes() error {
	l, err := net.Listen(h.cfg.ListenType, fmt.Sprintf("%s:%s", h.cfg.ListenAddr, h.cfg.ListenPort))

	if err != nil {
		return err
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
	defer c.Close()
	remoteAddr := c.RemoteAddr()
	remoteIP, _, err := net.SplitHostPort(remoteAddr.String())
	if err != nil {
		h.log.Error(err.Error())
		return
	}

	response := fmt.Sprintf("hello boys!\nremote addr: %s | remote IP: %s\n", remoteAddr.String(), remoteIP)

	_, err = c.Write([]byte(response))
	if err != nil {
		h.log.Error(err.Error())
		return
	}
}

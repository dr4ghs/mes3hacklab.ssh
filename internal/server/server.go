package server

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"charm.land/log/v2"
	"charm.land/wish/v2"
	"charm.land/wish/v2/activeterm"
	"charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/constants"
	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui"
)

type Server struct{}

func New() Server {
	return Server{}
}

func (s *Server) Start() (err error) {
	var srv *ssh.Server
	srv, err = wish.NewServer(
		wish.WithAddress(net.JoinHostPort(
			constants.SrvHost(),
			constants.SrvPort(),
		)),
		wish.WithHostKeyPath(constants.SSHKeyPath()),
		wish.WithMiddleware(
			bubbletea.Middleware(tui.Handler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		return
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Starting SSH server", "host", constants.SrvHost(), "port", constants.SrvPort())
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			return
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	log.Info("Stopping SSH server")
	if err = srv.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		return
	}

	return
}

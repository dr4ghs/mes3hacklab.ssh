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
	"github.com/joho/godotenv"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/tui"
)

const (
	_EnvSrvHostName    = "SRV_HOST"
	_EnvSrvPortName    = "SRV_PORT"
	_EnvSSHKeyPathName = "SSH_KEY_PATH"
)

var (
	host       string
	port       string
	sshkeyPath string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot find .env file: %v\n", err)
	}

	if host = os.Getenv(_EnvSrvHostName); host == "" {
		host = "localhost"
	}

	if port = os.Getenv(_EnvSrvPortName); port == "" {
		port = "46593"
	}

	if sshkeyPath = os.Getenv(_EnvSSHKeyPathName); sshkeyPath == "" {
		sshkeyPath = "~/.ssh/id_ed25519"
	}
}

type Server struct{}

func New() Server {
	return Server{}
}

func (s *Server) Start() (err error) {
	var srv *ssh.Server
	srv, err = wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(sshkeyPath),
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
		log.Info("Starting SSH server", "host", host, "port", port)
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

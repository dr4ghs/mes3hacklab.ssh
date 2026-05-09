package main

import (
	"charm.land/log/v2"

	"github.com/dr4hgs/mes3hacklab.ssh/internal/server"
)

func main() {
	srv := server.New()
	if err := srv.Start(); err != nil {
		log.Error("Could not start server", "error", err)
	}
}

package main

import (
	"forward/internal/config"
	"forward/internal/server"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	srv, err := server.New()
	if err != nil {
		panic(err)
	}

	if err := srv.Start(); err != nil {
		panic(err)
	}
}

package server

import (
	"context"
	"fmt"
	"forward/internal/client"
	"forward/internal/config"
	"forward/internal/handler"
	"forward/internal/repository"
	"forward/internal/router"
	"forward/internal/service"

	"github.com/labstack/echo/v5"
)

type Server struct {
	config *config.ServerConfig
	client *client.Client
	router *echo.Echo
}

func New() (*Server, error) {
	cfg := config.GetServerConfig()

	ctx := context.Background()

	clt, err := client.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	repo := repository.New(clt)
	svc := service.New(repo)
	hdlr := handler.New(svc)
	rtr := router.New(hdlr)

	return &Server{
		config: cfg,
		client: clt,
		router: rtr,
	}, nil
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	return s.router.Start(addr)
}

func (s *Server) Shutdown() {
	s.client.Close()
}

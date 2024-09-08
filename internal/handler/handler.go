package handler

import (
	"log/slog"

	"github.com/tracker-tv/actor-api/internal/service"
)

type Handler struct {
	Logger  *slog.Logger
	Service *service.Service
}

func NewHandler(logger *slog.Logger, service *service.Service) *Handler {
	return &Handler{Logger: logger, Service: service}
}

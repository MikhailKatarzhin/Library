package main

import (
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/MikhailKatarzhin/Library/internal/library"
	bookmemory "github.com/MikhailKatarzhin/Library/internal/library/book/repository/memory"
	"github.com/MikhailKatarzhin/Library/pkg/logger"
)

const (
	port    = 8080
	timeout = 5 * time.Second
)

func main() {
	server := library.NewHTTPServer(
		port,
		library.NewService(
			bookmemory.NewRepository(),
		),
		library.WithTimeout(timeout),
	)
	server.Start()

	resp, err := http.Post(
		"http://localhost:8080/book",
		"application/json",
		strings.NewReader(`{"title": "test", "description": "test"}`),
	)
	if err != nil {
		logger.I().Error("failed to send request", zap.Error(err))
	} else {
		logger.I().Info("response", zap.Int("status", resp.StatusCode))
	}

	if err = server.Close(); err != nil {
		logger.I().Error("failed to close http server", zap.Error(err))
	}

	logger.I().Info("application is closed")
}

package library

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/MikhailKatarzhin/Library/internal/library/book/info"
	"github.com/MikhailKatarzhin/Library/pkg/logger"
)

const DefaultTimeout = 30 * time.Second

type HTTPServerOption func(*HTTPServer)

func WithTimeout(timeout time.Duration) HTTPServerOption {
	return func(server *HTTPServer) {
		server.serv.ReadTimeout = timeout
		server.serv.WriteTimeout = timeout
	}
}

type HTTPServer struct {
	serv           *http.Server
	libraryService *Service
}

func NewHTTPServer(
	port int, libraryService *Service, options ...HTTPServerOption,
) *HTTPServer {
	result := &HTTPServer{
		serv: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			ReadTimeout:  DefaultTimeout,
			WriteTimeout: DefaultTimeout,
		},
		libraryService: libraryService,
	}

	for _, option := range options {
		option(result)
	}

	result.registerHandlers()

	return result
}

func (s *HTTPServer) Start() {
	go func() {
		if err := s.serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.I().Error("failed to start http server", zap.Error(err))
		}
	}()
}

func (s *HTTPServer) Close() error {
	return s.serv.Close() //nolint:wrapcheck
}

func (s *HTTPServer) registerHandlers() {
	http.HandleFunc("/book", s.handleBook)
}

func (s *HTTPServer) handleBook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.handleBookPost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type BookJSON struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *HTTPServer) handleBookPost(w http.ResponseWriter, r *http.Request) {
	book := new(BookJSON)
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		response := fmt.Sprintf("failed to decode book: %v", err)
		if _, err = w.Write([]byte(response)); err != nil {
			logger.I().Error(
				"failed to write response",
				zap.Error(err), zap.String("response", response),
			)
		}

		return
	}

	if err := s.libraryService.AddNewBook(r.Context(), info.BookInfo{
		Title:       book.Title,
		Description: book.Description,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.I().Error("failed to add new book", zap.Error(err))

		return
	}

	w.WriteHeader(http.StatusCreated)
}

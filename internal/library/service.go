package library

import (
	"context"
	"fmt"

	"github.com/MikhailKatarzhin/Library/internal/library/book"
	"github.com/MikhailKatarzhin/Library/internal/library/book/info"
)

type BookRepository interface {
	SaveBook(ctx context.Context, book book.Book) error
}

type Service struct {
	bookRepo BookRepository
}

func NewService(bookRepo BookRepository) *Service {
	return &Service{
		bookRepo: bookRepo,
	}
}

func (s *Service) AddNewBook(
	ctx context.Context, bookinfo info.BookInfo,
) error {
	if err := s.bookRepo.SaveBook(ctx, book.NewBook(bookinfo)); err != nil {
		return fmt.Errorf("failed to save book: %w", err)
	}

	return nil
}

package bookmemory

import (
	"context"
	"errors"

	"github.com/MikhailKatarzhin/Library/internal/library/book"
)

var ErrNotFound = errors.New("book not found")

type Repository struct {
	books []*book.Book
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SaveBook(ctx context.Context, book book.Book) error {
	r.books = append(r.books, &book)

	return nil
}

func (r *Repository) FindBookByTitle(ctx context.Context, title string) (*book.Book, error) {
	for _, b := range r.books {
		if b.Info().Title == title {
			return b, nil
		}
	}

	return nil, ErrNotFound
}

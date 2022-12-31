package bookmemory

import (
	"context"

	"github.com/MikhailKatarzhin/Library/internal/library/book"
)

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

// Package book describe the model of the book.
package book

import (
	"github.com/MikhailKatarzhin/Library/internal/library/book/id"
	"github.com/MikhailKatarzhin/Library/internal/library/book/info"
)

type Book struct {
	id   id.ID
	info info.BookInfo
}

func NewBook(info info.BookInfo) Book {
	return Book{
		id:   id.NewID(),
		info: info,
	}
}

func (b Book) ID() id.ID {
	return b.id
}

func (b Book) Info() info.BookInfo {
	return b.info
}

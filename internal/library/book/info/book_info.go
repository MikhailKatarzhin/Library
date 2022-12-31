// Package info provides the book's meta information.
package info

import (
	"github.com/MikhailKatarzhin/Library/internal/library/book/id"
)

type BookInfo struct {
	// Use it later.
	authors     []id.ID
	Title       string
	Description string
}

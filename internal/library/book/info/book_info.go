// Package info provides the book's meta information.
package info

import (
	"github.com/MikhailKatarzhin/Library/internal/library/book/id"
)

type BookInfo struct {
	authors     []*id.ID //nolint:unused // Use it later.
	Title       string
	Description string
}

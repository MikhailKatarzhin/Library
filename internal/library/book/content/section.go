package content

import (
	"github.com/MikhailKatarzhin/Library/internal/library/book/id"
)

type Section struct {
	id      id.ID
	content string
}

func NewSection(sectionId id.ID) *Section {
	return &Section{
		id: sectionId,
	}
}

func (s Section) ID() id.ID {
	return s.id
}

func (s Section) Content() string {
	return s.content
}

func (s *Section) UpdateContent(newContent string) {
	s.content = newContent
}

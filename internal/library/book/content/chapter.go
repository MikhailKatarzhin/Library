package content

import (
	"github.com/MikhailKatarzhin/Library/internal/library/book/id"
)

type Chapter struct {
	id        id.ID
	title     string
	sections  []*Section
	idToIndex map[id.ID]int
}

func (c Chapter) ID() id.ID {
	return c.id
}

func (c Chapter) Title() string {
	return c.title
}

func (c Chapter) Sections() []*Section {
	return c.sections
}

func (c Chapter) IdToIndex() map[id.ID]int {
	return c.idToIndex
}

func NewChapter(contentID id.ID) *Chapter {
	return &Chapter{
		id:        contentID,
		idToIndex: make(map[id.ID]int, 0),
	}
}

func (c Chapter) FindSectionById(sectionID id.ID) *Section {
	return c.Sections()[c.IdToIndex()[sectionID]]
}

func (c Chapter) FindSectionByIndex(index int) *Section {
	return c.Sections()[index]
}

func (c *Chapter) UpdateSection(
	sectionID id.ID,
	f func(*Section),
) {
	sect := c.FindSectionById(sectionID)
	f(sect)
}

func (c *Chapter) AddSectionToEnd(section *Section) error {
	err := c.RemoveSection(section.ID())
	if err != nil {
		return err
	}
	c.sections = append(c.sections, section)
	c.idToIndex[section.ID()] = len(c.idToIndex)
	return nil
}

func (c *Chapter) InsertSectionAtIndex(section *Section, index int) error {
	err := c.RemoveSection(section.ID())
	if err != nil {
		return err
	}
	c.sections = append(append(c.sections[:index], section), c.sections[index:]...)
	c.idToIndex[section.ID()] = index
	for i := index + 1; i < len(c.sections); {
		c.idToIndex[id.ID(i)] = i
	}
	return nil
}

func (c *Chapter) RemoveSection(sectionId id.ID) error {
	return c.removeSectionByIdIfExist(sectionId)
}

func (c *Chapter) removeSectionByIdIfExist(sectionId id.ID) error {
	_, isExists := c.idToIndex[sectionId]
	if isExists {
		index := c.idToIndex[sectionId]
		if len(c.sections) < index+1 {
			c.sections = c.sections[:index]
		} else {
			c.sections = append(c.sections[:index], c.sections[index+1:]...)
		}
		delete(c.idToIndex, sectionId)
		for i := index; i < len(c.sections); {
			c.idToIndex[id.ID(i)] = i
		}
	}
	return nil
}

func (c *Chapter) SwapSectionsByIDs(id1, id2 id.ID) {
	_, isExists := c.idToIndex[id1]
	if !isExists {
		return
	}
	_, isExists = c.idToIndex[id2]
	if !isExists {
		return
	}

	index1 := c.IdToIndex()[id1]
	index2 := c.IdToIndex()[id2]
	c.IdToIndex()[id1] = index2
	c.IdToIndex()[id2] = index1
	tmpChapter := c.Sections()[index1]
	c.Sections()[index1] = c.Sections()[index2]
	c.Sections()[index2] = tmpChapter
}

func (c *Chapter) SwapSectionsByIndexes(index1, index2 int) {
	if index1 < len(c.sections) && index2 < len(c.sections) {
		chapter1 := c.Sections()[index1]
		c.Sections()[index1] = c.Sections()[index2]
		c.Sections()[index2] = chapter1
		c.IdToIndex()[chapter1.ID()] = index2
		c.IdToIndex()[c.Sections()[index1].ID()] = index1
	}
}

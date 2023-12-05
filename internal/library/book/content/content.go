package content

import (
	"github.com/MikhailKatarzhin/Library/internal/library/book/id"
)

type Content struct {
	id        id.ID
	chapters  []*Chapter
	idToIndex map[id.ID]int
}

func (c Content) ID() id.ID {
	return c.id
}

func (c Content) Chapters() []*Chapter {
	return c.chapters
}

func (c Content) IdToIndex() map[id.ID]int {
	return c.idToIndex
}

func NewContent(contentID id.ID) *Content {
	return &Content{
		id:        contentID,
		idToIndex: make(map[id.ID]int, 0),
	}
}

func (c Content) FindChapterById(chapterID id.ID) *Chapter {
	return c.Chapters()[c.IdToIndex()[chapterID]]
}

func (c Content) FindChapterByIndex(index int) *Chapter {
	return c.Chapters()[index]
}

func (c *Content) UpdateChapter(
	chapterID id.ID,
	f func(*Chapter),
) {
	chap := c.FindChapterById(chapterID)
	f(chap)
}

func (c *Content) AddChapterToEnd(chapter *Chapter) error {
	err := c.removeChapterByIdIfExist(chapter.ID())
	if err != nil {
		return err
	}
	c.chapters = append(c.chapters, chapter)
	c.idToIndex[chapter.ID()] = len(c.idToIndex)
	return nil
}

func (c *Content) InsertChapterAtIndex(chapter *Chapter, index int) error {
	err := c.removeChapterByIdIfExist(chapter.ID())
	if err != nil {
		return err
	}
	c.chapters = append(append(c.chapters[:index], chapter), c.chapters[index:]...)
	c.idToIndex[chapter.ID()] = index
	for i := index + 1; i < len(c.chapters); {
		c.idToIndex[id.ID(i)] = i
	}
	return nil
}

func (c *Content) RemoveChapter(chapterId id.ID) error {
	return c.removeChapterByIdIfExist(chapterId)
}

func (c *Content) removeChapterByIdIfExist(chapterId id.ID) error {
	_, isExists := c.idToIndex[chapterId]
	if isExists {
		index := c.idToIndex[chapterId]
		if len(c.chapters) < index+1 {
			c.chapters = c.chapters[:index]
		} else {
			c.chapters = append(c.chapters[:index], c.chapters[index+1:]...)
		}
		delete(c.idToIndex, chapterId)
		for i := index; i < len(c.chapters); {
			c.idToIndex[id.ID(i)] = i
		}
	}
	return nil
}

func (c *Content) SwapChaptersByIDs(id1, id2 id.ID) {
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
	tmpChapter := c.Chapters()[index1]
	c.Chapters()[index1] = c.Chapters()[index2]
	c.Chapters()[index2] = tmpChapter
}

func (c *Content) SwapChaptersByIndexes(index1, index2 int) {
	if index1 < len(c.chapters) && index2 < len(c.chapters) {
		chapter1 := c.chapters[index1]
		c.Chapters()[index1] = c.Chapters()[index2]
		c.Chapters()[index2] = chapter1
		c.IdToIndex()[chapter1.ID()] = index2
		c.IdToIndex()[c.Chapters()[index1].ID()] = index1
	}
}

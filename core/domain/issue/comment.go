package domain

import (
	"fmt"
	"gopt/core/domain/common/model"
	"time"
)

type Comment struct {
	author            CommentAuthor // 4
	model.TimeTracked               // 2
	content           string        // 3
	parentItemId      int           // 1
	model.Entity                    // 5
}

func NewComment(parentItemId int, content string) Comment {
	comment := Comment{
		Entity:       model.Entity{},
		TimeTracked:  model.NewTimeTracked(time.Now(), time.Now()),
		parentItemId: parentItemId,
		content:      content,
	}
	return comment
}

func (c Comment) GetContent() string {
	return c.content
}

func (c Comment) String() string {
	return fmt.Sprintf("Comment[id:%v; parentId:%v; content:%v]", c.GetId(), c.parentItemId, c.content)
}

func (c Comment) GetParentItemId() int {
	return c.parentItemId
}

type CommentAuthor struct {
	Name string
}

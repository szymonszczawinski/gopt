package domain

import (
	"fmt"
	"gosi/domain/common/model"
	"time"
)

type Comment struct {
	model.Entity
	model.TimeTracked
	parentItemId int
	content      string
	author       CommentAuthor
}

func NewComment(parentItemId int, content string) Comment {
	comment := Comment{
		Entity: model.Entity{},
		TimeTracked: model.TimeTracked{
			Created: time.Now(),
			Updated: time.Now(),
		},
		parentItemId: parentItemId,
		content:      content,
	}
	return comment
}

func (self Comment) GetContent() string {
	return self.content
}

func (self Comment) String() string {
	return fmt.Sprintf("Comment[id:%v; parentId:%v; content:%v]", self.GetId(), self.parentItemId, self.content)
}

func (self Comment) GetParentItemId() int {
	return self.parentItemId
}

type CommentAuthor struct {
	Name string
}

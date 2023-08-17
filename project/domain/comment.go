package domain

import (
	"fmt"
	"gosi/common/model"
	"gosi/user"
	"time"
)

type Comment struct {
	model.Entity
	model.TimeTracked
	parentItemId int
	content      string
	author       user.User
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

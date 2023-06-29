package domain

import (
	"fmt"
	"gosi/users"
	"time"
)

type Comment struct {
	Entity
	TimeTracked
	parentItemId int
	content      string
	author       users.User
}

func (self Comment) GetContent() string {
	return self.content
}

func (self Comment) String() string {
	return fmt.Sprintf("Comment[id:%v; parentId:%v; content:%v]", self.id, self.parentItemId, self.content)
}

func (self Comment) GetParentItemId() int {
	return self.parentItemId
}

func NewComment(parentItemId int, content string) Comment {
	comment := Comment{
		Entity: Entity{},
		TimeTracked: TimeTracked{
			created: time.Now(),
			updated: time.Now(),
		},
		parentItemId: parentItemId,
		content:      content,
	}
	return comment
}

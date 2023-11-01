package project

import (
	"time"

	"github.com/uptrace/bun"
)

type ProjectRow struct {
	bun.BaseModel `bun:"table:project"`
	Id            int       `bun:"id,pk,autoincrement"`
	Created       time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Updated       time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Name          string
	ProjectKey    string
	Description   string
	StateId       int
	LifecycleId   int
	CreatedById   int
	OwnerId       int
}

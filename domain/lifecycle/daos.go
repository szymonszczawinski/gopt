package lifecycle

import "github.com/uptrace/bun"

type LifecycleStateRow struct {
	bun.BaseModel `bun:"table:lifecyclestate"`
	Id            int `bun:"id,pk,autoincrement"`
	Name          string
}

type LifecycleRow struct {
	bun.BaseModel `bun:"table:lifecycle"`
	Id            int `bun:"id,pk,autoincrement"`
	Name          string
	StartStateId  int
}

type StateTransition struct {
	bun.BaseModel `bun:"table:statetransition"`
	LifecycleId   int
	FromStateId   int
	ToStateId     int
}

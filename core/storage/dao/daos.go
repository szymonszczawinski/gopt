package dao

import "time"

type ProjectRow struct {
	Id          int
	Created     time.Time
	Updated     time.Time
	Name        string
	ItemKey     string
	ItemNumber  int
	Description string
	StateId     int
	LifecycleId int
	CreatedById int
}

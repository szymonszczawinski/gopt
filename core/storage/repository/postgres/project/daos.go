package project

import (
	"time"
)

// representation of database project entry
type ProjectRow struct {
	Created     time.Time
	Updated     time.Time
	Name        string
	ProjectKey  string
	Description string
	CreatorName string
	OwnerName   string
	StateName   string
	StateId     int
	LifecycleId int
	CreatedById int
	OwnerId     int
	Id          int
}

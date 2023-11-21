package model

import "time"

type Named interface {
	GetName() string
}

type Entity struct {
	Id int
}

func (e Entity) GetId() int {
	return e.Id
}

func (e *Entity) SetId(id int) {
	e.Id = id
}

type TimeTracked struct {
	Created time.Time
	Updated time.Time
}

func (tt TimeTracked) GetCreationTime() time.Time {
	return tt.Created
}

func (tt TimeTracked) GetLastUpdateTime() time.Time {
	return tt.Updated
}

func (tt *TimeTracked) SetLastUpdateTime(lastUpdated time.Time) {
	tt.Updated = lastUpdated
}

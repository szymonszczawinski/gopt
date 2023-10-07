package model

import "time"

type Named interface {
	GetName() string
}

type Entity struct {
	Id int
}

func (self Entity) GetId() int {
	return self.Id
}

func (self *Entity) SetId(id int) {
	self.Id = id
}

type TimeTracked struct {
	Created time.Time
	Updated time.Time
}

func (self TimeTracked) GetCreationTime() time.Time {
	return self.Created
}

func (self TimeTracked) GetLastUpdateTime() time.Time {
	return self.Updated
}

func (self *TimeTracked) SetLastUpdateTime(lastUpdated time.Time) {
	self.Updated = lastUpdated
}

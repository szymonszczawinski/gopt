package domain

import "time"

type Entity struct {
	id int
}

func (self Entity) GetId() int {
	return self.id
}

func (self *Entity) SetId(id int) {
	self.id = id
}

type TimeTracked struct {
	created time.Time
	updated time.Time
}

func (self TimeTracked) GetCreationTime() time.Time {
	return self.created
}

func (self TimeTracked) GetLastUpdateTime() time.Time {
	return self.updated
}

func (self *TimeTracked) SetLastUpdateTime(lastUpdated time.Time) {
	self.updated = lastUpdated
}

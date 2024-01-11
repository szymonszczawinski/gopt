package user

import "github.com/uptrace/bun"

type UserRow struct {
	bun.BaseModel `bun:"table:users"`
	Id            int `bun:"id,pk,autoincrement"`
	FirstName     string
	LastName      string
	Email         string
}

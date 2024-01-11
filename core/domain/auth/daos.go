package auth

import "github.com/uptrace/bun"

type AuthCredentialsDao struct {
	bun.BaseModel `bun:"table:usercredentials"`
	Id            int    `bun:"id,pk,autoincrement"`
	Username      string `bun:",notnull"`
	Password      string `bun:",notnull"`
}

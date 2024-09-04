package auth

import "github.com/uptrace/bun"

type AuthCredentialsDao struct {
	bun.BaseModel `bun:"table:usercredentials"`
	Username      string `bun:",notnull"`
	Password      string `bun:",notnull"`
	Id            int    `bun:"id,pk,autoincrement"`
}

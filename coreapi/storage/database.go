package storage

import (
	"gosi/coreapi/service"

	"github.com/uptrace/bun"
)

type IBunDatabase interface {
	service.IComponent
	NewSelect() *bun.SelectQuery
	NewInsert() *bun.InsertQuery
	NewRaw(sql string, args ...interface{}) *bun.RawQuery
}

type ISqlDatabase interface {
	service.IComponent
}

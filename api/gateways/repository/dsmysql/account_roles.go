package dsmysql

import (
	"context"
	"echo-admin-templte/entity"
)

type AccountRoleDatasource interface {
	Insert(ctx context.Context, e entity.AccountRole) (entity.AccountRole, error)
}

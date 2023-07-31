package dsmysql

import (
	"context"
	"echo-admin-templte/entity"
)

type AccountDatasource interface {
	SelectByLogin(ctx context.Context, e entity.Login) (entity.Account, error)
	Insert(ctx context.Context, e entity.Account) (entity.Account, error)
}

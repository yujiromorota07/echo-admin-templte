package repository

import (
	"context"
	"echo-admin-templte/entity"
)

type AccountRepository interface {
	GetAccountByLogin(ctx context.Context, id entity.Login) (entity.Account, error)
	CreateAccount(ctx context.Context, e entity.Account) (entity.AccountDetail, error)
}

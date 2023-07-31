package inputport

import (
	"echo-admin-templte/entity"
)

type Account interface {
	CreateAccount(entity.Account) (entity.AccountDetail, error)
}

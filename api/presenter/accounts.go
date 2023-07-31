package presenter

import (
	"echo-admin-templte/entity"
	"echo-admin-templte/generated/account"

	"github.com/deepmap/oapi-codegen/pkg/types"
)

func GetAccount(e entity.AccountDetail) account.Account {
	return account.Account{
		Id:       int32(e.ID),
		Name:     e.Name,
		Email:    types.Email(e.Email),
		Password: &e.Password,
		RoleId:   int32(e.RoleID),
	}
}

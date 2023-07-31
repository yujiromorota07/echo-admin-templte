package account

import (
	"context"
	"echo-admin-templte/entity"
	"echo-admin-templte/usecase/inputport"
	"echo-admin-templte/usecase/repository"
	util "echo-admin-templte/utils"
)

type AccountUsecase struct {
	accountRepository repository.AccountRepository
}

func NewAccountUsecase(accountRepository repository.AccountRepository) inputport.Account {
	return AccountUsecase{
		accountRepository: accountRepository,
	}
}

func (accountUsecase AccountUsecase) CreateAccount(e entity.Account) (entity.AccountDetail, error) {
	ctx := context.Background()

	password := util.NewPassword(e.Password)
	a := entity.Account{
		Name:     e.Name,
		Email:    e.Email,
		Password: password.Hash,
	}

	account, err := accountUsecase.accountRepository.CreateAccount(ctx, a)
	if err != nil {
		return entity.AccountDetail{}, err
	}

	return account, nil
}

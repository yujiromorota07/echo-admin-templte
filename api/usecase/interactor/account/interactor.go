package account

import (
	"context"
	"echo-admin-templte/entity"
	"echo-admin-templte/usecase/inputport"
	"echo-admin-templte/usecase/repository"
	util "echo-admin-templte/utils"
)

type AccountUsecase struct {
	accountRepository  repository.AccountRepository
	transactionManager repository.TransactionManager
}

func NewAccountUsecase(accountRepository repository.AccountRepository, transactionManager repository.TransactionManager) inputport.Account {
	return AccountUsecase{
		accountRepository:  accountRepository,
		transactionManager: transactionManager,
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

	v, err := accountUsecase.transactionManager.Do(ctx, func(ctx context.Context) (interface{}, error) {
		return accountUsecase.accountRepository.CreateAccount(ctx, a)
	})

	if err != nil {
		return entity.AccountDetail{}, err
	}

	account := v.(entity.AccountDetail)

	return account, nil
}

package account

import (
	"context"
	"echo-admin-templte/entity"
	"echo-admin-templte/gateways/repository/dsmysql"
	"echo-admin-templte/usecase/repository"
)

type AccountRepository struct {
	accountDatasource     dsmysql.AccountDatasource
	accountRoleDatasource dsmysql.AccountRoleDatasource
}

func NewAccountRepository(
	accountDatasource dsmysql.AccountDatasource,
	accountRoleDatasource dsmysql.AccountRoleDatasource,
) repository.AccountRepository {
	return &AccountRepository{
		accountDatasource:     accountDatasource,
		accountRoleDatasource: accountRoleDatasource,
	}
}

func (accountRepo AccountRepository) GetAccountByLogin(ctx context.Context, e entity.Login) (entity.Account, error) {
	return accountRepo.accountDatasource.SelectByLogin(ctx, e)
}

func (accountRepo AccountRepository) CreateAccount(ctx context.Context, e entity.Account) (entity.AccountDetail, error) {
	account, err := accountRepo.accountDatasource.Insert(ctx, e)
	if err != nil {
		return entity.AccountDetail{}, err
	}

	accountRole, err := accountRepo.accountRoleDatasource.Insert(ctx, entity.AccountRole{
		AccountID: account.ID,
		RoleID:    uint32(1),
	})

	if err != nil {
		return entity.AccountDetail{}, err
	}

	return entity.AccountDetail{
		ID:       account.ID,
		Name:     account.Name,
		Email:    account.Email,
		Password: account.Password,
		RoleID:   accountRole.RoleID,
	}, nil
}

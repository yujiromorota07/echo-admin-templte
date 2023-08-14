//go:build wireinject
// +build wireinject

package di

import (
	accountCtrl "echo-admin-templte/controller/mainapi"
	loginCtrl "echo-admin-templte/controller/mainapi"

	accountDatasource "echo-admin-templte/gateways/datasource/dsmysqlimpl"
	accountRoleDatasource "echo-admin-templte/gateways/datasource/dsmysqlimpl"
	accountRepository "echo-admin-templte/gateways/repository/account"
	dsmysql "echo-admin-templte/gateways/repository/dsmysql"
	transactionManager "echo-admin-templte/gateways/repository/transaction"
	"echo-admin-templte/generated/account"
	"echo-admin-templte/generated/login"
	inputport "echo-admin-templte/usecase/inputport"
	accountUsecase "echo-admin-templte/usecase/interactor/account"
	loginUsecase "echo-admin-templte/usecase/interactor/login"
	repository "echo-admin-templte/usecase/repository"

	"github.com/google/wire"
)

func IntializeLoginController() login.ServerInterface {
	wire.Build(
		loginCtrl.NewLoginController,
		IntializeLoginUsecase,
	)

	return loginCtrl.LoginController{}
}

func IntializeAccountController() account.ServerInterface {
	wire.Build(
		accountCtrl.NewAccountController,
		IntializeAccountUsecase,
	)

	return accountCtrl.AccountController{}
}

func IntializeLoginUsecase() inputport.Login {
	wire.Build(
		loginUsecase.NewLoginUsecase,
		IntializeAccountRepository,
	)
	return loginUsecase.LoginUsecase{}
}

func IntializeAccountUsecase() inputport.Account {
	wire.Build(
		accountUsecase.NewAccountUsecase,
		IntializeAccountRepository,
		IntializeTransactionManager,
	)
	return accountUsecase.AccountUsecase{}
}

func IntializeAccountRepository() repository.AccountRepository {
	wire.Build(
		accountRepository.NewAccountRepository,
		IntializeAccountDatasource,
		IntializeAccountRoleDatasource,
	)

	return accountRepository.AccountRepository{}
}

func IntializeTransactionManager() repository.TransactionManager {
	wire.Build(
		transactionManager.NewTransactionManageer,
	)

	return transactionManager.TransactionManager{}
}

func IntializeAccountDatasource() dsmysql.AccountDatasource {
	wire.Build(
		accountDatasource.NewAccountDatasource,
	)

	return accountDatasource.AccountDatasource{}
}

func IntializeAccountRoleDatasource() dsmysql.AccountRoleDatasource {
	wire.Build(
		accountRoleDatasource.NewAccountRoleDatasource,
	)

	return accountRoleDatasource.AccountRoleDatasource{}
}

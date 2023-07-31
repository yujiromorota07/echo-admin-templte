package dsmysqlimpl

import (
	"context"
	"echo-admin-templte/entity"
	"echo-admin-templte/gateways/infra/inframysql"
	"echo-admin-templte/gateways/repository/dsmysql"
)

type AccountDatasource struct{}

func NewAccountDatasource() dsmysql.AccountDatasource {
	return AccountDatasource{}
}

const (
	querySelectAccountByLogin = "SELECT `id`, `name`, `email`, `password` FROM `accounts` WHERE `email`=? AND `deleted_at` IS NULL"
	queryInsertAccount        = "INSERT INTO accounts(name,email,password) VALUE(?,?,?)"
)

func (ds AccountDatasource) SelectByLogin(ctx context.Context, e entity.Login) (entity.Account, error) {
	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(querySelectAccountByLogin)
	if err != nil {
		return entity.Account{}, err
	}

	defer stmt.Close()

	var account entity.Account
	result := stmt.QueryRow(e.Email)
	if getErr := result.Scan(&account.ID, &account.Name, &account.Email, &account.Password); getErr != nil {
		return entity.Account{}, getErr
	}
	return account, nil
}

func (ds AccountDatasource) Insert(ctx context.Context, e entity.Account) (entity.Account, error) {
	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(queryInsertAccount)
	if err != nil {
		return entity.Account{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Email, e.Password)
	if err != nil {
		return entity.Account{}, err
	}

	id, _ := result.LastInsertId()
	e.ID = uint32(id)

	return e, nil
}

package dsmysqlimpl

import (
	"context"
	"echo-admin-templte/entity"
	"echo-admin-templte/gateways/infra/inframysql"
	"echo-admin-templte/gateways/repository/dsmysql"
	"fmt"
)

type AccountRoleDatasource struct{}

func NewAccountRoleDatasource() dsmysql.AccountRoleDatasource {
	return AccountRoleDatasource{}
}

const (
	queryInsertAccountRole = "INSERT INTO account_roles(account_id,role_id) VALUE(?,?)"
)

func (accountRolesDatasource AccountRoleDatasource) Insert(ctx context.Context, e entity.AccountRole) (entity.AccountRole, error) {
	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(queryInsertAccountRole)
	if err != nil {
		return entity.AccountRole{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.AccountID, e.RoleID)
	if err != nil {
		return entity.AccountRole{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return entity.AccountRole{}, err
	}

	e.ID = uint32(id)
	return e, nil
}

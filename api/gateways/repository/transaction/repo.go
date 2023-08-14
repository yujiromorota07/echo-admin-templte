package transaction

import (
	"context"
	"echo-admin-templte/gateways/infra/inframysql"
	"echo-admin-templte/usecase/repository"
)

type TransactionManager struct{}

func NewTransactionManageer() repository.TransactionManager {
	return TransactionManager{}
}

func (transactionManager TransactionManager) Do(ctx context.Context, runner func(ctx context.Context) (interface{}, error)) (interface{}, error) {

	tx, err := inframysql.Client.Begin()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, &inframysql.TxKey, tx)
	v, err := runner(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return v, nil

}

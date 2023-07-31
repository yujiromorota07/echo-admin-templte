package mainapi

import (
	"echo-admin-templte/entity"
	"echo-admin-templte/generated/account"
	"echo-admin-templte/presenter"
	"echo-admin-templte/usecase/inputport"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountController struct {
	accountUsecase inputport.Account
}

func NewAccountController(accountUsecase inputport.Account) account.ServerInterface {
	return &AccountController{
		accountUsecase: accountUsecase,
	}
}

func (accountCtrl AccountController) PostAccountCreate(ctx echo.Context) error {
	req := new(account.AccountCreate)
	if err := ctx.Bind(req); err != nil {
		return presenter.ErrorResponse(ctx, http.StatusBadRequest, "不正なパラメータです")
	}

	account, err := accountCtrl.accountUsecase.CreateAccount(entity.Account{
		Name:     req.Name,
		Email:    string(req.Email),
		Password: req.Password,
	})

	if err != nil {
		return presenter.ErrorResponse(ctx, http.StatusBadRequest, "不正なパラメータです")
	}

	return ctx.JSON(http.StatusOK, presenter.GetAccount(account))
}

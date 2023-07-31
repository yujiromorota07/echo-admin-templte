package mainapi

import (
	"echo-admin-templte/entity"
	"echo-admin-templte/generated/login"
	"echo-admin-templte/presenter"
	"echo-admin-templte/usecase/inputport"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginController struct {
	loginUsecase inputport.Login
}

func NewLoginController(loginUsecase inputport.Login) login.ServerInterface {
	return &LoginController{
		loginUsecase: loginUsecase,
	}
}

func (ctrl LoginController) PostLogin(ctx echo.Context) error {
	req := new(login.Login)
	if err := ctx.Bind(req); err != nil {
		return presenter.ErrorResponse(ctx, http.StatusUnprocessableEntity, "不正なパラメータです")
	}

	token, err := ctrl.loginUsecase.Login(entity.Login{
		Email:    string(req.Email),
		Password: req.Password,
	})

	if err != nil {
		return presenter.ErrorResponse(ctx, http.StatusUnprocessableEntity, "不正なパラメータです")
	}

	return ctx.JSON(http.StatusOK, token)
}

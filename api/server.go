package main

import (
	"echo-admin-templte/di"
	"echo-admin-templte/generated/account"
	"echo-admin-templte/generated/login"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	loginController := di.IntializeLoginController()
	loginAPI := e.Group("")
	login.RegisterHandlers(loginAPI, loginController)

	accountController := di.IntializeAccountController()
	accountAPI := e.Group("")
	account.RegisterHandlers(accountAPI, accountController)

	e.Logger.Fatal(e.Start(":8081"))
}

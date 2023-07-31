package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SuccessResponse -
func SuccessResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, errorResponse{
		Code:    http.StatusOK,
		Message: message,
	})
}

// ErrorResponse -
func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, errorResponse{
		Code:    code,
		Message: message,
	})
}

// ErrorResponse -
func AuthorityErrorResponse(c echo.Context) error {
	return c.JSON(http.StatusForbidden, errorResponse{
		Code:    http.StatusForbidden,
		Message: "アクセスする権限がありません",
	})
}

// ErrorResponse はエラーのレスポンス
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

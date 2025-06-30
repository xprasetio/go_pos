package middleware

import (
	"net/http"
	"pos-go/pkg/response"
	"pos-go/shared/constants"

	"github.com/labstack/echo/v4"
)

func AdminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		roleID, ok := c.Get("role_id").(int64)
		if !ok || roleID != 1 {
			return response.SendResponseHttp(c, http.StatusForbidden, constants.StatusForbidden, nil)
		}
		return next(c)
	}
}

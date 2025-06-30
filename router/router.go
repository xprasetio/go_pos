package router

import (
	user "pos-go/internal/user"
	"pos-go/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, userHandler *user.UserHandler, jwtMiddleware *middleware.JWTMiddleware) {

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	protected := e.Group("/api/v1")
	protected.Use(jwtMiddleware.Authenticate)

	// CategoryVilla endpoints (hanya admin)
	// adminOnly := middleware.AdminOnlyMiddleware
	// protected.DELETE("/category", catHandler.Delete, adminOnly)

	// TODO: Daftarkan route lain
}

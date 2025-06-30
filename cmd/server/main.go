package main

import (
	"net/http"
	"pos-go/config"
	"pos-go/container"
	user_domain "pos-go/internal/user"
	"pos-go/pkg/db"
	"pos-go/pkg/logger"
	"pos-go/pkg/middleware"
	"pos-go/router"

	"github.com/labstack/echo/v4"
)

func main() {
	// Inisialisasi logger
	logger.SetupLogger()

	// Inisialisasi konfigurasi
	config.LoadConfig()

	// Inisialisasi koneksi database
	db.InitDB()

	// Automigrate tabel users
	db.AutoMigrate(db.DB)

	// Setup DI container
	ctn := container.BuildContainer()
	defer ctn.Delete()

	userHandler := ctn.Get(container.UserHandlerDefName).(*user_domain.UserHandler)
	jwtMiddleware := ctn.Get(container.JWTMiddlewareDef).(*middleware.JWTMiddleware)

	e := echo.New()

	// Pass handler ke router
	router.SetupRouter(e, userHandler, jwtMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Aplikasi booking siap!"})
	})

	e.Logger.Fatal(e.Start(":8081"))
}

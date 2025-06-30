package container

import (
	"database/sql"
	user_domain "pos-go/internal/user"
	"pos-go/pkg/db"
	"pos-go/pkg/jwt"
	"pos-go/pkg/middleware"

	di "github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
)

func BuildContainer() di.Container {
	builder, _ := di.NewBuilder()

	// Register DB
	_ = builder.Add(di.Def{
		Name:  DBDefName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return db.DB, nil
		},
	})

	// Register UserRepository
	_ = builder.Add(di.Def{
		Name:  UserRepoDefName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			dbConn := ctn.Get(DBDefName).(*sql.DB)
			return user_domain.NewUserRepository(dbConn), nil
		},
	})

	// Register JWTService
	_ = builder.Add(di.Def{
		Name:  JWTServiceDefName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			secret := viper.GetString("app.jwt_secret")
			exp := viper.GetInt("app.jwt_exp")
			if exp == 0 {
				exp = 24 // default 24 jam
			}
			return jwt.NewJWTService(secret, exp), nil
		},
	})

	// Register UserService
	_ = builder.Add(di.Def{
		Name:  UserServiceDefName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			repo := ctn.Get(UserRepoDefName).(user_domain.UserRepository)
			jwtSvc := ctn.Get(JWTServiceDefName).(jwt.JWTService)
			return user_domain.NewUserService(repo, jwtSvc), nil
		},
	})

	// Register UserHandler
	_ = builder.Add(di.Def{
		Name:  UserHandlerDefName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			userService := ctn.Get(UserServiceDefName).(user_domain.UserService)
			return user_domain.NewUserHandler(userService), nil
		},
	})

	// JWTMidlleware
	_ = builder.Add(di.Def{
		Name:  JWTMiddlewareDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			jwtSvc := ctn.Get(JWTServiceDefName).(jwt.JWTService)
			userService := ctn.Get(UserServiceDefName).(user_domain.UserService)
			return middleware.NewJWTMiddleware(jwtSvc, userService), nil
		},
	})

	return builder.Build()
}

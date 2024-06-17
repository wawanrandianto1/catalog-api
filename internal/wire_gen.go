// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"catalog-be/internal/config"
	"catalog-be/internal/middlewares"
	"catalog-be/internal/modules/auth"
	"catalog-be/internal/modules/circle"
	"catalog-be/internal/modules/circle/bookmark"
	"catalog-be/internal/modules/circle/circle_fandom"
	"catalog-be/internal/modules/circle/circle_work_type"
	"catalog-be/internal/modules/event"
	"catalog-be/internal/modules/fandom"
	"catalog-be/internal/modules/product"
	"catalog-be/internal/modules/refresh_token"
	"catalog-be/internal/modules/user"
	"catalog-be/internal/modules/work_type"
	"catalog-be/internal/router"
	"catalog-be/internal/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeServer(db *gorm.DB, validator2 *validator.Validate) *router.HTTP {
	userRepo := user.NewUserRepo(db)
	userService := user.NewUserService(userRepo)
	config := internal_config.NewConfig()
	refreshTokenRepo := refreshtoken.NewRefreshTokenRepo(db)
	utilsUtils := utils.NewUtils()
	refreshTokenService := refreshtoken.NewRefreshTokenService(refreshTokenRepo, utilsUtils)
	circleRepo := circle.NewCircleRepo(db)
	circleWorkTypeRepo := circle_work_type.NewCircleWorkTypeRepo(db)
	circleWorkTypeService := circle_work_type.NewCircleWorkTypeService(circleWorkTypeRepo)
	circleFandomRepo := circle_fandom.NewCircleFandomRepo(db)
	circleFandomService := circle_fandom.NewCircleFandomService(circleFandomRepo)
	circleBookmarkRepo := bookmark.NewCircleBookmarkRepo(db)
	circleBookmarkService := bookmark.NewCircleBookmarkService(circleBookmarkRepo)
	productRepo := product.NewProductRepo(db)
	productService := product.NewProductService(productRepo)
	circleService := circle.NewCircleService(circleRepo, userService, utilsUtils, refreshTokenService, circleWorkTypeService, circleFandomService, circleBookmarkService, productService)
	authService := auth.NewAuthService(userService, config, refreshTokenService, utilsUtils, circleService)
	authHandler := auth.NewAuthHandler(authService, validator2)
	authMiddleware := middlewares.NewAuthMiddleware(userService)
	fandomRepo := fandom.NewFandomRepo(db)
	fandomService := fandom.NewFandomService(fandomRepo)
	fandomHandler := fandom.NewFandomHandler(fandomService, validator2)
	workTypeRepo := work_type.NewWorkTypeRepo(db)
	workTypeService := work_type.NewWorkTypeService(workTypeRepo)
	workTypeHandler := work_type.NewWorkTypeHandler(workTypeService)
	circleHandler := circle.NewCircleHandler(circleService, validator2, userService, circleBookmarkService)
	eventRepo := event.NewEventRepo(db)
	eventService := event.NewEventService(eventRepo, utilsUtils)
	eventHandler := event.NewEventHandler(eventService, validator2)
	http := router.NewHTTP(authHandler, authMiddleware, fandomHandler, workTypeHandler, circleHandler, eventHandler)
	return http
}

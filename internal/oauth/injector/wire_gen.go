// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package oauth

import (
	"edtech.id/internal/admin/repository"
	admin2 "edtech.id/internal/admin/usecase"
	"edtech.id/internal/oauth/delivery/http"
	oauth2 "edtech.id/internal/oauth/repository"
	oauth3 "edtech.id/internal/oauth/usecase"
	"edtech.id/internal/user/repository"
	user2 "edtech.id/internal/user/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *oauth.OauthHandler {
	oauthClientRepository := oauth2.NewOauthClientRepository(db)
	oauthAccessTokenRepository := oauth2.NewOauthAccessTokenRepository(db)
	oauthRefreshTokenRepository := oauth2.NewOauthRefreshTokenRepository(db)
	userRepository := user.NewUserRepository(db)
	userUseCase := user2.NewUserUseCase(userRepository)
	adminRepository := admin.NewAdminRepository(db)
	adminUseCase := admin2.NewAdminUseCase(adminRepository)
	oauthUseCase := oauth3.NewOauthUseCase(oauthClientRepository, oauthAccessTokenRepository, oauthRefreshTokenRepository, userUseCase, adminUseCase)
	oauthHandler := oauth.NewOauthHandler(oauthUseCase)
	return oauthHandler
}

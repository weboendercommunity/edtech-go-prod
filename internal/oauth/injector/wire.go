//go:build wireinject
// +build wireinject

package oauth

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	oauthHandler "edtech.id/internal/oauth/delivery/http"
	oauthRepository "edtech.id/internal/oauth/repository"
	oauthUseCase "edtech.id/internal/oauth/usecase"
	userRepository "edtech.id/internal/user/repository"
	userUseCase "edtech.id/internal/user/usecase"
)

func InitializedService(db *gorm.DB) *oauthHandler.OauthHandler {
	wire.Build(
		oauthHandler.NewOauthHandler,
		oauthUseCase.NewOauthUseCase,
		oauthRepository.NewOauthClientRepository,
		oauthRepository.NewOauthAccessTokenRepository,
		oauthRepository.NewOauthRefreshTokenRepository,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &oauthHandler.OauthHandler{}
}

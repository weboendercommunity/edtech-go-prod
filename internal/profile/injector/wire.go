//go:build wireinject
// +build wireinject

package profile

import (
	profileHandler "edtech.id/internal/profile/delivery/http"
	profileUseCase "edtech.id/internal/profile/usecase"
	userRepository "edtech.id/internal/user/repository"
	userUseCase "edtech.id/internal/user/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *profileHandler.ProfileHandler {
	wire.Build(
		profileHandler.NewProfileHandler,
		profileUseCase.NewProfileUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &profileHandler.ProfileHandler{}
}

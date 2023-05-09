//go:build wireinject
// +build wireinject

package register

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	registerHandler "edtech.id/internal/register/delivery/http"
	registerUseCase "edtech.id/internal/register/usecase"
	userRepository "edtech.id/internal/user/repository"
	userUseCase "edtech.id/internal/user/usecase"
	mail "edtech.id/pkg/mail/gomail"
)

func InitializedService(db *gorm.DB) *registerHandler.RegisterHandler {
	wire.Build(
		registerHandler.NewRegisterHandler,
		registerUseCase.NewRegisterUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
		mail.NewSmtpMail,
	)

	return &registerHandler.RegisterHandler{}
}

package main

import (
	"github.com/gin-gonic/gin"

	mysql "edtech.id/pkg/db/mysql"

	userRepository "edtech.id/internal/user/repository"
	userUseCase "edtech.id/internal/user/usecase"

	registerHandler "edtech.id/internal/register/delivery/http"
	registerUseCase "edtech.id/internal/register/usecase"
	mail "edtech.id/pkg/mail/sendgrid"
)

func main() {
	db := mysql.DB()

	r := gin.Default()

	mail := mail.NewMail()

	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)

	registerUseCase := registerUseCase.NewRegisterUseCase(userUseCase, mail)

	registerHandler.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)
	r.Run("127.0.0.1:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

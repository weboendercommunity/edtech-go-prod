package utils

import (
	"math/rand"

	"github.com/gin-gonic/gin"

	oauthDto "edtech.id/internal/oauth/dto"
)

func RandString(number int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, number)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func GetCurrentUser(ctx *gin.Context) *oauthDto.MapClaimsResponseBody {
	user, _ := ctx.Get("user")

	return user.(*oauthDto.MapClaimsResponseBody)
}

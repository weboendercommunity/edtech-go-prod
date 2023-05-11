package utils

import (
	"math/rand"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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

func Paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset

		// if offset below 0, set to 1
		if page <= 0 {
			page = 1
		}

		pageSize := limit

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset = (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}

func GetFileName(fileUrl string) string {
	fileName := filepath.Base(fileUrl)

	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func IsVideo(fileUrl string) bool {
	return filepath.Ext(fileUrl) == ".mp4"
}

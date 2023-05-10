package middleware

import (
	"edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthAdmin(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	if !user.IsAdmin {
		ctx.JSON(403, utils.Response(403, "Forbidden", "You don't have access"))
		ctx.Abort()
		return
	}

	ctx.Next()
}

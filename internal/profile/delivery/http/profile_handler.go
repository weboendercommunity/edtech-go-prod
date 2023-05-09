package profile

import (
	"net/http"

	middleware "edtech.id/internal/middleware"
	profileUseCase "edtech.id/internal/profile/usecase"
	utils "edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileUseCase profileUseCase.ProfileUseCase
}

func NewProfileHandler(profileUseCase profileUseCase.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{profileUseCase}
}

func (handler *ProfileHandler) Route(r *gin.RouterGroup) {
	authorized := r.Group("/api")

	authorized.Use(middleware.AuthJwt)
	{
		authorized.GET("/profile", handler.GetProfile)
	}
}

func (handler *ProfileHandler) GetProfile(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	// TODO: Get Profile

	data, err := handler.profileUseCase.GetProfile(user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", nil))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", data))
}

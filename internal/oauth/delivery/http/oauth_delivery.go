package oauth

import (
	"net/http"

	oauthDto "edtech.id/internal/oauth/dto"
	oauthUseCase "edtech.id/internal/oauth/usecase"
	utils "edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type OauthHandler struct {
	oauthUseCase oauthUseCase.OauthUseCase
}

func NewOauthHandler(oauthUseCase oauthUseCase.OauthUseCase) *OauthHandler {
	return &OauthHandler{oauthUseCase}
}

func (handler *OauthHandler) Route(r *gin.RouterGroup) {

	oauthRouter := r.Group("/api")

	oauthRouter.POST("/oauth/login", handler.Login)
}

func (oh *OauthHandler) Login(ctx *gin.Context) {
	var input oauthDto.LoginRequestBody

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(400, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	// calls user data

	dataUser, err := oh.oauthUseCase.Login(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(500, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", dataUser))
}

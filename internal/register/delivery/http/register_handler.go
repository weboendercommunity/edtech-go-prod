package register

import (
	"net/http"

	"github.com/gin-gonic/gin"

	regiterUseCase "edtech.id/internal/register/usecase"
	userDto "edtech.id/internal/user/dto"
	"edtech.id/pkg/utils"
)

type RegisterHandler struct {
	registerUseCase regiterUseCase.RegisterUseCase
}

func NewRegisterHandler(registerUseCase regiterUseCase.RegisterUseCase) *RegisterHandler {
	return &RegisterHandler{registerUseCase}
}

func (rh *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/register", rh.Register)
}

func (rh *RegisterHandler) Register(ctx *gin.Context) {

	// validate input
	var registerRequestInput userDto.UserRequestBody

	err := ctx.ShouldBindJSON(&registerRequestInput)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(400, "Bad Request", err.Error()))
		ctx.Abort()
		return
	}

	inputErr := rh.registerUseCase.Register(registerRequestInput)

	if inputErr != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(500, "Internal Server Error", inputErr.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(201, "created", "Success please check your email"))

}

package order

import (
	"net/http"

	"edtech.id/internal/middleware"
	orderDto "edtech.id/internal/order/dto"
	orderUsecase "edtech.id/internal/order/usecase"
	"edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUsecase orderUsecase.OrderUsecase
}

func (oh *OrderHandler) Create(ctx *gin.Context) {
	var input orderDto.OrderRequestBody

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UserID = user.ID
	input.Email = user.Email

	data, err := oh.orderUsecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(200, "Success", data))
}

func (oh *OrderHandler) Route(r *gin.RouterGroup) {
	orderRouter := r.Group("/api/order")

	orderRouter.Use(middleware.AuthJwt)
	{
		orderRouter.POST("/", oh.Create)
	}
}

func NewOrderHandler(orderUsecase orderUsecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase}
}

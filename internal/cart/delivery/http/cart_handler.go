package handler

import (
	"net/http"
	"strconv"

	cartDto "edtech.id/internal/cart/dto"
	cartUsecase "edtech.id/internal/cart/usecase"
	"edtech.id/internal/middleware"
	"edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUsecase cartUsecase.CartUsecase
}

func (handler *CartHandler) Route(r *gin.RouterGroup) {
	cartRouter := r.Group("/api/cart")

	cartRouter.Use(middleware.AuthJwt)
	{
		cartRouter.GET("/", handler.FindById)
		cartRouter.POST("/", handler.Create)
		cartRouter.DELETE("/:id", handler.Delete)
	}
}

func (cartHandler *CartHandler) FindById(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := cartHandler.cartUsecase.FindByUserId(int(user.ID), offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(200, "Success", data))
}

func (cartHandler *CartHandler) Create(ctx *gin.Context) {
	var input cartDto.CartRequestBody

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UserID = user.ID

	data, err := cartHandler.cartUsecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusOK, "Created", data))
}

func (cartHandler *CartHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := utils.GetCurrentUser(ctx)

	err := cartHandler.cartUsecase.Delete(id, int(user.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", "Success"))
}

func NewCartHandler(cartUsecase cartUsecase.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase}
}

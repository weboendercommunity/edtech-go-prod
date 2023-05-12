package discount

import (
	"net/http"
	"strconv"

	discountDto "edtech.id/internal/discount/dto"
	discountUsecase "edtech.id/internal/discount/usecase"
	"edtech.id/internal/middleware"
	"edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type DiscountHandler struct {
	discountUsecase discountUsecase.DiscountUsecase
}

func (discountHandler *DiscountHandler) Route(r *gin.RouterGroup) {
	discountRouter := r.Group("/api/discounts")
	discountRouter.GET("/", discountHandler.FindAll)
	discountRouter.GET("/:id", discountHandler.FindById)
	discountRouter.GET("/code/:code", discountHandler.FindByCode)

	discountRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		discountRouter.POST("/", discountHandler.Create)
		discountRouter.PATCH("/:id", discountHandler.Update)
		discountRouter.DELETE("/:id", discountHandler.Delete)
	}
}

func (discountHandler *DiscountHandler) Create(ctx *gin.Context) {
	var input discountDto.DiscountRequestBody

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = user.ID

	discounts, err := discountHandler.discountUsecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", discounts))
}

func (discountHandler *DiscountHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	discounts := discountHandler.discountUsecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", discounts))
}

func (discountHandler *DiscountHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	discount, err := discountHandler.discountUsecase.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", discount))
}

func (discountHandler *DiscountHandler) FindByCode(ctx *gin.Context) {
	code := ctx.Param("code")

	discount, err := discountHandler.discountUsecase.FindByCode(code)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", discount))
}

func (discountHandler *DiscountHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input discountDto.DiscountRequestBody

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UpdatedBy = user.ID

	discount, err := discountHandler.discountUsecase.Update(id, input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", discount))
}

func (discountHandler *DiscountHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := discountHandler.discountUsecase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", nil))
}

func NewDiscountHandler(discountUsecase discountUsecase.DiscountUsecase) *DiscountHandler {
	return &DiscountHandler{discountUsecase}
}

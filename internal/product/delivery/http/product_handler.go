package product

import (
	"net/http"
	"strconv"

	"edtech.id/internal/middleware"
	productDto "edtech.id/internal/product/dto"
	productUsecase "edtech.id/internal/product/usecase"
	"edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase productUsecase.ProductUseCase
}

func NewProductHandler(
	productUsecase productUsecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUsecase}
}

func (ph *ProductHandler) Route(r *gin.RouterGroup) {
	productRouter := r.Group("/api/products")

	productRouter.GET("", ph.FindAll)
	productRouter.GET("/:id", ph.FindById)
	productRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		productRouter.POST("", ph.Create)
		productRouter.PATCH("/:id", ph.Update)
		productRouter.DELETE("/:id", ph.Delete)
	}
}

func (ph *ProductHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := ph.productUsecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", data))
}

func (ph *ProductHandler) FindById(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := ph.productUsecase.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "Not Found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", data))
}

func (ph *ProductHandler) Create(ctx *gin.Context) {
	var input productDto.ProductRequestBody

	requestErr := ctx.ShouldBind(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = user.ID

	newProduct, err := ph.productUsecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", newProduct))
}

func (ph *ProductHandler) Update(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	var input productDto.ProductRequestBody

	requestErr := ctx.ShouldBind(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UpdatedBy = user.ID

	updatedProduct, err := ph.productUsecase.Update(id, input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", updatedProduct))
}

func (ph *ProductHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ph.productUsecase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", nil))
}

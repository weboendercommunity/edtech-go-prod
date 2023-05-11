package product_category

import (
	"net/http"
	"strconv"

	"edtech.id/internal/middleware"
	productCategoryDto "edtech.id/internal/product_category/dto"
	productCategoryUsecase "edtech.id/internal/product_category/usecase"
	"edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler struct {
	productCategoryUsecase productCategoryUsecase.ProductCategoryUseCase
}

func NewProductCategoryHandler(productCategoryUsecase productCategoryUsecase.ProductCategoryUseCase) *ProductCategoryHandler {
	return &ProductCategoryHandler{productCategoryUsecase}
}

func (productCategoryHandler *ProductCategoryHandler) Route(r *gin.RouterGroup) {
	productCategoryRouter := r.Group("/api")

	productCategoryRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		productCategoryRouter.POST("/product-categories", productCategoryHandler.Create)
		productCategoryRouter.PATCH("/product-categories/:id", productCategoryHandler.Update)
		productCategoryRouter.DELETE("/product-categories/:id", productCategoryHandler.Delete)
	}
	r.GET("/api/product-categories", productCategoryHandler.FindAll)
	r.GET("/api/product-categories/:id", productCategoryHandler.FindById)

}

func (productCategoryHandler *ProductCategoryHandler) Create(ctx *gin.Context) {
	var input productCategoryDto.ProductCategoryRequestBody

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = user.ID

	productCategory, err := productCategoryHandler.productCategoryUsecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", productCategory))
}

func (productCategoryHandler *ProductCategoryHandler) Update(ctx *gin.Context) {
	var input productCategoryDto.ProductCategoryRequestBody
	id, _ := strconv.Atoi(ctx.Param("id"))

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UpdatedBy = user.ID

	productCategory, err := productCategoryHandler.productCategoryUsecase.Update(id, input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", productCategory))
}

func (productCategoryHandler *ProductCategoryHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	productCategories := productCategoryHandler.productCategoryUsecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", productCategories))
}

func (productCategoryHandler *ProductCategoryHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	productCategory, err := productCategoryHandler.productCategoryUsecase.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusNotFound, "Not Found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", productCategory))
}

func (productCategoryHandler *ProductCategoryHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := productCategoryHandler.productCategoryUsecase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", nil))
}

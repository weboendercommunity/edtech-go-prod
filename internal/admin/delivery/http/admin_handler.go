package admin

import (
	"net/http"
	"strconv"

	adminDto "edtech.id/internal/admin/dto"
	adminUseCase "edtech.id/internal/admin/usecase"
	"edtech.id/internal/middleware"
	"edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase adminUseCase.AdminUseCase
}

func NewAdminHandler(adminUseCase adminUseCase.AdminUseCase) *AdminHandler {
	return &AdminHandler{adminUseCase}
}

func (adminHanlder *AdminHandler) Route(r *gin.RouterGroup) {
	adminRouter := r.Group("/api/admin")

	adminRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		adminRouter.GET("/", adminHanlder.FindAll)
		adminRouter.GET("/:id", adminHanlder.FindById)
		adminRouter.POST("/", adminHanlder.Create)
		adminRouter.PATCH("/:id", adminHanlder.Update)
		adminRouter.DELETE("/:id", adminHanlder.Delete)
	}
}

func (adminHanlder *AdminHandler) Create(ctx *gin.Context) {
	var input adminDto.AdminRequestBody

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = user.ID

	// Create Data

	_, err := adminHanlder.adminUseCase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusOK, "Created", "Success"))
}

func (adminHandler *AdminHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input adminDto.AdminRequestBody

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UpdatedBy = user.ID

	updatedUser, err := adminHandler.adminUseCase.Update(id, input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", updatedUser))

}

func (adminHandler *AdminHandler) FindAll(ctx *gin.Context) {
	// api/admin?offset=0&limit=10

	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	data := adminHandler.adminUseCase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", data))
}

func (adminHandler *AdminHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := adminHandler.adminUseCase.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "Not Found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", data))
}

func (adminHandler *AdminHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := adminHandler.adminUseCase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "Not Found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Success", "Admin Deleted"))
}

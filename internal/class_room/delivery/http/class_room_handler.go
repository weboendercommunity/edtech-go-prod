package class_room

import (
	"net/http"
	"strconv"

	classRoomUsecase "edtech.id/internal/class_room/usecase"
	"edtech.id/internal/middleware"
	"edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ClassRoomHandler struct {
	classRoomUsecase classRoomUsecase.ClassRoomUsecase
}

func (crh *ClassRoomHandler) Route(r *gin.RouterGroup) {
	classRoomRouter := r.Group("/api/classroom")

	classRoomRouter.Use(middleware.AuthJwt)
	{
		classRoomRouter.GET("/", crh.FindAllByUserID)
	}
}

func NewClassRoomHandler(classRoomUsecase classRoomUsecase.ClassRoomUsecase) *ClassRoomHandler {
	return &ClassRoomHandler{classRoomUsecase}
}

func (crh *ClassRoomHandler) FindAllByUserID(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	classRooms, err := crh.classRoomUsecase.FindAllByUserID(offset, limit, int(user.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(200, utils.Response(http.StatusOK, "Success", classRooms))
}

package webhook

import (
	"net/http"

	webhookDto "edtech.id/internal/webhook/dto"
	webhookUsecase "edtech.id/internal/webhook/usecase"
	utils "edtech.id/pkg/utils"
	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	webhookUsecase webhookUsecase.WebhookUsecase
}

func NewWebhookHandler(webhookUsecase webhookUsecase.WebhookUsecase) *WebhookHandler {
	return &WebhookHandler{webhookUsecase}
}

func (webhookHandler *WebhookHandler) Route(ctx *gin.RouterGroup) {
	webhookRouter := ctx.Group("/api/webhook")
	{
		webhookRouter.POST("/xendit", webhookHandler.Xendit)
	}
}

func (webhookHandler *WebhookHandler) Xendit(ctx *gin.Context) {
	var input webhookDto.WebhookRequestBody

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	paymentErr := webhookHandler.webhookUsecase.UpdatePayment(input.ID)

	if paymentErr != nil {
		ctx.JSON(http.StatusInternalServerError,
			utils.Response(http.StatusInternalServerError,
				"internal server error", paymentErr.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", paymentErr))

}

//go:build wireinject
// +build wireinject

package webhook

import (
	cartRepository "edtech.id/internal/cart/repository"
	cartUseCase "edtech.id/internal/cart/usecase"
	classRoomRepository "edtech.id/internal/class_room/repository"
	classRoomUsecase "edtech.id/internal/class_room/usecase"
	discountRepository "edtech.id/internal/discount/repository"
	discountUseCase "edtech.id/internal/discount/usecase"
	orderRepository "edtech.id/internal/order/repository"
	orderUsecase "edtech.id/internal/order/usecase"
	orderDetailRepository "edtech.id/internal/order_detail/repository"
	orderDetailUseCase "edtech.id/internal/order_detail/usecase"
	paymentUsecase "edtech.id/internal/payment/usecase"
	productRepository "edtech.id/internal/product/repository"
	productUsecase "edtech.id/internal/product/usecase"
	webhookHandler "edtech.id/internal/webhook/delivery/http"
	webhookUseCase "edtech.id/internal/webhook/usecase"
	fileUpload "edtech.id/pkg/fileupload/cloudinary"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *webhookHandler.WebhookHandler {
	wire.Build(
		webhookHandler.NewWebhookHandler,
		webhookUseCase.NewWebhookUsecase,
		classRoomUsecase.NewClassRoomUsecase,
		classRoomRepository.NewClassRoomRepository,
		orderUsecase.NewOrderUsecase,
		orderRepository.NewOrderRepository,
		orderDetailUseCase.NewOrderDetailUsecase,
		orderDetailRepository.NewOrderDetailRepository,
		cartUseCase.NewCartUsecase,
		cartRepository.NewCartRepository,
		discountUseCase.NewDiscountUsecase,
		discountRepository.NewDiscountRepository,
		paymentUsecase.NewPaymentUsecase,
		productUsecase.NewProductUsecase,
		productRepository.NewProductRepository,
		fileUpload.NewFileUpload,
	)

	return &webhookHandler.WebhookHandler{}
}

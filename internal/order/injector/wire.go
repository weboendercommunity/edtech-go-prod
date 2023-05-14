//go: build wireinject
//go:build wireinject
// +build wireinject

package order

import (
	cartRepository "edtech.id/internal/cart/repository"
	cartUseCase "edtech.id/internal/cart/usecase"
	discountRepository "edtech.id/internal/discount/repository"
	discountUseCase "edtech.id/internal/discount/usecase"
	orderHandler "edtech.id/internal/order/delivery/http"
	orderRepository "edtech.id/internal/order/repository"
	orderUseCase "edtech.id/internal/order/usecase"
	orderDetailRepository "edtech.id/internal/order_detail/repository"
	orderDetailUseCase "edtech.id/internal/order_detail/usecase"
	paymentUsecase "edtech.id/internal/payment/usecase"
	productRepository "edtech.id/internal/product/repository"
	productUsecase "edtech.id/internal/product/usecase"
	fileUpload "edtech.id/pkg/fileupload/cloudinary"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *orderHandler.OrderHandler {
	wire.Build(
		orderHandler.NewOrderHandler,
		orderUseCase.NewOrderUsecase,
		orderRepository.NewOrderRepository,
		orderDetailUseCase.NewOrderDetailUsecase,
		orderDetailRepository.NewOrderDetailRepository,
		cartUseCase.NewCartUsecase,
		cartRepository.NewCartRepository,
		discountUseCase.NewDiscountUsecase,
		discountRepository.NewDiscountRepository,
		paymentUsecase.NewPaymentUsecase,
		productUsecase.NewProductUseCase,
		productRepository.NewProductRepository,
		fileUpload.NewFileUpload,
	)

	return &orderHandler.OrderHandler{}
}

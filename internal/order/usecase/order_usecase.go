package order

import (
	"errors"
	"strconv"

	cartUsecase "edtech.id/internal/cart/usecase"
	discountEntity "edtech.id/internal/discount/entity"
	discountUsecase "edtech.id/internal/discount/usecase"
	orderDto "edtech.id/internal/order/dto"
	orderEntity "edtech.id/internal/order/entity"
	orderRepository "edtech.id/internal/order/repository"
	orderDetailDto "edtech.id/internal/order_detail/dto"
	orderDetailUsecase "edtech.id/internal/order_detail/usecase"
	paymentDto "edtech.id/internal/payment/dto"
	paymentUsecase "edtech.id/internal/payment/usecase"
	productEntity "edtech.id/internal/product/entity"
	productUsecase "edtech.id/internal/product/usecase"
	"github.com/google/uuid"
)

type OrderUsecase interface {
	Create(orderDto orderDto.OrderRequestBody) (*orderEntity.Order, error)
	FindAll(offset int, limit int) []orderEntity.Order
	FindById(id int) (*orderEntity.Order, error)
}

type OrderUsecaseImpl struct {
	orderRepository    orderRepository.OrderRepository
	cartUsecase        cartUsecase.CartUsecase
	discountUsecase    discountUsecase.DiscountUsecase
	productUsecase     productUsecase.ProductUseCase
	orderDetailUsecase orderDetailUsecase.OrderDetailUsecase
	paymentUsecase     paymentUsecase.PaymentUsecase
}

// Create implements OrderUsecase
func (ou *OrderUsecaseImpl) Create(orderDto orderDto.OrderRequestBody) (*orderEntity.Order, error) {
	price := 0
	totalPrice := 0
	description := ""

	var products []productEntity.Product

	order := orderEntity.Order{
		UserID:      orderDto.UserID,
		Status:      "pending",
		CreatedByID: &orderDto.UserID,
	}

	var discount *discountEntity.Discount

	carts := ou.cartUsecase.FindByUserId(int(orderDto.UserID), 0, 99999)

	// check carts is empty
	if len(carts) == 0 {
		if orderDto.ProductID == nil {
			return nil, errors.New("cart is empty")
		}
	}

	// check discount
	if orderDto.DiscountCode != nil {
		dataDiscount, err := ou.discountUsecase.FindByCode(*orderDto.DiscountCode)

		if err != nil {
			return nil, errors.New("discount not found")
		}

		discount = dataDiscount

		if discount.RemainingQuantity == 0 {
			return nil, errors.New("discount is expired")
		}

		// TODO: check if discount date is expired
	}

	// handle if cart is not empty
	if len(carts) > 0 {
		for _, cart := range carts {
			product, err := ou.productUsecase.FindById(int(cart.ProductID))

			if err != nil {
				return nil, err
			}

			products = append(products, *product)
		}
	} else if orderDto.ProductID != nil {
		// handle if user not have cart and order product directly
		product, err := ou.productUsecase.FindById(int(*orderDto.ProductID))

		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	// calculate price
	for index, product := range products {
		price += int(product.Price)

		i := strconv.Itoa(index + 1)

		description += i + ". Product: " + product.Title + "\n"
	}

	// assign total price
	totalPrice = price

	// apply discount
	if discount != nil {
		if discount.Type == discountEntity.Fixed {
			totalPrice = price - discount.Value
		} else if discount.Type == discountEntity.Precentage {
			totalPrice = price - (price * discount.Value / 100)
		}

		order.DiscountID = &discount.ID
	}

	order.Price = int64(price)           // actual price
	order.TotalPrice = int64(totalPrice) // total price after discount

	// set external id
	externalId := uuid.New().String()
	order.ExternalID = externalId

	// insert order
	createdOrder, err := ou.orderRepository.Create(order)

	if err != nil {
		return nil, err
	}

	// insert order detail
	for _, product := range products {
		orderDetail := orderDetailDto.OrderDetailRequestBody{
			OrderID:   createdOrder.ID,
			ProductID: product.ID,
			Price:     product.Price,
			CreatedBy: order.UserID,
		}

		_, err := ou.orderDetailUsecase.Create(orderDetail)

		if err != nil {
			return nil, err
		}
	}

	// handle with payment gateway
	dataPayment := paymentDto.PaymentRequestBody{
		ExternalID:  externalId,
		Amount:      createdOrder.TotalPrice,
		PayerEmail:  orderDto.Email,
		Description: description,
	}

	payment, err := ou.paymentUsecase.Create(dataPayment)

	if err != nil {
		return nil, err
	}

	createdOrder.CheckoutLink = payment.InvoiceURL

	updatedOrder, err := ou.orderRepository.Update(*createdOrder)

	if err != nil {
		return nil, err
	}

	// update remaining quantity discount
	if discount != nil {

		_, err := ou.discountUsecase.UpdateRemainingQuantity(int(discount.ID), 1, "-")

		if err != nil {
			return nil, err
		}
	}

	// delete cart
	deletedCart := ou.cartUsecase.DeleteByUserId(int(orderDto.UserID))

	if deletedCart != nil {
		return nil, errors.New("failed to delete cart")
	}

	return updatedOrder, nil

}

// FindAll implements OrderUsecase
func (ou *OrderUsecaseImpl) FindAll(offset int, limit int) []orderEntity.Order {
	return ou.orderRepository.FindAll(offset, limit)
}

// FindById implements OrderUsecase
func (ou *OrderUsecaseImpl) FindById(id int) (*orderEntity.Order, error) {
	return ou.orderRepository.FindById(id)
}

func NewOrderUsecase(
	orderRepository orderRepository.OrderRepository,
	cartUsecase cartUsecase.CartUsecase,
	discountUsecase discountUsecase.DiscountUsecase,
	productUsecase productUsecase.ProductUseCase,
	orderDetailUsecase orderDetailUsecase.OrderDetailUsecase,
	paymentUsecase paymentUsecase.PaymentUsecase,
) OrderUsecase {
	return &OrderUsecaseImpl{
		orderRepository,
		cartUsecase,
		discountUsecase,
		productUsecase,
		orderDetailUsecase,
		paymentUsecase}
}

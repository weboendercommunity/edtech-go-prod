package cart

import (
	"errors"

	cartDto "edtech.id/internal/cart/dto"
	cartEntity "edtech.id/internal/cart/entity"
	cartRepository "edtech.id/internal/cart/repository"
)

type CartUsecase interface {
	FindByUserId(userId int, offset int, limit int) []cartEntity.Cart
	FindById(id int) (*cartEntity.Cart, error)
	Create(cartDto cartDto.CartRequestBody) (*cartEntity.Cart, error)
	Delete(id int, userId int) error
}

type CartUsecaseImpl struct {
	cartRepository cartRepository.CartRepository
}

// Create implements CartUsecase
func (cu *CartUsecaseImpl) Create(cartDto cartDto.CartRequestBody) (*cartEntity.Cart, error) {
	cart := cartEntity.Cart{
		ProductID: cartDto.ProductID,
		UserID:    cartDto.UserID,
	}

	//TODO: validate if product exist

	createdCart, err := cu.cartRepository.Create(cart)

	if err != nil {
		return nil, err
	}

	return createdCart, nil
}

// Delete implements CartUsecase
func (cu *CartUsecaseImpl) Delete(id int, userId int) error {
	cart, err := cu.cartRepository.FindById(id)

	if err != nil {
		return err
	}

	if cart.UserID != int64(userId) {
		return errors.New("you are not authorized to delete this cart")
	}

	err = cu.cartRepository.Delete(*cart)

	if err != nil {
		return err
	}

	return nil
}

// FindById implements CartUsecase
func (cu *CartUsecaseImpl) FindById(id int) (*cartEntity.Cart, error) {
	return cu.cartRepository.FindById(id)
}

// FindByUserId implements CartUsecase
func (cu *CartUsecaseImpl) FindByUserId(userId int, offset int, limit int) []cartEntity.Cart {
	return cu.cartRepository.FindByUserId(userId, offset, limit)
}

func NewCartUsecase(cartRepository cartRepository.CartRepository) CartUsecase {
	return &CartUsecaseImpl{cartRepository}
}

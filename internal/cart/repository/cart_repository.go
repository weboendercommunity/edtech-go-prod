package cart

import (
	cartEntity "edtech.id/internal/cart/entity"
	"edtech.id/pkg/utils"
	"gorm.io/gorm"
)

type CartRepository interface {
	FindById(id int) (*cartEntity.Cart, error)
	FindByUserId(userId int, offset int, limit int) []cartEntity.Cart
	Create(cartEntity cartEntity.Cart) (*cartEntity.Cart, error)
	Delete(cartEntity cartEntity.Cart) error
}

type CartRepositoryImpl struct {
	db *gorm.DB
}

// Create implements CartRepository
func (cr *CartRepositoryImpl) Create(cartEntity cartEntity.Cart) (*cartEntity.Cart, error) {
	createdCart := cr.db.Create(&cartEntity)

	if createdCart.Error != nil {
		return nil, createdCart.Error
	}

	return &cartEntity, nil
}

// Delete implements CartRepository
func (cr *CartRepositoryImpl) Delete(cartEntity cartEntity.Cart) error {
	deletedCart := cr.db.Delete(&cartEntity)

	if deletedCart.Error != nil {
		return deletedCart.Error
	}

	return nil
}

// FindById implements CartRepository
func (cr *CartRepositoryImpl) FindById(id int) (*cartEntity.Cart, error) {
	var cart cartEntity.Cart

	dataCart := cr.db.Preload("User").Preload("Product").First(&cart, id)

	if dataCart.Error != nil {
		return nil, dataCart.Error
	}

	return &cart, nil
}

// FindByUserId implements CartRepository
func (cr *CartRepositoryImpl) FindByUserId(userId int, offset int, limit int) []cartEntity.Cart {
	var carts []cartEntity.Cart

	cr.db.Scopes(utils.Paginate(offset, limit)).Preload("User").Preload("Product").Where("user_id = ?", userId).Find(&carts)

	return carts
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &CartRepositoryImpl{db}
}

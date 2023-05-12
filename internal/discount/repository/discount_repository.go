package discount

import (
	discountEntity "edtech.id/internal/discount/entity"
	"edtech.id/pkg/utils"
	"gorm.io/gorm"
)

type DiscountRepository interface {
	FindAll(offset int, limit int) []discountEntity.Discount
	FindById(id int) (*discountEntity.Discount, error)
	FindByCode(code string) (*discountEntity.Discount, error)
	Create(discountEntity discountEntity.Discount) (*discountEntity.Discount, error)
	Update(discountEntity discountEntity.Discount) (*discountEntity.Discount, error)
	Delete(discountEntity discountEntity.Discount) error
}

type DiscountRepositoryImpl struct {
	db *gorm.DB
}

// Create implements DiscountRepository
func (dr *DiscountRepositoryImpl) Create(discountEntity discountEntity.Discount) (*discountEntity.Discount, error) {
	createdDiscount := dr.db.Create(&discountEntity)

	if createdDiscount.Error != nil {
		return nil, createdDiscount.Error
	}

	return &discountEntity, nil
}

// Delete implements DiscountRepository
func (dr *DiscountRepositoryImpl) Delete(discountEntity discountEntity.Discount) error {
	deletedDiscount := dr.db.Delete(&discountEntity)

	if deletedDiscount.Error != nil {
		return deletedDiscount.Error
	}

	return nil
}

// FindAll implements DiscountRepository
func (dr *DiscountRepositoryImpl) FindAll(offset int, limit int) []discountEntity.Discount {
	var discounts []discountEntity.Discount

	dr.db.Scopes(utils.Paginate(offset, limit)).Find(&discounts)

	return discounts
}

// FindByCode implements DiscountRepository
func (dr *DiscountRepositoryImpl) FindByCode(code string) (*discountEntity.Discount, error) {
	var discount discountEntity.Discount

	dataDiscount := dr.db.Where("code = ?", code).First(&discount)

	if dataDiscount.Error != nil {
		return nil, dataDiscount.Error
	}

	return &discount, nil
}

// FindById implements DiscountRepository
func (dr *DiscountRepositoryImpl) FindById(id int) (*discountEntity.Discount, error) {
	var discount discountEntity.Discount

	dataDiscount := dr.db.First(&discount, id)

	if dataDiscount.Error != nil {
		return nil, dataDiscount.Error
	}

	return &discount, nil
}

// Update implements DiscountRepository
func (dr *DiscountRepositoryImpl) Update(discountEntity discountEntity.Discount) (*discountEntity.Discount, error) {
	updatedDiscount := dr.db.Save(&discountEntity)

	if updatedDiscount.Error != nil {
		return nil, updatedDiscount.Error
	}

	return &discountEntity, nil
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &DiscountRepositoryImpl{db}
}

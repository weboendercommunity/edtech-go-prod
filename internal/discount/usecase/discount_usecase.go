package discount

import (
	"database/sql"
	"errors"

	discountDto "edtech.id/internal/discount/dto"
	discountEntity "edtech.id/internal/discount/entity"
	discountRepository "edtech.id/internal/discount/repository"
)

type DiscountUsecase interface {
	FindAll(offset int, limit int) []discountEntity.Discount
	FindById(id int) (*discountEntity.Discount, error)
	FindByCode(code string) (*discountEntity.Discount, error)
	Create(discountDto discountDto.DiscountRequestBody) (*discountEntity.Discount, error)
	Update(id int, discountDto discountDto.DiscountRequestBody) (*discountEntity.Discount, error)
	Delete(id int) error
	UpdateRemainingQuantity(id int, quantity int, operator string) (*discountEntity.Discount, error)
}

type DiscountUsecaseImpl struct {
	discountRepository discountRepository.DiscountRepository
}

// Create implements DiscountUsecase
func (du *DiscountUsecaseImpl) Create(discountDto discountDto.DiscountRequestBody) (*discountEntity.Discount, error) {
	discount := discountEntity.Discount{
		Name:              discountDto.Name,
		Code:              discountDto.Code,
		Quantity:          discountDto.Quantity,
		Value:             discountDto.Value,
		RemainingQuantity: discountDto.Quantity,
		Type:              discountDto.Type,
		StartDate: sql.NullTime{
			Time:  discountDto.StartDate,
			Valid: true,
		},
		EndDate: sql.NullTime{
			Time:  discountDto.EndDate,
			Valid: true,
		},
		CreatedByID: &discountDto.CreatedBy,
	}

	dataDiscount, err := du.discountRepository.Create(discount)

	if err != nil {
		return nil, err
	}

	return dataDiscount, nil
}

// Delete implements DiscountUsecase
func (du *DiscountUsecaseImpl) Delete(id int) error {
	// find discount by id

	discount, err := du.discountRepository.FindById(id)

	if err != nil {
		return err
	}

	// delete discount

	deletedDiscount := du.discountRepository.Delete(*discount)

	if deletedDiscount != nil {
		return deletedDiscount
	}

	return nil
}

// FindAll implements DiscountUsecase
func (du *DiscountUsecaseImpl) FindAll(offset int, limit int) []discountEntity.Discount {
	return du.discountRepository.FindAll(offset, limit)
}

// FindByCode implements DiscountUsecase
func (du *DiscountUsecaseImpl) FindByCode(code string) (*discountEntity.Discount, error) {
	return du.discountRepository.FindByCode(code)
}

// FindById implements DiscountUsecase
func (du *DiscountUsecaseImpl) FindById(id int) (*discountEntity.Discount, error) {
	return du.discountRepository.FindById(id)
}

// Update implements DiscountUsecase
func (du *DiscountUsecaseImpl) Update(id int, discountDto discountDto.DiscountRequestBody) (*discountEntity.Discount, error) {
	// find discount by id

	discount, err := du.discountRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	// update discount

	discount.Name = discountDto.Name
	discount.Code = discountDto.Code
	discount.Quantity = discountDto.Quantity
	discount.Value = discountDto.Value
	discount.RemainingQuantity = discountDto.Quantity
	discount.Type = discountDto.Type
	discount.UpdatedByID = &discountDto.UpdatedBy
	discount.StartDate.Time = discountDto.StartDate
	discount.EndDate.Time = discountDto.EndDate

	dataDiscount, err := du.discountRepository.Update(*discount)

	if err != nil {
		return nil, err
	}

	return dataDiscount, nil
}

// UpdateRemainingQuantity implements DiscountUsecase
func (du *DiscountUsecaseImpl) UpdateRemainingQuantity(id int, quantity int, operator string) (*discountEntity.Discount, error) {
	// find discount by id

	discount, err := du.discountRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	if operator == "+" {
		discount.RemainingQuantity += quantity
	} else if operator == "-" {
		discount.RemainingQuantity -= quantity
	} else {
		return nil, errors.New("operator not valid")
	}

	updateDiscount, err := du.discountRepository.Update(*discount)

	if err != nil {
		return nil, err
	}

	return updateDiscount, nil
}

func NewDiscountUsecase(discountRepository discountRepository.DiscountRepository) DiscountUsecase {
	return &DiscountUsecaseImpl{discountRepository}
}

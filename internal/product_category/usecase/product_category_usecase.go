package product_category

import (
	productCategoryDto "edtech.id/internal/product_category/dto"
	productCategoryEntity "edtech.id/internal/product_category/entity"
	productCategoryRepository "edtech.id/internal/product_category/repository"
)

type ProductCategoryUseCase interface {
	FindAll(offset int, limit int) []productCategoryEntity.ProductCategory
	FindById(id int) (*productCategoryEntity.ProductCategory, error)
	Create(productCategoryDto productCategoryDto.ProductCategoryRequestBody) (*productCategoryEntity.ProductCategory, error)
	Update(id int, productCategoryDto productCategoryDto.ProductCategoryRequestBody) (*productCategoryEntity.ProductCategory, error)
	Delete(id int) error
}

type ProductCategoryUseCaseImpl struct {
	productCategoryRepository productCategoryRepository.ProductCategoryRepository
}

// Create implements ProductCategoryUseCase
func (pcu *ProductCategoryUseCaseImpl) Create(productCategoryDto productCategoryDto.ProductCategoryRequestBody) (*productCategoryEntity.ProductCategory, error) {
	productCategory := productCategoryEntity.ProductCategory{
		Name:        productCategoryDto.Name,
		CreatedByID: productCategoryDto.CreatedBy,
	}

	createdProductCategory, err := pcu.productCategoryRepository.Create(productCategory)

	if err != nil {
		return nil, err
	}

	return createdProductCategory, nil
}

// Delete implements ProductCategoryUseCase
func (pcu *ProductCategoryUseCaseImpl) Delete(id int) error {
	productCategory, err := pcu.productCategoryRepository.FindById(id)

	if err != nil {
		return err
	}

	err = pcu.productCategoryRepository.Delete(*productCategory)

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductCategoryUseCase
func (pcu *ProductCategoryUseCaseImpl) FindAll(offset int, limit int) []productCategoryEntity.ProductCategory {
	return pcu.productCategoryRepository.FindAll(offset, limit)
}

// FindById implements ProductCategoryUseCase
func (pcu *ProductCategoryUseCaseImpl) FindById(id int) (*productCategoryEntity.ProductCategory, error) {
	return pcu.productCategoryRepository.FindById(id)
}

// Update implements ProductCategoryUseCase
func (pcu *ProductCategoryUseCaseImpl) Update(id int, productCategoryDto productCategoryDto.ProductCategoryRequestBody) (*productCategoryEntity.ProductCategory, error) {
	productCategory, err := pcu.productCategoryRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	productCategory.Name = productCategoryDto.Name
	productCategory.UpdatedByID = &productCategoryDto.UpdatedBy

	updateProductCategory, err := pcu.productCategoryRepository.Update(*productCategory)

	if err != nil {
		return nil, err
	}

	return updateProductCategory, nil
}

func NewProductCategoryUseCase(productCategoryRepository productCategoryRepository.ProductCategoryRepository) ProductCategoryUseCase {
	return &ProductCategoryUseCaseImpl{productCategoryRepository}
}

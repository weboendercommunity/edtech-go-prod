package product_category

import (
	productCategoryEntity "edtech.id/internal/product_category/entity"
	"edtech.id/pkg/utils"
	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	FindAll(offset int, limit int) []productCategoryEntity.ProductCategory
	FindById(id int) (*productCategoryEntity.ProductCategory, error)
	Create(productCategoryEntity productCategoryEntity.ProductCategory) (*productCategoryEntity.ProductCategory, error)
	Update(productCategoryEntity productCategoryEntity.ProductCategory) (*productCategoryEntity.ProductCategory, error)
	Delete(productCategoryEntity productCategoryEntity.ProductCategory) error
}

type ProductCategoryRepositoryImpl struct {
	db *gorm.DB
}

// Create implements ProductCategoryRepository
func (pcr *ProductCategoryRepositoryImpl) Create(productCategoryEntity productCategoryEntity.ProductCategory) (*productCategoryEntity.ProductCategory, error) {
	createdProductCategory := pcr.db.Create(&productCategoryEntity)

	if createdProductCategory.Error != nil {
		return nil, createdProductCategory.Error
	}

	return &productCategoryEntity, nil
}

// Delete implements ProductCategoryRepository
func (pcr *ProductCategoryRepositoryImpl) Delete(productCategoryEntity productCategoryEntity.ProductCategory) error {
	deletedProductCategory := pcr.db.Delete(&productCategoryEntity)

	if deletedProductCategory.Error != nil {
		return deletedProductCategory.Error
	}

	return nil
}

// FindAll implements ProductCategoryRepository
func (pcr *ProductCategoryRepositoryImpl) FindAll(offset int, limit int) []productCategoryEntity.ProductCategory {
	var productCategories []productCategoryEntity.ProductCategory

	pcr.db.Scopes(utils.Paginate(offset, limit)).Find(&productCategories)

	return productCategories
}

// FindById implements ProductCategoryRepository
func (pcr *ProductCategoryRepositoryImpl) FindById(id int) (*productCategoryEntity.ProductCategory, error) {
	var productCategory productCategoryEntity.ProductCategory

	dataProductCategory := pcr.db.First(&productCategory, id)

	if dataProductCategory.Error != nil {
		return nil, dataProductCategory.Error
	}

	return &productCategory, nil
}

// Update implements ProductCategoryRepository
func (pcr *ProductCategoryRepositoryImpl) Update(productCategoryEntity productCategoryEntity.ProductCategory) (*productCategoryEntity.ProductCategory, error) {
	updatedProductCategory := pcr.db.Save(&productCategoryEntity)

	if updatedProductCategory.Error != nil {
		return nil, updatedProductCategory.Error
	}

	return &productCategoryEntity, nil
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &ProductCategoryRepositoryImpl{db}
}

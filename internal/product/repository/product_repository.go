package product

import (
	productEntity "edtech.id/internal/product/entity"
	"edtech.id/pkg/utils"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(offset int, limit int) []productEntity.Product
	FindById(id int) (*productEntity.Product, error)
	Create(productEntity productEntity.Product) (*productEntity.Product, error)
	Update(productEntity productEntity.Product) (*productEntity.Product, error)
	Delete(productEntity productEntity.Product) error
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

// Create implements ProductRepository
func (pr *ProductRepositoryImpl) Create(productEntity productEntity.Product) (*productEntity.Product, error) {
	createdProduct := pr.db.Create(&productEntity)

	if createdProduct.Error != nil {
		return nil, createdProduct.Error
	}

	return &productEntity, nil
}

// Delete implements ProductRepository
func (pr *ProductRepositoryImpl) Delete(productEntity productEntity.Product) error {
	deletedProduct := pr.db.Delete(&productEntity)

	if deletedProduct.Error != nil {
		return deletedProduct.Error
	}

	return nil
}

// FindAll implements ProductRepository
func (pr *ProductRepositoryImpl) FindAll(offset int, limit int) []productEntity.Product {
	var products []productEntity.Product

	pr.db.Scopes(utils.Paginate(offset, limit)).Preload("ProductCategory").Find(&products)

	return products
}

// FindById implements ProductRepository
func (pr *ProductRepositoryImpl) FindById(id int) (*productEntity.Product, error) {
	var product productEntity.Product

	dataProduct := pr.db.Preload("ProductCategory").First(&product, id)

	if dataProduct.Error != nil {
		return nil, dataProduct.Error
	}

	return &product, nil
}

// Update implements ProductRepository
func (pr *ProductRepositoryImpl) Update(productEntity productEntity.Product) (*productEntity.Product, error) {
	updateProduct := pr.db.Save(&productEntity)

	if updateProduct.Error != nil {
		return nil, updateProduct.Error
	}

	return &productEntity, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}

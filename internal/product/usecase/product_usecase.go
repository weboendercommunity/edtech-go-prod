package product

import (
	"fmt"

	productDto "edtech.id/internal/product/dto"
	productEntity "edtech.id/internal/product/entity"
	productRepository "edtech.id/internal/product/repository"
	fileUpload "edtech.id/pkg/fileupload/cloudinary"
)

type ProductUseCase interface {
	FindAll(offset int, limit int) []productEntity.Product
	FindById(id int) (*productEntity.Product, error)
	Create(productDto productDto.ProductRequestBody) (*productEntity.Product, error)
	Update(id int, productDto productDto.ProductRequestBody) (*productEntity.Product, error)
	Delete(id int) error
}

type ProductUseCaseImpl struct {
	productRepository productRepository.ProductRepository
	fileUpload        fileUpload.FileUpload
}

// Create implements ProductUseCase
func (pu *ProductUseCaseImpl) Create(productDto productDto.ProductRequestBody) (*productEntity.Product, error) {
	dataProduct := productEntity.Product{
		ProductCategoryID: productDto.ProductCategoryId,
		Title:             productDto.Title,
		Description:       productDto.Description,
		Price:             productDto.Price,
		CreatedByID:       productDto.CreatedBy,
	}

	if productDto.Image != nil {
		image, err := pu.fileUpload.Upload(*productDto.Image)
		if err != nil {
			return nil, err
		}

		if image != nil {
			dataProduct.Image = image
		}
	}

	if productDto.Video != nil {
		video, err := pu.fileUpload.Upload(*productDto.Video)
		if err != nil {
			return nil, err
		}

		if video != nil {
			dataProduct.Video = video
		}
	}

	product, err := pu.productRepository.Create(dataProduct)

	if err != nil {
		return nil, err
	}

	return product, nil
}

// Delete implements ProductUseCase
func (pu *ProductUseCaseImpl) Delete(id int) error {
	product, err := pu.productRepository.FindById(id)

	if err != nil {
		return err
	}

	if product.Image != nil {
		_, err := pu.fileUpload.Delete(*product.Image)

		if err != nil {
			return err
		}
	}

	if product.Video != nil {
		_, err := pu.fileUpload.Delete(*product.Video)

		if err != nil {
			return err
		}
	}

	err = pu.productRepository.Delete(*product)

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductUseCase
func (pu *ProductUseCaseImpl) FindAll(offset int, limit int) []productEntity.Product {
	return pu.productRepository.FindAll(offset, limit)
}

// FindById implements ProductUseCase
func (pu *ProductUseCaseImpl) FindById(id int) (*productEntity.Product, error) {
	return pu.productRepository.FindById(id)
}

// Update implements ProductUseCase
func (pu *ProductUseCaseImpl) Update(id int, productDto productDto.ProductRequestBody) (*productEntity.Product, error) {
	product, err := pu.productRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	product.Title = productDto.Title
	product.Description = productDto.Description
	product.Price = productDto.Price
	product.UpdatedByID = &productDto.UpdatedBy

	// upload image
	if productDto.Image != nil {
		if product.Image != nil {
			_, err := pu.fileUpload.Delete(*product.Image)

			if err != nil {
				return nil, err
			}
		}

		image, err := pu.fileUpload.Upload(*productDto.Image)

		if err != nil {
			return nil, err
		}

		if image != nil {
			product.Image = image
		}
	}

	// upload video
	if productDto.Video != nil {
		if product.Video != nil {
			res, err := pu.fileUpload.Delete(*product.Video)

			if err != nil {
				return nil, err
			}

			fmt.Println(*res)
		}

		video, err := pu.fileUpload.Upload(*productDto.Video)

		if err != nil {
			return nil, err
		}

		if video != nil {
			product.Video = video
		}
	}

	updatedProduct, err := pu.productRepository.Update(*product)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func NewProductUseCase(
	productRepository productRepository.ProductRepository,
	fileUpload fileUpload.FileUpload) ProductUseCase {
	return &ProductUseCaseImpl{productRepository, fileUpload}
}

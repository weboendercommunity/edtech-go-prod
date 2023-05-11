package product

import "mime/multipart"

type ProductRequestBody struct {
	ProductCategoryId int64                 `form:"product_category_id" binding:"required"`
	Title             string                `form:"title" binding:"required"`
	Image             *multipart.FileHeader `form:"image"`
	Video             *multipart.FileHeader `form:"video"`
	Description       string                `form:"description"`
	Price             int64                 `form:"price" binding:"required"`
	CreatedBy         int64                 `form:"created_by"`
	UpdatedBy         int64                 `form:"updated_by"`
}

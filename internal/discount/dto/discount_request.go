package discount

import (
	"time"

	discountEntity "edtech.id/internal/discount/entity"
)

type DiscountRequestBody struct {
	Name            string                      `json:"name" binding:"required"`
	Code            string                      `json:"code" binding:"required"`
	Quantity        int                         `json:"quantity" binding:"required,number"`
	Value           int                         `json:"value" binding:"required,number"`
	RemaingQuantity int64                       `json:"remaining_quantity" binding:"number"`
	Type            discountEntity.DiscountType `json:"type" binding:"required"`
	CreatedBy       int64                       `json:"created_by"`
	UpdatedBy       int64                       `json:"updated_by"`
	StartDate       time.Time                   `json:"start_date"`
	EndDate         time.Time                   `json:"end_date"`
}

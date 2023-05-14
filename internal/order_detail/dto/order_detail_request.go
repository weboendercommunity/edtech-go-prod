package order_detail

type OrderDetailRequestBody struct {
	ID        int64 `json:"id"`
	Price     int64 `json:"price"`
	ProductID int64 `json:"product_id"`
	OrderID   int64 `json:"order_id"`
	CreatedBy int64 `json:"created_by"`
}

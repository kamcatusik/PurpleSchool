package order

import "4-order-api/internal/models"

type QuantProductID struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
type OrderRequest struct {
	UserID   uint             `json:"user_id"`
	Products []QuantProductID `json:"quantproduct_id"`
}
type OrderResponse struct {
	Order          *models.Order
	QuantProductID []QuantProductID `json:"quant_prod_id"`
}

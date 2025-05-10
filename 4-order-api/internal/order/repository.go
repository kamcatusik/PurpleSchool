package order

import (
	"4-order-api/internal/models"
	"4-order-api/pkg/db"
	"errors"

	"gorm.io/gorm"
)

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}
func (repo *OrderRepository) CreateOrder(req *OrderRequest) (*OrderResponse, error) {
	orderResp := &OrderResponse{
		Order: &models.Order{
			UserId: req.UserID,
		},
		QuantProductID: make([]QuantProductID, 0, len(req.Products)),
	}

	for _, orderProdId := range req.Products {

		var product models.Product
		err := repo.Database.DB.Model(&models.Product{}).Where("id = ?", orderProdId.ProductID).First(&product).Error
		if err != nil {
			return nil, errors.New("товар не найден")
		}

		if orderProdId.Quantity > product.Quantity {
			return nil, errors.New("недостаточно товара")

		}
		//orderResp.Order.Products = append(orderResp.Order.Products, &product)
		orderResp.QuantProductID = append(orderResp.QuantProductID, orderProdId)

	}

	repo.Database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(orderResp.Order).Error; err != nil {
			return err
		}
		for _, prodId := range req.Products {
			orderProduct := models.OrderProduct{
				OrderID:   orderResp.Order.ID,
				ProductID: prodId.ProductID,
				Quantity:  prodId.Quantity,
			}
			if err := tx.Create(&orderProduct).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return orderResp, nil
}

func (repo *OrderRepository) FindOrderId(userId uint) ([]*models.Order, error) {
	var orders []*models.Order

	err := repo.Database.DB.
		Where("user_id = ?", userId).
		Preload("Products").
		Preload("Products.OrderProduct", func(db *gorm.DB) *gorm.DB {
			return db.Where("order_id IN (?)", repo.Database.DB.
				Table("orders").
				Select("id").
				Where("user_id = ?", userId))
		}).
		Find(&orders).Error

	if err != nil || len(orders) == 0 {
		return nil, err
	}

	return orders, nil
}
func (repo *OrderRepository) GetOrder(orderId uint) (*models.Order, error) {
	var order models.Order
	result := repo.Database.DB.Where("id =?", orderId).
		Preload("Products").
		Preload("Products.OrderProduct", "order_id =?", orderId).
		First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

package order

import (
	"4-order-api/internal/models"
	"4-order-api/pkg/db"
	"errors"
	"fmt"

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
func (repo *OrderRepository) CreateOrder(order *models.Order, quantProd []QuantProductID) (*models.Order, error) {

	repo.Database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		for _, prodquant := range quantProd {
			var product models.Product
			if err := tx.First(&product, prodquant.ProductID).Error; err != nil {
				return errors.New("товар не найден")
			}
			if product.Quantity < prodquant.Quantity {
				return errors.New("недостаточно товара")
			}
			orderProduct := models.OrderProduct{
				OrderID:   order.ID,
				ProductID: prodquant.ProductID,
				Quantity:  prodquant.Quantity,
			}
			fmt.Println(prodquant.Quantity)
			if err := tx.Create(&orderProduct).Error; err != nil {
				return err
			}

		}
		return nil
	})
	fmt.Println(order.Products)

	return order, nil
}

func (repo *OrderRepository) FindOrderId(userId uint) ([]*models.Order, error) {
	var orders []*models.Order

	err := repo.Database.DB.Debug().
		Where("user_id = ?", userId).
		Preload("Products").
		Preload("Products.OrderProduct", func(db *gorm.DB) *gorm.DB {
			return db.Where("order_id IN (?)", repo.Database.DB.
				Table("orders").
				Select("id").
				Where("user_id = ?", userId))
		}).
		Find(&orders).Error

	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, errors.New("заказов не найдено")
	}

	return orders, nil
}
func (repo *OrderRepository) GetOrder(orderId uint) (*models.Order, error) {
	var order models.Order
	result := repo.Database.DB.Debug().
		Where("id =?", orderId).
		Preload("Products").
		Preload("Products.OrderProduct", "order_id =?", orderId).
		First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

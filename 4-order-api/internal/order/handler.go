package order

import (
	"4-order-api/configs"
	"4-order-api/internal/models"
	"4-order-api/internal/product"
	"4-order-api/pkg/middleware"
	"4-order-api/pkg/req"
	"4-order-api/pkg/resp"
	"fmt"
	"net/http"
	"strconv"
)

type OrderHandlerDeps struct {
	*OrderRepository
	*configs.Config
	*product.ProductRepository
}
type OrderHandler struct {
	*configs.Config
	*OrderRepository
	*product.ProductRepository
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		Config:            deps.Config,
		OrderRepository:   deps.OrderRepository,
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("POST /order", middleware.Auth(handler.order, deps.Config))
	router.HandleFunc("GET /order/{id}", middleware.Auth(handler.getOrder, deps.Config))
	router.HandleFunc("GET /my-orders/{id}", middleware.Auth(handler.getOrderByUser, deps.Config))
}
func (handler *OrderHandler) order(w http.ResponseWriter, request *http.Request) {
	body, err := req.HandleBody[OrderRequest](w, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	productIDs := make([]uint, len(body.Products))
	for i, p := range body.Products {
		productIDs[i] = p.ProductID
	}
	products, err := handler.ProductRepository.FindProductById(productIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderResp := &models.Order{
		UserId:   body.UserID,
		Products: products,
	}

	CreatedOrder, err := handler.OrderRepository.CreateOrder(orderResp, body.Products)

	newOrder := GetResponseOrder(CreatedOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productMap := make(map[uint]*ProductResponse)
	for i := range newOrder.Products {
		productMap[newOrder.Products[i].ID] = &newOrder.Products[i]
	}

	for _, reqProduct := range body.Products {
		if existing, ok := productMap[reqProduct.ProductID]; ok {

			existing.Quantity = reqProduct.Quantity

		}
	}
	fmt.Println(newOrder)

	resp.Json(w, newOrder, 201)
}

func (handler *OrderHandler) getOrder(w http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	resId, _ := strconv.Atoi(idStr)

	getOrders, err := handler.OrderRepository.GetOrder(uint(resId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order := GetResponseOrder(getOrders)
	resp.Json(w, order, http.StatusOK)
}
func (handler *OrderHandler) getOrderByUser(w http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	resId, _ := strconv.Atoi(idStr)
	getOrders, err := handler.OrderRepository.FindOrderId(uint(resId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orders := AllGetOrderByUserResponse{
		Orders: make([]OrderResponse, 0, len(getOrders)),
	}
	for _, order := range getOrders {

		orders.Orders = append(orders.Orders, GetResponseOrder(order))
	}

	resp.Json(w, orders, http.StatusOK)
}
func GetResponseOrder(getOrders *models.Order) OrderResponse {
	order := OrderResponse{
		ID:        getOrders.ID,
		UserID:    getOrders.UserId,
		CreatedAt: getOrders.CreatedAt,
		UpdatedAt: getOrders.UpdatedAt,
		Products:  make([]ProductResponse, 0, len(getOrders.Products)),
	}
	for _, product := range getOrders.Products {

		var quantity uint
		for _, q := range product.OrderProduct {
			if order.ID == q.OrderID {
				quantity = q.Quantity
				break
			}
		}
		order.Products = append(order.Products, ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Images:      product.Images,
			Quantity:    quantity,
		})

	}
	return order
}

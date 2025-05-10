package order

import (
	"4-order-api/configs"
	"4-order-api/internal/product"
	"4-order-api/pkg/req"
	"4-order-api/pkg/resp"
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
	router.HandleFunc("POST /order", handler.order)
	router.HandleFunc("GET /order/{id}", handler.getOrder)
	router.HandleFunc("GET /my-orders/{id}", handler.getOrderByUser)
}
func (handler *OrderHandler) order(w http.ResponseWriter, request *http.Request) {
	body, err := req.HandleBody[OrderRequest](w, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	NewOrder, err := handler.OrderRepository.CreateOrder(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp.Json(w, NewOrder, 201)
}

func (handler *OrderHandler) getOrder(w http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	resId, _ := strconv.Atoi(idStr)

	getOrders, err := handler.OrderRepository.GetOrder(uint(resId))
	if err != nil {
		resp.Json(w, "Заказ не найден", http.StatusNotFound)
		return
	}
	resp.Json(w, getOrders, http.StatusOK)
}
func (handler *OrderHandler) getOrderByUser(w http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	resId, _ := strconv.Atoi(idStr)
	getOrders, err := handler.OrderRepository.FindOrderId(uint(resId))
	if err != nil {
		resp.Json(w, "Заказ не найден", http.StatusNotFound)
		return
	}
	resp.Json(w, getOrders, http.StatusOK)
}

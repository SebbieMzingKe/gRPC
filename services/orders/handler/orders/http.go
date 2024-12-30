package handlers

import (
	"kitchen/services/common/genproto/orders"
	"kitchen/services/orders/types"
	"net/http"

	"kitchen/services/common/util"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHadler(ordersService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: ordersService,
	}
	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}


func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	order := &orders.Order{
		OrderID: 42,
		CustomerID: req.GetCustomerID(),
		ProductID: req.GetProductID(),
		Quantity: req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	util.WriteJSON(w, http.StatusOK, res)
}
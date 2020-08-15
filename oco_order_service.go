package binance

import (
	"context"
)

// CreateOCOService create OCO order
type CreateOCOService struct {
	c                    *Client
	symbol               string
	listClientOrderID    *string
	side                 SideType
	quantity             *string
	limitClientOrderID   *string
	price                *string
	limitIcebergQty      *string
	stopClientOrderID    *string
	stopPrice            *string
	stopLimitPrice       *string
	stopIcebergQty       *string
	stopLimitTimeInForce *TimeInForceType
	newOrderRespType     *NewOrderRespType
}

// Symbol set symbol
func (s *CreateOCOService) Symbol(symbol string) *CreateOCOService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOCOService) Side(side SideType) *CreateOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *CreateOCOService) Quantity(quantity string) *CreateOCOService {
	s.quantity = &quantity
	return s
}

// LimitClientOrderID set limitClientOrderID
func (s *CreateOCOService) LimitClientOrderID(limitClientOrderID string) *CreateOCOService {
	s.limitClientOrderID = &limitClientOrderID
	return s
}

// Price set price
func (s *CreateOCOService) Price(price string) *CreateOCOService {
	s.price = &price
	return s
}

// limitIcebergQuantity set limitIcebergQuantity
func (s *CreateOCOService) limitIcebergQuantity(limitIcebergQty string) *CreateOCOService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// StopClientOrderID set stopClientOrderID
func (s *CreateOCOService) StopClientOrderID(stopClientOrderID string) *CreateOCOService {
	s.stopClientOrderID = &stopClientOrderID
	return s
}

// StopPrice set stop price
func (s *CreateOCOService) StopPrice(stopPrice string) *CreateOCOService {
	s.stopPrice = &stopPrice
	return s
}

// StopLimitPrice set stop limit price
func (s *CreateOCOService) StopLimitPrice(stopLimitPrice string) *CreateOCOService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopIcebergQty set stop limit price
func (s *CreateOCOService) StopIcebergQty(stopIcebergQty string) *CreateOCOService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *CreateOCOService) StopLimitTimeInForce(stopLimitTimeInForce TimeInForceType) *CreateOCOService {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

// NewOrderRespType set icebergQuantity
func (s *CreateOCOService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *CreateOCOService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":    s.symbol,
		"side":      s.side,
		"quantity":  *s.quantity,
		"price":     *s.price,
		"stopPrice": *s.stopPrice,
	}

	if s.listClientOrderID != nil {
		m["listClientOrderId"] = *s.listClientOrderID
	}
	if s.limitClientOrderID != nil {
		m["limitClientOrderId"] = *s.limitClientOrderID
	}
	if s.limitIcebergQty != nil {
		m["limitIcebergQty"] = *s.limitIcebergQty
	}
	if s.stopClientOrderID != nil {
		m["stopClientOrderId"] = *s.stopClientOrderID
	}
	if s.stopLimitPrice != nil {
		m["stopLimitPrice"] = *s.stopLimitPrice
	}
	if s.stopIcebergQty != nil {
		m["stopIcebergQty"] = *s.stopIcebergQty
	}
	if s.stopLimitTimeInForce != nil {
		m["stopLimitTimeInForce"] = *s.stopLimitTimeInForce
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateOCOService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOCOResponse, err error) {
	data, err := s.createOrder(ctx, "/api/v3/order/oco", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOCOResponse define create OCO order response
type CreateOCOResponse struct {
	OrderListID       int64                  `json:"orderListId"`
	ContingencyType   string                 `json:"contingencyType"`
	ListStatusType    OCOListStatusType      `json:"listStatusType"`
	ListOrderStatus   OCOListOrderStatusType `json:"listOrderStatus"`
	ListClientOrderID string                 `json:"listClientOrderId"`
	TransactionTime   int64                  `json:"transactionTime"`
	Symbol            string                 `json:"symbol"`
	Orders            []*OCOOrder            `json:"orders"`
	OrderReports      []*OCOOrderReport      `json:"orderReports"`
}

// OCOResponse define OCO order response
type OCOResponse struct {
	OrderListID       int64                  `json:"orderListId"`
	ContingencyType   string                 `json:"contingencyType"`
	ListStatusType    OCOListStatusType      `json:"listStatusType"`
	ListOrderStatus   OCOListOrderStatusType `json:"listOrderStatus"`
	ListClientOrderID string                 `json:"listClientOrderId"`
	TransactionTime   int64                  `json:"transactionTime"`
	Symbol            string                 `json:"symbol"`
	Orders            []*OCOOrder            `json:"orders"`
}

// OCOOrder may be returned in an array of OCOOrder in a CreateOCOResponse.
type OCOOrder struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

// OCOOrderReport may be returned in an array of OCOOrderReport in a CreateOCOResponse.
type OCOOrderReport struct {
	Symbol                   string          `json:"symbol"`
	OrderID                  int64           `json:"orderId"`
	OrderListID              int64           `json:"orderListId"`
	ClientOrderID            string          `json:"clientOrderId"`
	TransactionTime          int64           `json:"transactionTime"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
	StopPrice                string          `json:"stopPrice"`
}

// ListOpenOCOService list opened OCO orders
type ListOpenOCOService struct {
	c      *Client
}

// Do send request
func (s *ListOpenOCOService) Do(ctx context.Context, opts ...RequestOption) (res []*OCOResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/openOrderList",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*OCOResponse{}, err
	}
	res = make([]*OCOResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*OCOResponse{}, err
	}
	return res, nil
}

// GetOCOService get an OCO order
type GetOCOService struct {
	c                 *Client
	orderListID       *int64
	origClientOrderID *string
}

// OrderListID set orderListID
func (s *GetOCOService) OrderListID(orderListID int64) *GetOCOService {
	s.orderListID = &orderListID
	return s
}

// OrigClientOrderID set listClientOrderID
func (s *GetOCOService) OrigClientOrderID(origClientOrderID string) *GetOCOService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetOCOService) Do(ctx context.Context, opts ...RequestOption) (res *OCOResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/orderList",
		secType:  secTypeSigned,
	}
	if s.orderListID != nil {
		r.setParam("orderListId", *s.orderListID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListOCOService all account orders; active, canceled, or filled
type ListOCOService struct {
	c         *Client
	symbol    string
	fromID    *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListOCOService) Symbol(symbol string) *ListOCOService {
	s.symbol = symbol
	return s
}

// FromID set orderListID
func (s *ListOCOService) FromID(fromID int64) *ListOCOService {
	s.fromID = &fromID
	return s
}

// StartTime set starttime
func (s *ListOCOService) StartTime(startTime int64) *ListOCOService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListOCOService) EndTime(endTime int64) *ListOCOService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListOCOService) Limit(limit int) *ListOCOService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListOCOService) Do(ctx context.Context, opts ...RequestOption) (res []*OCOResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v3/allOrderList",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*OCOResponse{}, err
	}
	res = make([]*OCOResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*OCOResponse{}, err
	}
	return res, nil
}

// CancelOCOService cancel an OCO order
type CancelOCOService struct {
	c                 *Client
	symbol            string
	orderListID       *int64
	listClientOrderID *string
	newClientOrderID  *string
}

// Symbol set symbol
func (s *CancelOCOService) Symbol(symbol string) *CancelOCOService {
	s.symbol = symbol
	return s
}

// OrderListID set orderListID
func (s *CancelOCOService) OrderListID(orderListID int64) *CancelOCOService {
	s.orderListID = &orderListID
	return s
}

// ListClientOrderID set listClientOrderID
func (s *CancelOCOService) ListClientOrderID(listClientOrderID string) *CancelOCOService {
	s.listClientOrderID = &listClientOrderID
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CancelOCOService) NewClientOrderID(newClientOrderID string) *CancelOCOService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// Do send request
func (s *CancelOCOService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOCOResponse, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/api/v3/orderList",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderListID != nil {
		r.setFormParam("orderListId", *s.orderListID)
	}
	if s.listClientOrderID != nil {
		r.setFormParam("listClientOrderId", *s.listClientOrderID)
	}
	if s.newClientOrderID != nil {
		r.setFormParam("newClientOrderId", *s.newClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelOCOResponse define response of canceling OCO order
type CancelOCOResponse struct {
	OrderListID       int64                  `json:"orderListId"`
	ContingencyType   string                 `json:"contingencyType"`
	ListStatusType    OCOListStatusType      `json:"listStatusType"`
	ListOrderStatus   OCOListOrderStatusType `json:"listOrderStatus"`
	ListClientOrderID string                 `json:"listClientOrderId"`
	TransactionTime   int64                  `json:"transactionTime"`
	Symbol            string                 `json:"symbol"`
	Orders            []*OCOOrder            `json:"orders"`
	OrderReports      []*OCOOrderReport      `json:"orderReports"`
}

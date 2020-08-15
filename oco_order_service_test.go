package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseOrderTestSuite struct {
	baseTestSuite
}

type orderServiceTestSuite struct {
	baseOrderTestSuite
}

func TestOrderService(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (s *orderServiceTestSuite) TestCreateOCO() {
	data := []byte(`{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "C3wyj4WVEktd7u9aVBRXcN",
		"transactionTime": 1574040868128,
		"symbol": "LTCBTC",
		"orders": [
		  {
			"symbol": "LTCBTC",
			"orderId": 2,
			"clientOrderId": "pO9ufTiFGg3nw2fOdgeOXa"
		  },
		  {
			"symbol": "LTCBTC",
			"orderId": 3,
			"clientOrderId": "TXOvglzXuaubXAaENpaRCB"
		  }
		],
		"orderReports": [
		  {
			"symbol": "LTCBTC",
			"origClientOrderId": "pO9ufTiFGg3nw2fOdgeOXa",
			"orderId": 2,
			"orderListId": 0,
			"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
			"price": "1.00000000",
			"origQty": "10.00000000",
			"executedQty": "0.00000000",
			"cummulativeQuoteQty": "0.00000000",
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "STOP_LOSS",
			"side": "SELL",
			"stopPrice": "1.00000000"
		  },
		  {
			"symbol": "LTCBTC",
			"origClientOrderId": "TXOvglzXuaubXAaENpaRCB",
			"orderId": 3,
			"orderListId": 0,
			"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
			"price": "3.00000000",
			"origQty": "10.00000000",
			"executedQty": "0.00000000",
			"cummulativeQuoteQty": "0.00000000",
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "LIMIT_MAKER",
			"side": "SELL"
		  }
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	timeInForce := TimeInForceTypeGTC
	quantity := "10"
	price := "3"
	stopPrice := "3.1"
	stopLimitPrice := "3.2"
	limitClientOrderID := "myOrder1"
	newOrderRespType := NewOrderRespTypeFULL
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":               symbol,
			"side":                 side,
			"quantity":             quantity,
			"price":                price,
			"stopPrice":            stopPrice,
			"stopLimitPrice":       stopLimitPrice,
			"stopLimitTimeInForce": timeInForce,
			"limitClientOrderId":   limitClientOrderID,
			"newOrderRespType":     newOrderRespType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOCOService().
		Symbol(symbol).
		Side(side).
		Quantity(quantity).
		Price(price).
		StopPrice(stopPrice).
		StopLimitPrice(stopLimitPrice).
		StopLimitTimeInForce(timeInForce).
		LimitClientOrderID(limitClientOrderID).
		NewOrderRespType(newOrderRespType).
		Do(newContext())

	s.r().NoError(err)
	e := &CreateOCOResponse{
		OrderListID:       0,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "EXECUTING",
		ListClientOrderID: "C3wyj4WVEktd7u9aVBRXcN",
		TransactionTime:   1574040868128,
		Symbol:            "LTCBTC",
		Orders: []*OCOOrder{
			{
				Symbol:        "LTCBTC",
				OrderID:       2,
				ClientOrderID: "pO9ufTiFGg3nw2fOdgeOXa",
			},
			{
				Symbol:        "LTCBTC",
				OrderID:       3,
				ClientOrderID: "TXOvglzXuaubXAaENpaRCB",
			},
		},
		OrderReports: []*OCOOrderReport{
			{
				Symbol:                   "LTCBTC",
				OrderID:                  2,
				OrderListID:              0,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
				Price:                    "1.00000000",
				OrigQuantity:             "10.00000000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeNew,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeStopLoss,
				Side:                     SideTypeSell,
				StopPrice:                "1.00000000",
			},
			{
				Symbol:                   "LTCBTC",
				OrderID:                  3,
				OrderListID:              0,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
				Price:                    "3.00000000",
				OrigQuantity:             "10.00000000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeNew,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeLimitMaker,
				Side:                     SideTypeSell,
			},
		},
	}
	s.assertCreateOCOResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertCreateOCOResponseEqual(e, a *CreateOCOResponse) {
	r := s.r()
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")

	r.Len(a.OrderReports, len(e.OrderReports))
	for idx, orderReport := range e.OrderReports {
		s.assertOCOReportEqual(orderReport, a.OrderReports[idx])
	}

	r.Len(a.Orders, len(e.Orders))
	for idx, order := range e.Orders {
		s.assertOCOOrderEqual(order, a.Orders[idx])
	}
}

func (s *baseOrderTestSuite) assertOCOReportEqual(e, a *OCOOrderReport) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.OrderID, a.OrderID, "OrderListID")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	// r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
}

func (s *baseOrderTestSuite) assertOCOOrderEqual(e, a *OCOOrder) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.OrderID, a.OrderID, "OrderListID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}

func (s *orderServiceTestSuite) TestListOpenOCO() {
	data := []byte(`[
	  {
		"orderListId": 31,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "wuB13fmulKj3YjdqWEcsnp",
		"transactionTime": 1565246080644,
		"symbol": "LTCBTC",
		"orders": [
		  {
			"symbol": "LTCBTC",
			"orderId": 4,
			"clientOrderId": "r3EH2N76dHfLoSZWIUw1bT"
		  },
		  {
			"symbol": "LTCBTC",
			"orderId": 5,
			"clientOrderId": "Cv1SnyPD3qhqpbjpYEHbd2"
		  }
		]
	  }
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	orders, err := s.client.NewListOpenOCOService().Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &OCOResponse{
		OrderListID:       31,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "EXECUTING",
		ListClientOrderID: "wuB13fmulKj3YjdqWEcsnp",
		TransactionTime:   1565246080644,
		Symbol:            "LTCBTC",
		Orders: []*OCOOrder{
			{
				Symbol:        "LTCBTC",
				OrderID:       4,
				ClientOrderID: "r3EH2N76dHfLoSZWIUw1bT",
			},
			{
				Symbol:        "LTCBTC",
				OrderID:       5,
				ClientOrderID: "Cv1SnyPD3qhqpbjpYEHbd2",
			},
		},
	}
	s.assertOCOEqual(e, orders[0])
}

func (s *baseOrderTestSuite) assertOCOEqual(e, a *OCOResponse) {
	r := s.r()
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Len(a.Orders, len(e.Orders))

	for idx, order := range e.Orders {
		s.assertOCOOrderEqual(order, a.Orders[idx])
	}
}

func (s *orderServiceTestSuite) TestGetOCO() {
	data := []byte(`{
	  "orderListId": 27,
	  "contingencyType": "OCO",
	  "listStatusType": "EXEC_STARTED",
	  "listOrderStatus": "EXECUTING",
	  "listClientOrderId": "h2USkA5YQpaXHPIrkd96xE",
	  "transactionTime": 1565245656253,
	  "symbol": "LTCBTC",
	  "orders": [
		{
		  "symbol": "LTCBTC",
		  "orderId": 4,
		  "clientOrderId": "qD1gy3kc3Gx0rihm9Y3xwS"
		},
		{
		  "symbol": "LTCBTC",
		  "orderId": 5,
		  "clientOrderId": "ARzZ9I00CPM8i3NhmU9Ega"
		}
	  ]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	orderListID := int64(27)
	origClientOrderID := "h2USkA5YQpaXHPIrkd96xE"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"orderListId":       orderListID,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})
	order, err := s.client.NewGetOCOService().OrderListID(orderListID).
		OrigClientOrderID(origClientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &OCOResponse{
		OrderListID:       27,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "EXECUTING",
		ListClientOrderID: "h2USkA5YQpaXHPIrkd96xE",
		TransactionTime:   1565245656253,
		Symbol:            "LTCBTC",
		Orders: []*OCOOrder{
			{
				Symbol:        "LTCBTC",
				OrderID:       4,
				ClientOrderID: "qD1gy3kc3Gx0rihm9Y3xwS",
			},
			{
				Symbol:        "LTCBTC",
				OrderID:       5,
				ClientOrderID: "ARzZ9I00CPM8i3NhmU9Ega",
			},
		},
	}
	s.assertOCOEqual(e, order)
}

func (s *orderServiceTestSuite) TestListOCO() {
	data := []byte(`[
	  {
		"orderListId": 29,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "amEEAXryFzFwYF1FeRpUoZ",
		"transactionTime": 1565245913483,
		"symbol": "LTCBTC",
		"orders": [
		  {
			"symbol": "LTCBTC",
			"orderId": 4,
			"clientOrderId": "oD7aesZqjEGlZrbtRpy5zB"
		  },
		  {
			"symbol": "LTCBTC",
			"orderId": 5,
			"clientOrderId": "Jr1h6xirOxgeJOUuYQS7V3"
		  }
		]
	  }
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	fromID := int64(29)
	limit := 3
	startTime := int64(1565245913483)
	endTime := int64(1565245913484)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"fromId":    fromID,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewListOCOService().Symbol(symbol).
		FromID(fromID).StartTime(startTime).EndTime(endTime).
		Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &OCOResponse{
		OrderListID:       29,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "EXECUTING",
		ListClientOrderID: "amEEAXryFzFwYF1FeRpUoZ",
		TransactionTime:   1565245913483,
		Symbol:            "LTCBTC",
		Orders: []*OCOOrder{
			{
				Symbol:        "LTCBTC",
				OrderID:       4,
				ClientOrderID: "oD7aesZqjEGlZrbtRpy5zB",
			},
			{
				Symbol:        "LTCBTC",
				OrderID:       5,
				ClientOrderID: "Jr1h6xirOxgeJOUuYQS7V3",
			},
		},
	}
	s.assertOCOEqual(e, orders[0])
}

func (s *orderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
	  "orderListId": 0,
	  "contingencyType": "OCO",
	  "listStatusType": "ALL_DONE",
	  "listOrderStatus": "ALL_DONE",
	  "listClientOrderId": "C3wyj4WVEktd7u9aVBRXcN",
	  "transactionTime": 1574040868128,
	  "symbol": "LTCBTC",
	  "orders": [
		{
		  "symbol": "LTCBTC",
		  "orderId": 2,
		  "clientOrderId": "pO9ufTiFGg3nw2fOdgeOXa"
		},
		{
		  "symbol": "LTCBTC",
		  "orderId": 3,
		  "clientOrderId": "TXOvglzXuaubXAaENpaRCB"
		}
	  ],
	  "orderReports": [
		{
		  "symbol": "LTCBTC",
		  "origClientOrderId": "pO9ufTiFGg3nw2fOdgeOXa",
		  "orderId": 2,
		  "orderListId": 0,
		  "clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
		  "price": "1.00000000",
		  "origQty": "10.00000000",
		  "executedQty": "0.00000000",
		  "cummulativeQuoteQty": "0.00000000",
		  "status": "CANCELED",
		  "timeInForce": "GTC",
		  "type": "STOP_LOSS_LIMIT",
		  "side": "SELL",
		  "stopPrice": "1.00000000"
		},
		{
		  "symbol": "LTCBTC",
		  "origClientOrderId": "TXOvglzXuaubXAaENpaRCB",
		  "orderId": 3,
		  "orderListId": 0,
		  "clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
		  "price": "3.00000000",
		  "origQty": "10.00000000",
		  "executedQty": "0.00000000",
		  "cummulativeQuoteQty": "0.00000000",
		  "status": "CANCELED",
		  "timeInForce": "GTC",
		  "type": "LIMIT_MAKER",
		  "side": "SELL"
		}
	  ]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderListID := int64(0)
	listClientOrderID := "C3wyj4WVEktd7u9aVBRXcN"
	newClientOrderID := "cancelMyOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":            symbol,
			"orderListId":       orderListID,
			"listClientOrderId": listClientOrderID,
			"newClientOrderId":  newClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelOCOService().Symbol(symbol).
		OrderListID(orderListID).ListClientOrderID(listClientOrderID).
		NewClientOrderID(newClientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CancelOCOResponse{
		OrderListID:       0,
		ContingencyType:   "OCO",
		ListStatusType:    "ALL_DONE",
		ListOrderStatus:   "ALL_DONE",
		ListClientOrderID: "C3wyj4WVEktd7u9aVBRXcN",
		TransactionTime:   1574040868128,
			Symbol:            "LTCBTC",
		Orders: []*OCOOrder{
			{
				Symbol:        "LTCBTC",
				OrderID:       2,
				ClientOrderID: "pO9ufTiFGg3nw2fOdgeOXa",
			},
			{
				Symbol:        "LTCBTC",
				OrderID:       3,
				ClientOrderID: "TXOvglzXuaubXAaENpaRCB",
			},
		},
		OrderReports: []*OCOOrderReport{
			{
				Symbol:                   "LTCBTC",
				OrderID:                  2,
				OrderListID:              0,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
				Price:                    "1.00000000",
				OrigQuantity:             "10.00000000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeCanceled,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeStopLoss,
				Side:                     SideTypeSell,
				StopPrice:                "1.00000000",
			},
			{
				Symbol:                   "LTCBTC",
				OrderID:                  3,
				OrderListID:              0,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
				Price:                    "3.00000000",
				OrigQuantity:             "10.00000000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeCanceled,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeLimitMaker,
				Side:                     SideTypeSell,
			},
		},
	}
	s.assertCancelOrderResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertCancelOrderResponseEqual(e, a *CancelOCOResponse) {
	r := s.r()
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")

	r.Len(a.OrderReports, len(e.OrderReports))
	for idx, orderReport := range e.OrderReports {
		s.assertOCOReportEqual(orderReport, a.OrderReports[idx])
	}

	r.Len(a.Orders, len(e.Orders))
	for idx, order := range e.Orders {
		s.assertOCOOrderEqual(order, a.Orders[idx])
	}
}

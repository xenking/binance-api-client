package binance

import (
	"fmt"
	"strings"
	"time"
)

var (
	baseURL = "wss://stream.binance.com:9443/ws"
	//baseFutureURL   = "wss://fstream.binance.com/ws"
	//combinedBaseURL = "wss://stream.binance.com:9443/stream?streams="
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
	)

// WsBookTickerEvent define websocket individual book ticker event
type WsBookTickerEvent struct {
	UpdateID    int64  `json:"u"`
	Symbol      string `json:"s"`
	BidPrice    string `json:"b"`
	BidQuantity string `json:"B"`
	AskPrice    string `json:"a"`
	AskQuantity string `json:"A"`
}

// WsBookTickerHandler handle websocket partial depth event
type WsBookTickerHandler func(event *WsBookTickerEvent)

// WsBookTickerServe serve websocket partial depth handler with a symbol
func WsBookTickerServe(symbol string, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@bookTicker", baseURL, strings.ToLower(symbol))

	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBookTickerEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

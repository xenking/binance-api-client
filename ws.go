package binance

import (
	"time"

	"github.com/dgrr/fastws"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {

	c, err := fastws.Dial(cfg.Endpoint)

	if err != nil {
		return nil, nil, err
	}
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		defer func() {
			cerr := c.Close()
			if cerr != nil {
				errHandler(cerr)
			}
		}()
		defer close(doneC)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		var msg []byte
		for {
			select {
			case <-stopC:
				return
			default:
				_, msg, err = c.ReadMessage(msg[:0])
				if err != nil {
					errHandler(err)
					return
				}
				handler(msg)
			}
		}
	}()
	return
}

func keepAlive(c *fastws.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()

	go func() {
		defer ticker.Stop()
		for {
			err := c.SendCode(fastws.CodePing, fastws.StatusNone, []byte{})
			if err != nil {
				return
			}
			<-ticker.C
			if time.Now().Sub(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}

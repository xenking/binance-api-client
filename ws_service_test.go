package binance

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)


type websocketServiceTestSuite struct {
	suite.Suite
	origWsServe func(*WsConfig, WsHandler, ErrHandler) (chan struct{}, chan struct{}, error)
	serveCount  int
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(websocketServiceTestSuite))
}

func (s *websocketServiceTestSuite) r() *require.Assertions {
	return s.Require()
}


func (s *websocketServiceTestSuite) SetupTest() {
	s.origWsServe = wsServe
}

func (s *websocketServiceTestSuite) TearDownTest() {
	wsServe = s.origWsServe
	s.serveCount = 0
}

func (s *websocketServiceTestSuite) mockWsServe(data []byte, err error) {
	wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, innerErr error) {
		s.serveCount++
		doneC = make(chan struct{})
		stopC = make(chan struct{})
		go func() {
			<-stopC
			close(doneC)
		}()
		handler(data)
		if err != nil {
			errHandler(err)
		}
		return doneC, stopC, nil
	}
}

func (s *websocketServiceTestSuite) assertWsServe(count ...int) {
	e := 1
	if len(count) > 0 {
		e = count[0]
	}
	s.r().Equal(e, s.serveCount)
}

func (s *websocketServiceTestSuite) TestBookTickerServe() {
	data := []byte(`{
		"u":400900217,    
		"s":"BNBUSDT",    
		"b":"25.35190000", 
		"B":"31.21000000", 
		"a":"25.36520000",
		"A":"40.66000000"  
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBookTickerServe("ETHBTC", func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			UpdateID:    400900217,
			Symbol:      "BNBUSDT",
			BidPrice:    "25.35190000",
			BidQuantity: "31.21000000",
			AskPrice:    "25.36520000",
			AskQuantity: "40.66000000",
		}
		s.assertWsBookTickerEventEqual(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsBookTickerEventEqual(e, a *WsBookTickerEvent) {
	r := s.r()
	r.Equal(e.UpdateID, a.UpdateID, "UpdateID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.BidPrice, a.BidPrice, "AskQuantity")
	r.Equal(e.BidQuantity, a.BidQuantity, "AskQuantity")
	r.Equal(e.AskPrice, a.AskPrice, "AskQuantity")
	r.Equal(e.AskQuantity, a.AskQuantity, "AskQuantity")

}

package spot

import (
	"github.com/shopspring/decimal"
)

const (
	DEPTH_SIZE_FIVE   = 5
	DEPTH_SIZE_TEN    = 10
	DEPTH_SIZE_TWENTY = 20
)

const (
	STEP0 = "step0"
	STEP1 = "step1"
	STEP2 = "step2"
	STEP3 = "step3"
	STEP4 = "step4"
	STEP5 = "step5"
)

type GetDepthOptionalRequest struct {
	Size int
}

const (
	MIN1  = "1min"
	MIN5  = "5min"
	MIN15 = "15min"
	MIN30 = "30min"
	MIN60 = "60min"
	HOUR4 = "4hour"
	DAY1  = "1day"
	MON1  = "1mon"
	WEEK1 = "1week"
	YEAR1 = "1year"
)

type GetCandlestickOptionalRequest struct {
	Period string
	Size   int
}

type GetCandlestickResponse struct {
	Status string        `json:"status"`
	Ch     string        `json:"ch"`
	Ts     int64         `json:"ts"`
	Data   []Candlestick `json:"data"`
}

type Candlestick struct {
	Amount decimal.Decimal `json:"amount"`
	Open   decimal.Decimal `json:"open"`
	Close  decimal.Decimal `json:"close"`
	High   decimal.Decimal `json:"high"`
	Id     int64           `json:"id"`
	Count  int64           `json:"count"`
	Low    decimal.Decimal `json:"low"`
	Vol    decimal.Decimal `json:"vol"`
}

type GetDepthResponse struct {
	Status string `json:"status"`
	Ch     string `json:"ch"`
	Ts     int64  `json:"ts"`
	Tick   *Depth `json:"tick"`
}

type Depth struct {
	Timestamp int64               `json:"ts"`
	Version   int64               `json:"version"`
	Bids      [][]decimal.Decimal `json:"bids"`
	Asks      [][]decimal.Decimal `json:"asks"`
}

type GetHistoricalTradeOptionalRequest struct {
	Size int
}

type GetHistoricalTradeResponse struct {
	Status string      `json:"status"`
	Ch     string      `json:"ch"`
	Ts     int64       `json:"ts"`
	Data   []TradeTick `json:"data"`
}

type GetLast24hCandlestick struct {
	Status string       `json:"status"`
	Ch     string       `json:"ch"`
	Ts     int64        `json:"ts"`
	Tick   *Candlestick `json:"tick"`
}

type GetLast24hCandlestickAskBidResponse struct {
	Status string             `json:"status"`
	Ch     string             `json:"ch"`
	Ts     int64              `json:"ts"`
	Tick   *CandlestickAskBid `json:"tick"`
}
type CandlestickAskBid struct {
	Amount  decimal.Decimal   `json:"amount"`
	Open    decimal.Decimal   `json:"open"`
	Close   decimal.Decimal   `json:"close"`
	High    decimal.Decimal   `json:"high"`
	Id      int64             `json:"id"`
	Count   int64             `json:"count"`
	Low     decimal.Decimal   `json:"low"`
	Vol     decimal.Decimal   `json:"vol"`
	Version int64             `json:"version"`
	Bid     []decimal.Decimal `json:"bid"`
	Ask     []decimal.Decimal `json:"ask"`
}

type GetAllSymbolsLast24hCandlesticksAskBidResponse struct {
	Status string              `json:"status"`
	Ts     int64               `json:"ts"`
	Data   []SymbolCandlestick `json:"data"`
}
type SymbolCandlestick struct {
	Amount  decimal.Decimal `json:"amount"`
	Open    decimal.Decimal `json:"open"`
	Close   decimal.Decimal `json:"close"`
	High    decimal.Decimal `json:"high"`
	Symbol  string          `json:"symbol"`
	Count   int64           `json:"count"`
	Low     decimal.Decimal `json:"low"`
	Vol     decimal.Decimal `json:"vol"`
	Bid     decimal.Decimal `json:"bid"`
	BidSize decimal.Decimal `json:"bidSize"`
	Ask     decimal.Decimal `json:"ask"`
	AskSize decimal.Decimal `json:"askSize"`
}

type GetLatestTradeResponse struct {
	Status string     `json:"status"`
	Ch     string     `json:"ch"`
	Ts     int64      `json:"ts"`
	Tick   *TradeTick `json:"tick"`
}
type TradeTick struct {
	Id   int64 `json:"id"`
	Ts   int64 `json:"ts"`
	Data []struct {
		Amount    decimal.Decimal `json:"amount"`
		TradeId   int64           `json:"trade-id"`
		Ts        int64           `json:"ts"`
		Id        decimal.Decimal `json:"id"`
		Price     decimal.Decimal `json:"price"`
		Direction string          `json:"direction"`
	}
}

type SubscribeBestBidOfferResponse struct {
	WebSocketResponseBase
	Tick *struct {
		QuoteTime int64           `json:"quoteTime"`
		Symbol    string          `json:"symbol"`
		Bid       decimal.Decimal `json:"bid"`
		BidSize   decimal.Decimal `json:"bidSize"`
		Ask       decimal.Decimal `json:"ask"`
		AskSize   decimal.Decimal `json:"askSize"`
	}
}

type SubscribeCandlestickResponse struct {
	base WebSocketResponseBase
	Tick *Tick
	Data []Tick
}

type Tick struct {
	Id     int64           `json:"id"`
	Amount decimal.Decimal `json:"amount"`
	Count  int             `json:"count"`
	Open   decimal.Decimal `json:"open"`
	Close  decimal.Decimal `json:"close"`
	Low    decimal.Decimal `json:"low"`
	High   decimal.Decimal `json:"high"`
	Vol    decimal.Decimal `json:"vol"`
}

type SubscribeDepthResponse struct {
	WebSocketResponseBase
	Data *Depth
	Tick *Depth
}

type SubscribeLast24hCandlestickResponse struct {
	WebSocketResponseBase
	Data *Candlestick
	Tick *Candlestick
}

type SubscribeMarketByPriceResponse struct {
	WebSocketResponseBase
	Tick *MarketByPrice
	Data *MarketByPrice
}

type MarketByPrice struct {
	SeqNum     int64               `json:"seqNum"`
	PrevSeqNum int64               `json:"prevSeqNum"`
	Bids       [][]decimal.Decimal `json:"bids"`
	Asks       [][]decimal.Decimal `json:"asks"`
}

type SubscribeTradeResponse struct {
	WebSocketResponseBase
	Tick *struct {
		Id        int64 `json:"id"`
		Timestamp int64 `json:"ts"`
		Data      []Trade
	}
}

type Trade struct {
	TradeId   int64           `json:"tradeId"`
	Amount    decimal.Decimal `json:"amount"`
	Price     decimal.Decimal `json:"price"`
	Timestamp int64           `json:"ts"`
	Direction string          `json:"direction"`
}

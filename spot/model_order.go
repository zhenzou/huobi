package spot

import (
	"github.com/shopspring/decimal"
)

type CancelOrderByClientResponse struct {
	Status       string `json:"status"`
	Data         int    `json:"data"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
}

type CancelOrderByIdResponse struct {
	Status       string `json:"status"`
	Data         string `json:"data"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	OrderState   string `json:"order-state"`
}

type CancelOrdersByCriteriaResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         *struct {
		SuccessCount int `json:"success-count"`
		FailedCount  int `json:"failed-count"`
		NextId       int `json:"next-id"`
	}
}

type CancelOrdersByCriteriaRequest struct {
	AccountId string `json:"account-id"`
	Symbol    string `json:"symbol,omitempty"`
	Side      string `json:"side,omitempty"`
	Size      int    `json:"size,omitempty"`
}

type CancelOrdersByIdsRequest struct {
	OrderIds       []string `json:"order-ids,omitempty"`
	ClientOrderIds []string `json:"client-order-ids,omitempty"`
}

type CancelOrdersByIdsResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         *struct {
		Success []string `json:"success"`
		Failed  []struct {
			OrderId       string `json:"order-id"`
			ClientOrderId string `json:"client-order-id"`
			ErrorCode     string `json:"err-code"`
			ErrorMessage  string `json:"err-msg"`
		}
	}
}

type GetHistoryOrdersResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         []struct {
		Id               int64  `json:"id"`
		ClientOrderId    string `json:"client-order-id"`
		AccountId        int    `json:"account-id"`
		UserId           int    `json:"user-id"`
		Amount           string `json:"amount"`
		Symbol           string `json:"symbol"`
		Price            string `json:"price"`
		CreatedAt        int64  `json:"created-at"`
		CanceledAt       int64  `json:"canceled-at"`
		FinishedAt       int64  `json:"finished-at"`
		Type             string `json:"type"`
		FilledAmount     string `json:"field-amount"`
		FilledCashAmount string `json:"field-cash-amount"`
		FilledFees       string `json:"field-fees"`
		Source           string `json:"source"`
		State            string `json:"state"`
		Exchange         string `json:"exchange"`
		Batch            string `json:"batch"`
		StopPrice        string `json:"stop-price"`
		Operator         string `json:"operator"`
	}
}

type GetMatchResultsResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         []struct {
		Id                int64  `json:"id"`
		OrderId           int64  `json:"order-id"`
		MatchId           int64  `json:"match-id"`
		TradeId           int64  `json:"trade-id"`
		Symbol            string `json:"symbol"`
		Price             string `json:"price"`
		CreatedAt         int64  `json:"created-at"`
		Type              string `json:"type"`
		FilledAmount      string `json:"filled-amount"`
		FilledFees        string `json:"filled-fees"`
		FeeCurrency       string `json:"fee-currency"`
		Source            string `json:"source"`
		Role              string `json:"role"`
		FilledPoints      string `json:"filled-points"`
		FeeDeductCurrency string `json:"fee-deduct-currency"`
		FeeDeductState    string `json:"fee-deduct-state"`
	}
}

type GetOpenOrdersResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         []OpenOrder
}

type OpenOrder struct {
	Id               int64           `json:"id"`
	ClientOrderId    string          `json:"client-order-id"`
	AccountId        int             `json:"account-id"`
	Amount           decimal.Decimal `json:"amount"`
	Symbol           string          `json:"symbol"`
	Price            decimal.Decimal `json:"price"`
	CreatedAt        int64           `json:"created-at"`
	Type             string          `json:"type"`
	FilledAmount     decimal.Decimal `json:"filled-amount"`
	FilledCashAmount decimal.Decimal `json:"filled-cash-amount"`
	FilledFees       decimal.Decimal `json:"filled-fees"`
	Source           string          `json:"source"`
	State            string          `json:"state"`
	StopPrice        decimal.Decimal `json:"stop-price"`
	Operator         string          `json:"operator"`
}

type GetOrderResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         *struct {
		AccountId        int    `json:"account-id"`
		Amount           string `json:"amount"`
		Id               int64  `json:"id"`
		ClientOrderId    string `json:"client-order-id"`
		Symbol           string `json:"symbol"`
		Price            string `json:"price"`
		CreatedAt        int64  `json:"created-at"`
		Type             string `json:"type"`
		FilledAmount     string `json:"field-amount"`
		FilledCashAmount string `json:"field-cash-amount"`
		FilledFees       string `json:"field-fees"`
		Source           string `json:"source"`
		State            string `json:"state"`
	}
}

type GetTransactFeeRateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Symbol          string `json:"symbol"`
		MakerFeeRate    string `json:"makerFeeRate"`
		TakerFeeRate    string `json:"takerFeeRate"`
		ActualMakerRate string `json:"actualMakerRate"`
		ActualTakerRate string `json:"actualTakerRate"`
	}
}

type PlaceOrderRequest struct {
	AccountId     string `json:"account-id"`
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
	Amount        string `json:"amount"`
	Price         string `json:"price,omitempty"`
	Source        string `json:"source,omitempty"`
	ClientOrderId string `json:"client-order-id,omitempty"`
	StopPrice     string `json:"stop-price,omitempty"`
	Operator      string `json:"operator,omitempty"`
}

type PlaceOrderResponse struct {
	Status       string `json:"status"`
	Data         string `json:"data"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
}

type PlaceOrdersResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         []PlaceOrderResult
}

type PlaceOrderResult struct {
	OrderId       int64  `json:"order-id"`
	ClientOrderId string `json:"client-order-id"`
	ErrorCode     string `json:"err-code"`
	ErrorMessage  string `json:"err-msg"`
}

type RequestOrdersRequest struct {
	Op        string `json:"op"`
	Topic     string `json:"topic"`
	ClientId  string `json:"cid"`
	AccountId int    `json:"account-id"`
	Symbol    string `json:"symbol"`
	Types     string `json:"types"`
	States    string `json:"states"`
	StartDate string `json:"start-date"`
	EndDate   string `json:"end-date"`
	From      string `json:"from"`
	Direct    string `json:"direct"`
	Size      string `json:"size"`
}

type RequestOrdersV1Response struct {
	Op        string  `json:"op"`
	Timestamp int64   `json:"ts"`
	Topic     string  `json:"topic"`
	ClientId  string  `json:"cid"`
	ErrorCode int     `json:"err-code"`
	Data      []Order `json:"data"`
}

type RequestOrderV1Response struct {
	Op        string `json:"op"`
	Timestamp int64  `json:"ts"`
	Topic     string `json:"topic"`
	ClientId  string `json:"cid"`
	ErrorCode int    `json:"err-code"`
	Data      Order  `json:"data"`
}

type Order struct {
	AccountId        int    `json:"account-id"`
	Amount           string `json:"amount"`
	Id               int64  `json:"id"`
	Symbol           string `json:"symbol"`
	Price            string `json:"price"`
	CreatedAt        int64  `json:"created-at"`
	Type             string `json:"type"`
	FilledAmount     string `json:"filled-amount"`
	FilledCashAmount string `json:"filled-cash-amount"`
	FilledFees       string `json:"filled-fees"`
	Source           string `json:"source"`
	State            string `json:"state"`
}

type SubscribeOrderV1Response struct {
	Op        string `json:"op"`
	Timestamp int64  `json:"ts"`
	Topic     string `json:"topic"`
	Data      struct {
		MatchId          int    `json:"match-id"`
		OrderId          int64  `json:"order-id"`
		Symbol           string `json:"symbol"`
		OrderState       string `json:"order-state"`
		Role             string `json:"role"`
		Price            string `json:"price"`
		FilledAmount     string `json:"filled-amount"`
		FilledCashAmount string `json:"filled-cash-amount"`
		UnfilledAmount   string `json:"unfilled-amount"`
		ClientOrderId    string `json:"client-order-id"`
		OrderType        string `json:"order-type"`
	}
}

type SubscribeOrderV2Response struct {
	WebSocketV2ResponseBase
	Data *struct {
		EventType       string `json:"eventType"`
		Symbol          string `json:"symbol"`
		AccountId       int64  `json:"accountId"`
		OrderId         int64  `json:"orderId"`
		ClientOrderId   string `json:"clientOrderId"`
		OrderSide       string `json:"orderSide"`
		OrderPrice      string `json:"orderPrice"`
		OrderSize       string `json:"orderSize"`
		OrderValue      string `json:"orderValue"`
		Type            string `json:"type"`
		OrderStatus     string `json:"orderStatus"`
		OrderCreateTime int64  `json:"orderCreateTime"`
		TradePrice      string `json:"tradePrice"`
		TradeVolume     string `json:"tradeVolume"`
		TradeId         int64  `json:"tradeId"`
		TradeTime       int64  `json:"tradeTime"`
		Aggressor       bool   `json:"aggressor"`
		RemainAmt       string `json:"remainAmt"`
		LastActTime     int64  `json:"lastActTime"`
		ErrorCode       int    `json:"errCode"`
		ErrorMessage    string `json:"errMessage"`
	}
}

type SubscribeTradeClearResponse struct {
	WebSocketV2ResponseBase
	Data *struct {
		EventType       string `json:"eventType"`
		Symbol          string `json:"symbol"`
		OrderId         int64  `json:"orderId"`
		TradePrice      string `json:"tradePrice"`
		TradeVolume     string `json:"tradeVolume"`
		OrderSide       string `json:"orderSide"`
		OrderType       string `json:"orderType"`
		Aggressor       bool   `json:"aggressor"`
		TradeId         int64  `json:"tradeId"`
		TradeTime       int64  `json:"tradeTime"`
		TransactFee     string `json:"transactFee"`
		FeeCurrency     string `json:"feeCurrency"`
		FeeDeduct       string `json:"feeDeduct"`
		FeeDeductType   string `json:"feeDeductType"`
		AccountId       int64  `json:"accountId"`
		Source          string `json:"source"`
		OrderPrice      string `json:"orderPrice"`
		OrderSize       string `json:"orderSize"`
		OrderValue      string `json:"orderValue"`
		ClientOrderId   string `json:"clientOrderId"`
		StopPrice       string `json:"stopPrice"`
		Operator        string `json:"operator"`
		OrderCreateTime int64  `json:"orderCreateTime"`
		OrderStatus     string `json:"orderStatus"`
	}
}

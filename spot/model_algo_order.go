package spot

type CancelOrdersRequest struct {
	ClientOrderIds []string `json:"clientOrderIds"`
}

type CancelOrdersResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Accepted []string `json:"accepted"`
		Rejected []string `json:"rejected"`
	}
}
type GetAlgoHistoryOrdersResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		AccountId       int    `json:"accountId"`
		Source          string `json:"source"`
		ClientOrderId   string `json:"clientOrderId"`
		Symbol          string `json:"symbol"`
		OrderPrice      string `json:"orderPrice"`
		OrderSize       string `json:"orderSize"`
		OrderValue      string `json:"orderValue"`
		OrderSide       string `json:"orderSide"`
		TimeInForce     string `json:"timeInForce"`
		OrderType       string `json:"orderType"`
		StopPrice       string `json:"stopPrice"`
		TrailingRate    string `json:"trailingRate"`
		OrderOrigTime   int64  `json:"orderOrigTime"`
		LastActTime     int64  `json:"lastActTime"`
		OrderCreateTime int64  `json:"orderCreateTime"`
		OrderStatus     string `json:"orderStatus"`
		ErrCode         int    `json:"errCode"`
		ErrMessage      string `json:"errMessage"`
	}
	NextId int64 `json:"nextId"`
}

type GetAlgoOpenOrdersResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		AccountId     int    `json:"accountId"`
		Source        string `json:"source"`
		ClientOrderId string `json:"clientOrderId"`
		Symbol        string `json:"symbol"`
		OrderPrice    string `json:"orderPrice"`
		OrderSize     string `json:"orderSize"`
		OrderValue    string `json:"orderValue"`
		OrderSide     string `json:"orderSide"`
		TimeInForce   string `json:"timeInForce"`
		OrderType     string `json:"orderType"`
		StopPrice     string `json:"stopPrice"`
		TrailingRate  string `json:"trailingRate"`
		OrderOrigTime int64  `json:"orderOrigTime"`
		LastActTime   int64  `json:"lastActTime"`
		OrderStatus   string `json:"orderStatus"`
	}
	NextId int64 `json:"nextId"`
}

type GetSpecificOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		AccountId       int    `json:"accountId"`
		Source          string `json:"source"`
		ClientOrderId   string `json:"clientOrderId"`
		OrderId         string `json:"orderId"`
		Symbol          string `json:"symbol"`
		OrderPrice      string `json:"orderPrice"`
		OrderSize       string `json:"orderSize"`
		OrderValue      string `json:"orderValue"`
		OrderSide       string `json:"orderSide"`
		TimeInForce     string `json:"timeInForce"`
		OrderType       string `json:"orderType"`
		StopPrice       string `json:"stopPrice"`
		TrailingRate    string `json:"trailingRate"`
		OrderOrigTime   int64  `json:"orderOrigTime"`
		LastActTime     int64  `json:"lastActTime"`
		OrderCreateTime int64  `json:"orderCreateTime"`
		OrderStatus     string `json:"orderStatus"`
		ErrCode         int    `json:"errCode"`
		ErrMessage      string `json:"errMessage"`
	}
}

type PlaceAlgoOrderRequest struct {
	AccountId     int    `json:"accountId"`
	Symbol        string `json:"symbol"`
	OrderPrice    string `json:"orderPrice"`
	OrderSide     string `json:"orderSide"`
	OrderSize     string `json:"orderSize"`
	OrderValue    string `json:"orderValue"`
	TimeInForce   string `json:"timeInForce"`
	OrderType     string `json:"orderType"`
	ClientOrderId string `json:"clientOrderId"`
	StopPrice     string `json:"stopPrice"`
	TrailingRate  string `json:"trailingRate"`
}

type PlaceAlgoOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ClientOrderId string `json:"clientOrderId"`
	}
}

package futures

type ContractInfo struct {
	Symbol string `json:"symbol"`

	ContractCode string `json:"contract_code"`

	ContractSize float64 `json:"contract_size"`

	PriceTick float64 `json:"price_tick"`

	SettlementDate string `json:"settlement_date"`

	CreateDate string `json:"create_date"`

	DeliveryTime string `json:"delivery_time"`

	ContractStatus int `json:"contract_status"`

	SupportMarginMode string `json:"support_margin_mode"`
}

type ReqKLineResponse struct {
	Rep string `json:"rep"`

	Status string `json:"status"`

	Id string `json:"id"`

	Wsid int64 `json:"wsid"`

	Data []struct {
		Id int64 `json:"id"`

		Vol float64 `json:"vol"`

		Count float64 `json:"count"`

		Open float64 `json:"open"`

		Close float64 `json:"close"`

		Low float64 `json:"low"`

		High float64 `json:"high"`

		Amount float64 `json:"amount"`

		TradeTurnover float64 `json:"trade_turnover"`
	} `json:"data"`
}

type ReqTradeDetailResponse struct {
	Rep string `json:"rep"`

	Status string `json:"status"`

	Id string `json:"id"`

	Ts int64 `json:"ts"`

	Data []struct {
		Id int64 `json:"id"`

		Price string `json:"price"`

		Amount string `json:"amount"`

		Direction string `json:"direction"`

		Ts int64 `json:"ts"`
	} `json:"data"`
}

type SubBBOResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Ch string `json:"ch"`

		Mrid int64 `json:"mrid"`

		Id int64 `json:"id"`

		Ask []float64 `json:"ask"`

		Did []float64 `json:"bid"`

		Version int64 `json:"version"`

		Ts int64 `json:"ts"`
	} `json:"tick"`
}
type SubDepthResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Mrid int64 `json:"mrid"`

		Id int64 `json:"id"`

		Asks [][]float64 `json:"asks"`

		Bids [][]float64 `json:"bids"`

		Ts int64 `json:"ts"`

		Version int64 `json:"version"`

		Ch string `json:"ch"`

		TickEvent string `json:"event,omitempty"`
	} `json:"tick"`
}

type SubKLineResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Id int64 `json:"id"`

		Mrid int64 `json:"mrid"`

		Vol float64 `json:"vol"`

		Count float64 `json:"count"`

		Open float64 `json:"open"`

		Close float64 `json:"close"`

		Low float64 `json:"low"`

		High float64 `json:"high"`

		Amount float64 `json:"amount"`

		TradeTurnover float64 `json:"trade_turnover"`

		Ask []float64 `json:"ask,omitempty"`

		Bid []float64 `json:"bid,omitempty"`
	} `json:"tick"`
}

type SubTradeDetailResponse struct {
	Ch   string `json:"ch"`
	Ts   int64  `json:"ts"`
	Tick struct {
		Id   int64 `json:"id"`
		Ts   int64 `json:"ts"`
		Data []struct {
			Amount    float64 `json:"amount"`
			Ts        int64   `json:"ts"`
			Id        int64   `json:"id"`
			Price     float64 `json:"price"`
			Direction string  `json:"direction"`
		} `json:"data"`
	} `json:"tick"`
}

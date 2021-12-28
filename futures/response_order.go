package futures

type SubContractInfoResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	EventSender string `json:"event"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		ContractSize float64 `json:"contract_size"`

		PriceTick float64 `json:"price_tick"`

		SettlementDate string `json:"settlement_date"`

		CreateDate string `json:"create_date"`

		ContractStatus int `json:"contract_status"`
	} `json:"data"`
}

type SubFundingRateResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		FeeAsset string `json:"fee_asset"`

		FundingTime string `json:"funding_time"`

		FundingRate string `json:"funding_rate"`

		EstimatedRate string `json:"estimated_rate"`

		SettlementTime string `json:"settlement_time"`
	} `json:"data"`
}

type SubLiquidationOrdersResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Direction string `json:"direction"`

		Offset string `json:"offset"`

		Volume float64 `json:"volume"`

		Price float64 `json:"price"`

		CreatedAt int64 `json:"created_at"`
	} `json:"data"`
}

type SubOrdersResponse struct {
	Op              string  `json:"op"`
	Topic           string  `json:"topic"`
	Ts              int64   `json:"ts"`
	Uid             string  `json:"uid"`
	Symbol          string  `json:"symbol"`
	ContractCode    string  `json:"contract_code"`
	Volume          float64 `json:"volume"`
	Price           float64 `json:"price"`
	OrderPriceType  string  `json:"order_price_type"`
	Direction       string  `json:"direction"`
	Offset          string  `json:"offset"`
	Status          int     `json:"status"`
	LeverRate       int     `json:"lever_rate"`
	OrderId         string  `json:"order_id_str"`
	ClientOrderId   int64   `json:"client_order_id,omitempty"`
	OrderSource     string  `json:"order_source"`
	OrderType       int     `json:"order_type"`
	CreatedAt       int64   `json:"created_at"`
	TradeVolume     float64 `json:"trade_volume"`
	TradeTurnover   float64 `json:"trade_turnover"`
	Fee             float64 `json:"fee"`
	TradeAvgPrice   float64 `json:"trade_avg_price"`
	MarginAsset     string  `json:"margin_asset"`
	MarginFrozen    float64 `json:"margin_frozen"`
	Profit          float64 `json:"profit"`
	RealProfit      float64 `json:"real_profit"`
	LiquidationType string  `json:"liquidation_type"`
	CanceledAt      int64   `json:"canceled_at"`
	FeeAsset        string  `json:"fee_asset"`
	IsTpsl          int     `json:"is_tpsl"`
	Trade           []struct {
		Id            string  `json:"id"`
		TradeId       int64   `json:"trade_id"`
		TradeVolume   float64 `json:"trade_volume"`
		TradePrice    float64 `json:"trade_price"`
		TradeFee      float64 `json:"trade_fee"`
		TradeTurnover float64 `json:"trade_turnover"`
		CreatedAt     int64   `json:"created_at"`
		Role          string  `json:"role"`
		FeeAsset      string  `json:"fee_asset"`
	} `json:"trade"`
}

type SubPositionsResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Uid string `json:"uid"`

	EventSender string `json:"event"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Volume float64 `json:"volume"`

		Available float64 `json:"available"`

		Frozen float64 `json:"frozen"`

		CostOpen float64 `json:"cost_open"`

		CostHold float64 `json:"cost_hold"`

		ProfitUnreal float64 `json:"profit_unreal"`

		ProfitRate float64 `json:"profit_rate"`

		Profit float64 `json:"profit"`

		PositionMargin float64 `json:"position_margin"`

		LeverRate int `json:"lever_rate"`

		Direction string `json:"direction"`

		LastPrice float64 `json:"last_price"`

		MarginAsset string `json:"margin_asset"`
	} `json:"data"`
}

type SubTriggerOrderResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Uid string `json:"uid"`

	EventSender string `json:"event"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		TriggerType string `json:"trigger_type"`

		Volume float64 `json:"volume"`

		OrderType int `json:"order_type"`

		Direction string `json:"direction"`

		Offset string `json:"offset"`

		LeverRate int `json:"lever_rate"`

		OrderId int64 `json:"order_id"`

		OrderIdStr string `json:"order_id_str"`

		RelationOrderId string `json:"relation_order_id"`

		OrderPriceType string `json:"order_price_type"`

		Status int `json:"status"`

		OrderSource string `json:"order_source"`

		TriggerPrice float64 `json:"trigger_price"`

		TriggeredPrice float64 `json:"triggered_price,omitempty"`

		OrderPrice float64 `json:"order_price"`

		CreatedAt int64 `json:"created_at"`

		TriggeredAt int64 `json:"triggered_at,omitempty"`

		OrderInsertAt int64 `json:"order_insert_at"`

		CanceledAt int64 `json:"canceled_at,omitempty"`

		FailCode int `json:"fail_code,omitempty"`

		FailReason string `json:"fail_reason,omitempty"`
	} `json:"data"`
}

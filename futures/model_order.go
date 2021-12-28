package futures

const (
	OrderPriceTypeLimit     = "limit"
	OrderPriceTypeOpponent  = "opponent"
	OrderPriceTypeOptimal5  = "optimal_5"
	OrderPriceTypeOptimal10 = "optimal_10"
	OrderPriceTypeOptimal20 = "optimal_20"
)

const (
	OrderStatusPendingSubmit1           = 1
	OrderStatusPendingSubmit2           = 2
	OrderStatusSubmitted                = 3
	OrderStatusPartialTraded            = 4
	OrderStatusPartialTradedAndCanceled = 5
	OrderStatusAllTraded                = 6
	OrderStatusCanceled                 = 7
)

type GetHisMatchResponse struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"err_code,omitempty"`
	ErrorMessage string `json:"err_msg,omitempty"`
	Data         struct {
		Trades []struct {
			Id               string  `json:"id"`
			MatchId          int64   `json:"match_id"`
			OrderId          int64   `json:"order_id"`
			OrderIdStr       string  `json:"order_id_str"`
			Symbol           string  `json:"symbol"`
			OrderSource      string  `json:"order_source"`
			ContractCode     string  `json:"contract_code"`
			Direction        string  `json:"direction"`
			Offset           string  `json:"offset"`
			TradeVolume      float64 `json:"trade_volume"`
			TradePrice       float64 `json:"trade_price"`
			TradeTurnover    float64 `json:"trade_turnover"`
			CreateDate       int64   `json:"create_date"`
			OffsetProfitloss float64 `json:"offset_profitloss"`
			TradeFee         float64 `json:"trade_fee"`
			Role             string  `json:"role"`
			FeeAsset         string  `json:"fee_asset"`
			MarginMode       string  `json:"margin_mode"`
			MarginAccount    string  `json:"margin_account"`
		} `json:"trades"`
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}

type Order struct {
	OrderId      string  `json:"order_id_str"`
	Symbol       string  `json:"symbol"`
	ContractCode string  `json:"contract_code"`
	Volume       float64 `json:"volume"`
	Price        float64 `json:"price"`
	// 1：限价单（limit）
	//2：市价单（market）
	//3：对手价（opponent）
	//4：闪电平仓（lightning）
	//5：计划委托（trigger）
	//6：post_only
	//7：最优5档（optimal_5）
	//8：最优10档（optimal_10）
	//9：最优20档（optimal_20）
	//10：FOK
	//11：IOC
	//12：对手价_IOC（opponent_ioc）
	//13：闪电平仓_IOC（lightning_ioc)
	//14：最优5档_IOC（optimal_5_ioc）
	//15：最优10档_IOC（optimal_10_ioc）
	//16：最优20档_IOC（optimal_20_ioc）
	//17：对手价_FOK（opponent_fok）
	//18：闪电平仓_FOK（lightning_fok）
	//19：最优5档_FOK（optimal_5_fok）
	//40：最优10档_FOK（optimal_10_fok）
	//41：最优20档_FOK（optimal_20_fok）。
	OrderPriceType  int64   `json:"order_price_type"`
	OrderType       int     `json:"order_type"`
	Direction       string  `json:"direction"`
	Offset          string  `json:"offset"`
	LeverRate       int     `json:"lever_rate"`
	ClientOrderId   int64   `json:"client_order_id,omitempty"`
	TradeVolume     float64 `json:"trade_volume"`
	TradeTurnover   float64 `json:"trade_turnover"`
	Fee             float64 `json:"fee"`
	FeeAsset        string  `json:"fee_asset"`
	TradeAvgPrice   float64 `json:"trade_avg_price,omitempty"`
	MarginFrozen    float64 `json:"margin_frozen"`
	MarginAsset     string  `json:"margin_asset"`
	Profit          float64 `json:"profit"`
	RealProfit      float64 `json:"real_profit"`
	Status          int     `json:"status"`
	OrderSource     string  `json:"order_source"`
	LiquidationType string  `json:"liquidation_type"`
	CanceledAt      int64   `json:"canceled_at,omitempty"`
	MarginMode      string  `json:"margin_mode"`
	MarginAccount   string  `json:"margin_account"`
	CreateDate      int64   `json:"create_date"`
	UpdateDate      int64   `json:"update_date"`
	IsTpsl          int     `json:"is_tpsl"`
}

type GetOrderInfoResponse struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"err_code,omitempty"`
	ErrorMessage string `json:"err_msg,omitempty"`
	Data         []struct {
		Symbol          string  `json:"symbol"`
		ContractCode    string  `json:"contract_code"`
		Volume          float64 `json:"volume"`
		Price           float64 `json:"price"`
		OrderPriceType  string  `json:"order_price_type"`
		Direction       string  `json:"direction"`
		Offset          string  `json:"offset"`
		LeverRate       int     `json:"lever_rate"`
		OrderId         int64   `json:"order_id"`
		OrderIdStr      string  `json:"order_id_str"`
		ClientOrderId   int64   `json:"client_order_id,omitempty"`
		CreatedAt       int64   `json:"created_at"`
		TradeVolume     float64 `json:"trade_volume"`
		TradeTurnover   float64 `json:"trade_turnover"`
		Fee             float64 `json:"fee"`
		TradeAvgPrice   float64 `json:"trade_avg_price"`
		MarginAsset     string  `json:"margin_asset"`
		MarginFrozen    float64 `json:"margin_frozen"`
		Profit          float64 `json:"profit"`
		Status          int     `json:"status"`
		OrderType       int     `json:"order_type"`
		OrderSource     string  `json:"order_source"`
		FeeAsset        string  `json:"fee_asset"`
		LiquidationType string  `json:"liquidation_type"`
		CanceledAt      int64   `json:"canceled_at"`
		MarginMode      string  `json:"margin_mode"`
		MarginAccount   string  `json:"margin_account"`
		IsTpsl          int     `json:"is_tpsl"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetOrderDetailResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol          string  `json:"symbol"`
		ContractCode    string  `json:"contract_code"`
		LeverRate       int     `json:"lever_rate"`
		Direction       string  `json:"direction"`
		Offset          string  `json:"offset"`
		Volume          float64 `json:"volume"`
		Price           float64 `json:"price"`
		CreatedAt       int64   `json:"created_at"`
		CanceledAt      int64   `json:"canceled_at"`
		OrderSource     string  `json:"order_source"`
		OrderPriceType  string  `json:"order_price_type"`
		MarginAsset     string  `json:"margin_asset"`
		MarginFrozen    float64 `json:"margin_frozen"`
		Profit          float64 `json:"profit"`
		InstrumentPrice float64 `json:"instrument_price"`
		FinalInterest   float64 `json:"final_interest"`
		AdjustValue     float64 `json:"adjust_value"`
		Fee             float64 `json:"fee"`
		FeeAsset        string  `json:"fee_asset"`
		LiquidationType string  `json:"liquidation_type"`
		OrderId         int64   `json:"order_id"`
		OrderIdStr      string  `json:"order_id_str"`
		ClientOrderId   int64   `json:"client_order_id,omitempty"`
		OrderType       string  `json:"order_type"`
		Status          int     `json:"status"`
		TradeAvgPrice   float64 `json:"trade_avg_price,omitempty"`
		TradeTurnover   float64 `json:"trade_turnover"`
		TradeVolume     float64 `json:"trade_volume"`
		TotalPage       int     `json:"total_page"`
		CurrentPage     int     `json:"current_page"`
		TotalSize       int     `json:"total_size"`
		MarginMode      string  `json:"margin_mode"`
		MarginAccount   string  `json:"margin_account"`
		IsTpsl          int     `json:"is_tpsl"`
		Trades          []struct {
			Id string `json:"id"`

			TradeId int64 `json:"trade_id"`

			TradePrice float64 `json:"trade_price"`

			TradeVolume float64 `json:"trade_volume"`

			TradeTurnover float64 `json:"trade_turnover"`

			TradeFee float64 `json:"trade_fee"`

			Role string `json:"role"`

			CreatedAt int64 `json:"created_at"`
		} `json:"trades,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type LightningCloseResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		OrderId int64 `json:"order_id"`

		OrderIdStr string `json:"order_id_str"`

		ClientOrderId int64 `json:"client_order_id,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type PlaceBatchOrderResponse struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"err_code,omitempty"`
	ErrorMessage string `json:"err_msg,omitempty"`
	Data         struct {
		Errors []struct {
			Index        int    `json:"index"`
			ErrorCode    int    `json:"err_code"`
			ErrorMessage string `json:"err_msg"`
		} `json:"errors,omitempty"`

		Success []struct {
			Index         int    `json:"index"`
			OrderId       int64  `json:"order_id"`
			ClientOrderId int64  `json:"client_order_id,omitempty"`
			OrderIdStr    string `json:"order_id_str"`
		} `json:"success,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type SwitchLeverRateResponse struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"err_code,omitempty"`
	ErrorMessage string `json:"err_msg,omitempty"`
	Data         struct {
		ContractCode string `json:"contract_code"`
		MarginMode   string `json:"margin_mode"`
		LeverRate    int    `json:"lever_rate"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}

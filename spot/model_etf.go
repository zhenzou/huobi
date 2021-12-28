package spot

import "github.com/shopspring/decimal"

type GetSwapConfigResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    *SwapConfig `json:"data"`
}

type SwapConfig struct {
	PurchaseMinAmount   int64           `json:"purchase_min_amount"`
	PurchaseMaxAmount   int64           `json:"purchase_max_amount"`
	RedemptionMinAmount int64           `json:"redemption_min_amount"`
	RedemptionMaxAmount int64           `json:"redemption_max_amount"`
	PurchaseFeeRate     decimal.Decimal `json:"purchase_fee_rate"`
	RedemptionFeeRate   decimal.Decimal `json:"redemption_fee_rate"`
	EtfStatus           int             `json:"etf_status"`
	Unitprice           []*UnitPrice    `json:"unit_price"`
}
type UnitPrice struct {
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
}

type GetSwapListResponse struct {
	Code int         `json:"code"`
	Data []*SwapList `json:"data"`
}

type SwapList struct {
	Id         int64           `json:"id"`
	GmtCreated int64           `json:"gmt_created"`
	Currency   string          `json:"currency"`
	Amount     decimal.Decimal `json:"amount"`
	Type       int             `json:"type"`
	Status     int             `json:"status"`
	Detail     *struct {
		UsedCurrencyList   []*UnitPrice    `json:"used_ currency_list"`
		Rate               decimal.Decimal `json:"rate"`
		Fee                decimal.Decimal `json:"fee"`
		PointCardAmount    decimal.Decimal `json:"point_card_amount"`
		ObtainCurrencyList []*UnitPrice    `json:"obtain_currency_list"`
	}
}

type SwapRequest struct {
	EtfName string `json:"etf_name"`
	Amount  int64  `json:"amount"`
}

type SwapResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

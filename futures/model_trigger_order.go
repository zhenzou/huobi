package futures

import (
	"strconv"
)

type TriggerOrderStatus int64

func (t TriggerOrderStatus) String() string {
	return strconv.FormatInt(int64(t), 10)
}

const (
	TriggerOrderStatusPendingSubmit = 1
	TriggerOrderStatusSubmitted     = 2
	TriggerOrderStatusMatching      = 3
	TriggerOrderStatusSucceed       = 4
	TriggerOrderStatusFailed        = 5
	TriggerOrderStatusCanceled      = 6
)

type TriggerOrder struct {
	OrderId             int64   `json:"order_id"`
	Symbol              string  `json:"symbol"`
	ContractCode        string  `json:"contract_code"`
	TriggerType         string  `json:"trigger_type"`
	Volume              float64 `json:"volume"`
	OrderType           int     `json:"order_type"`
	Direction           string  `json:"direction"`
	Offset              string  `json:"offset"`
	LeverRate           int     `json:"lever_rate"`
	OrderIdStr          string  `json:"order_id_str"`
	RelationOrderId     string  `json:"relation_order_id"`
	OrderPriceType      int64   `json:"order_price_type"`
	Status              int     `json:"status"`
	OrderSource         string  `json:"order_source"`
	TriggerPrice        float64 `json:"trigger_price"`
	TriggeredPrice      float64 `json:"triggered_price,omitempty"`
	OrderPrice          float64 `json:"order_price"`
	CreatedAt           int64   `json:"created_at"`
	TriggeredAt         int64   `json:"triggered_at"`
	OrderInsertAt       int64   `json:"order_insert_at"`
	CanceledAt          int64   `json:"canceled_at,omitempty"`
	FailCode            int     `json:"fail_code,omitempty"`
	FailReason          string  `json:"fail_reason,omitempty"`
	UpdateTime          int64   `json:"update_time"`
	MarginMode          string  `json:"margin_mode"`
	MarginAccount       string  `json:"margin_account"`
	TpslOrderType       string  `json:"tpsl_order_type,omitempty"`
	SourceOrderIßd      string  `json:"source_order_id,omitempty"`
	RelationTpslOrderId string  `json:"relation_tpsl_order_id,omitempty"`
}

type GetRelationTpslTriggerOrderResponse struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"err_code,omitempty"`
	ErrorMessage string `json:"err_msg,omitempty"`
	Data         struct {
		Symbol         string  `json:"symbol"`
		ContractCode   string  `json:"contract_code"`
		Volume         float64 `json:"volume"`
		Price          float64 `json:"price"`
		OrderPriceType string  `json:"order_price_type"`
		Direction      string  `json:"direction"`
		Offset         string  `json:"offset"`
		LeverRate      int     `json:"lever_rate"`
		OrderId        int64   `json:"order_id"`
		OrderIdStr     string  `json:"order_id_str"`
		ClientOrderId  int64   `json:"client_order_id,omitempty"`
		CreatedAt      int64   `json:"created_at"`
		TradeVolume    float64 `json:"trade_volume"`
		TradeTurnover  float64 `json:"trade_turnover"`
		Fee            float64 `json:"fee"`
		TradeAvgPrice  float64 `json:"trade_avg_price"`
		//MarginAsset string `json:"margin_asset"`
		MarginFrozen float64 `json:"margin_frozen"`
		Profit       float64 `json:"profit"`
		Status       int     `json:"status"`
		OrderType    int     `json:"order_type"`
		OrderSource  string  `json:"order_source"`
		FeeAsset     string  `json:"fee_asset"`
		//LiquidationType string `json:"liquidation_type"`
		CanceledAt    int64  `json:"canceled_at"`
		MarginMode    string `json:"margin_mode"`
		MarginAccount string `json:"margin_account"`
		TpslOrderInfo []struct {
			Volume              float64 `json:"volume"`
			TpslOrderType       string  `json:"tpsl_order_type,omitempty"`
			Direction           string  `json:"direction"`
			OrderId             int64   `json:"order_id"`
			OrderIdStr          string  `json:"order_id_str"`
			TriggerType         string  `json:"trigger_type"`
			TriggerPrice        float64 `json:"trigger_price"`
			CreatedAt           int64   `json:"created_at"`
			OrderPriceType      string  `json:"order_price_type"`
			Status              int     `json:"status"`
			RelationTpslOrderId string  `json:"relation_tpsl_order_id,omitempty"`
			CanceledAt          int64   `json:"canceled_at,omitempty"`
			FailCode            int     `json:"fail_code,omitempty"`
			FailReason          string  `json:"fail_reason,omitempty"`
			TriggeredPrice      float64 `json:"triggered_price,omitempty"`
			RelationOrderId     string  `json:"relation_order_id"`
		} `json:"tpsl_order_info"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type TpslTriggerOrderResponse struct {
	Status       string `json:"status"`
	ErrorCode    int    `json:"err_code,omitempty"`
	ErrorMessage string `json:"err_msg,omitempty"`
	Data         struct {
		TpOrder struct {
			OrderId    int64  `json:"order_id"`
			OrderIdStr string `json:"order_id_str"`
		} `json:"tp_order"`
		SlOrder struct {
			OrderId    int64  `json:"order_id"`
			OrderIdStr string `json:"order_id_str"`
		} `json:"sl_order"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

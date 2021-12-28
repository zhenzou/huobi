package futures

type SubAccountsResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Uid string `json:"uid"`

	EventSender string `json:"event"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		MarginAsset string `json:"margin_asset"`

		MarginBalance float64 `json:"margin_balance"`

		MarginStatic float64 `json:"margin_static"`

		MarginPosition float64 `json:"margin_position"`

		MarginFrozen float64 `json:"margin_frozen"`

		MarginAvailable float64 `json:"margin_available"`

		ProfitReal float64 `json:"profit_real"`

		ProfitUnreal float64 `json:"profit_unreal"`

		WithdrawAvailable float64 `json:"withdraw_available"`

		RiskRate float64 `json:"risk_rate,omitempty"`

		LiquidationPrice float64 `json:"liquidation_price,omitempty"`

		LeverRate float64 `json:"lever_rate"`

		AdjustFactor float64 `json:"adjust_factor"`
	} `json:"data"`
}

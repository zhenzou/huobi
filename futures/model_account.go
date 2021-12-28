package futures

type AccountTransferResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		OrderId string `json:"order_id"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetAccountInfoResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

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

		WithdrawAvailable float64 `json:"withdraw_available,omitempty"`

		RiskRate float64 `json:"risk_rate,omitempty"`

		LiquidationPrice float64 `json:"liquidation_price,omitempty"`

		LeverRate float64 `json:"lever_rate"`

		AdjustFactor float64 `json:"adjust_factor"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		ContractDetail []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginPosition float64 `json:"margin_position"`

			MarginFrozen float64 `json:"margin_frozen"`

			MarginAvailable float64 `json:"margin_available"`

			ProfitUnreal float64 `json:"profit_unreal"`

			LiquidationPrice float64 `json:"liquidation_price"`

			LeverRate float64 `json:"lever_rate"`

			AdjustFactor float64 `json:"adjust_factor"`
		} `json:"contract_detail,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetAccountPositionResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

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

		MarginAsset string `json:"margin_asset"`

		PositionMargin float64 `json:"position_margin"`

		LeverRate int `json:"lever_rate"`

		Direction string `json:"direction"`

		LastPrice float64 `json:"last_price"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetAccountTransHisResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		FinancialRecord []struct {
			Id int64 `json:"id"`

			Ts int64 `json:"ts"`

			Asset string `json:"asset"`

			ContractCode string `json:"contract_code"`

			MarginAccount string `json:"margin_account"`

			Amount float64 `json:"amount"`

			// region only for account
			FaceMarginAccount string `json:"face_margin_account,omitempty"`

			FcType int `json:"type,omitempty"`
			// endregion

			// region only for sub account
			FromMarginAccount string `json:"from_margin_account,omitempty"`

			ToMarginAccount string `json:"to_margin_account,omitempty"`

			SubUid string `json:"sub_uid,omitempty"`

			SubAccountName string `json:"sub_account_name,omitempty"`

			TransferType int `json:"transfer_type,omitempty"`
			// endregion
		} `json:"financial_record"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetApiTradingStatusResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Ts int64 `json:"ts"`

	Data struct {
		IsDisable int `json:"is_disable"`

		OrderPriceTypes string `json:"order_price_types"`

		DisableReason string `json:"disable_reason"`

		DisableInterval int64 `json:"disable_interval"`

		RecoveryTime int64 `json:"recovery_time"`

		COR struct {
			OrdersThreshold int64 `json:"orders_threshold"`

			Orders int64 `json:"orders"`

			InvalidCancelOrders int64 `json:"invalid_cancel_orders"`

			CancelRatioThreshold float64 `json:"cancel_ratio_threshold"`

			CancelRatio float64 `json:"cancel_ratio"`

			IsTrigger int `json:"is_trigger"`

			IsActive int `json:"is_active"`
		}

		TDN struct {
			DisablesThreshold int64 `json:"disables_threshold"`

			Disables int64 `json:"disables"`

			IsTrigger int `json:"is_trigger"`

			IsActive int `json:"is_active"`
		}
	} `json:"data,omitempty"`
}

type GetAssetsPositionResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol,omitempty"`

		ContractCode string `json:"contract_code,omitempty"`

		MarginAsset string `json:"margin_asset"`

		MarginBalance float64 `json:"margin_balance"`

		MarginStatic float64 `json:"margin_static"`

		MarginPosition float64 `json:"margin_position"`

		MarginFrozen float64 `json:"margin_frozen"`

		MarginAvailable float64 `json:"margin_available,omitempty"`

		ProfitReal float64 `json:"profit_real"`

		ProfitUnreal float64 `json:"profit_unreal"`

		WithdrawAvailable float64 `json:"withdraw_available,omitempty"`

		RiskRate float64 `json:"risk_rate,omitempty"`

		LiquidationPrice float64 `json:"liquidation_price,omitempty"`

		LeverRate float64 `json:"lever_rate,omitempty"`

		AdjustFactor float64 `json:"adjust_factor,omitempty"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		ContractDetail []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginPosition float64 `json:"margin_position"`

			MarginFrozen float64 `json:"margin_frozen"`

			MarginAvailable float64 `json:"margin_available"`

			ProfitUnreal float64 `json:"profit_unreal"`

			LiquidationPrice float64 `json:"liquidation_price"`

			LeverRate float64 `json:"lever_rate"`

			AdjustFactor float64 `json:"adjust_factor"`
		} `json:"contract_detail,omitempty"`

		Positions []struct {
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

			MarginAsset string `json:"margin_asset"`

			PositionMargin float64 `json:"position_margin"`

			LeverRate int `json:"lever_rate"`

			Direction string `json:"direction"`

			LastPrice float64 `json:"last_price"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`
		} `json:"positions,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetAssetsPositionResponseSingle struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol,omitempty"`

		ContractCode string `json:"contract_code,omitempty"`

		MarginAsset string `json:"margin_asset"`

		MarginBalance float64 `json:"margin_balance"`

		MarginStatic float64 `json:"margin_static"`

		MarginPosition float64 `json:"margin_position"`

		MarginFrozen float64 `json:"margin_frozen"`

		MarginAvailable float64 `json:"margin_available,omitempty"`

		ProfitReal float64 `json:"profit_real"`

		ProfitUnreal float64 `json:"profit_unreal"`

		WithdrawAvailable float64 `json:"withdraw_available,omitempty"`

		RiskRate float64 `json:"risk_rate,omitempty"`

		LiquidationPrice float64 `json:"liquidation_price,omitempty"`

		LeverRate float64 `json:"lever_rate,omitempty"`

		AdjustFactor float64 `json:"adjust_factor,omitempty"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		ContractDetail []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginPosition float64 `json:"margin_position"`

			MarginFrozen float64 `json:"margin_frozen"`

			MarginAvailable float64 `json:"margin_available"`

			ProfitUnreal float64 `json:"profit_unreal"`

			LiquidationPrice float64 `json:"liquidation_price"`

			LeverRate float64 `json:"lever_rate"`

			AdjustFactor float64 `json:"adjust_factor"`
		} `json:"contract_detail,omitempty"`

		Positions []struct {
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

			MarginAsset string `json:"margin_asset"`

			PositionMargin float64 `json:"position_margin"`

			LeverRate int `json:"lever_rate"`

			Direction string `json:"direction"`

			LastPrice float64 `json:"last_price"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`
		} `json:"positions,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetFeeResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		OpenMakerFee string `json:"open_maker_fee"`

		OpenTakerFee string `json:"open_taker_fee"`

		CloseMakerFee string `json:"close_maker_fee"`

		CloseTakerFee string `json:"close_taker_fee"`

		FeeAsset string `json:"fee_asset"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetOrderLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		OrderPriceType string `json:"order_price_type"`

		List []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			OpenLimit float64 `json:"open_limit"`

			CloseLimit float64 `json:"close_limit"`
		} `json:"list"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetPositionLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		BuyLimit float64 `json:"buy_limit"`

		SellLimit float64 `json:"sell_limit"`

		MarginMode string `json:"margin_mode"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetSubAccountListResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		SubUid int64 `json:"sub_uid"`

		list []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginAsset string `json:"margin_asset"`

			MarginBalance float64 `json:"margin_balance"`

			LiquidationPrice float64 `json:"liquidation_price,omitempty"`

			RiskRate float64 `json:"risk_rate,omitempty"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`
		}
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetTransferLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		OpenLimit string `json:"open_limit"`

		TransferInMaxEach float64 `json:"transfer_in_max_each"`

		TransferInMinEach float64 `json:"transfer_in_min_each"`

		TransferOutMaxEach float64 `json:"transfer_out_max_each"`

		TransferOutMinEach float64 `json:"transfer_out_min_each"`

		TransferInMaxDaily float64 `json:"transfer_in_max_daily"`

		TransferOutMaxDaily float64 `json:"transfer_out_max_daily"`

		NetTransferInMaxDaily float64 `json:"net_transfer_in_max_daily"`

		NetTransferOutMaxDaily float64 `json:"net_transfer_out_max_daily"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetValidLeverRateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		ContractCode string `json:"contract_code"`

		MarginMode string `json:"margin_mode"`

		AvailableLeverRate string `json:"available_level_rate"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

package futures

type GetAdjustFactorFundResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		MarginMode string `json:"margin_mode"`

		List []struct {
			LeverRate float64 `json:"lever_rate"`

			Ladder []struct {
				MinSize float64 `json:"min_size"`

				MaxSize float64 `json:"max_size,omitempty"`

				Ladder int `json:"ladder"`

				AdjustFactor float64 `json:"adjust_factor"`
			} `json:"ladders"`
		} `json:"list"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetApiStateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		Open int `json:"open"`

		Close int `json:"close"`

		Cancel int `json:"cancel"`

		TransferIn int `json:"transfer_in"`

		TransferOut int `json:"transfer_out"`

		MasterTransferSub int `json:"master_transfer_sub"`

		SubTransferMaster int `json:"sub_transfer_master"`

		MasterTransferSubInnerIn int `json:"master_transfer_sub_inner_in"`

		MasterTransferSubInnerOut int `json:"master_transfer_sub_inner_out"`

		SubTransferMasterInnerIn int `json:"sub_transfer_master_inner_in"`

		SubTransferMasterInnerOut int `json:"sub_transfer_master_inner_out"`

		TransferInnerIn int `json:"transfer_inner_in"`

		TransferInnerOut int `json:"transfer_inner_out"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetBasisResponse struct {
	Ch string `json:"ch"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Id int64 `json:"id"`

		ContractPrice string `json:"contract_price"`

		IndexPrice string `json:"index_price"`

		Basis string `json:"basis"`

		BasisRate string `json:"basis_rate"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetContractInfoResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		ContractSize float64 `json:"contract_size"`

		PriceTick float64 `json:"price_tick"`

		SettlementDate string `json:"settlement_date"`

		CreateDate string `json:"create_date"`

		DeliveryTime string `json:"delivery_time"`

		ContractStatus int `json:"contract_status"`

		SupportMarginMode string `json:"support_margin_mode"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetDepthResponse struct {
	Ch string `json:"ch"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Tick struct {
		Asks [][]float64 `json:"asks"`

		Bids [][]float64 `json:"bids"`

		Ch string `json:"ch"`

		Id int64 `json:"id"`

		Mrid int64 `json:"mrid"`

		Ts int64 `json:"ts"`

		Version int64 `json:"version"`
	} `json:"tick,omitempty"`

	Ts int64 `json:"ts"`
}

type GetEliteRatioResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		ShortLongRatio []struct {
			BuyRatio float64 `json:"buy_ratio"`

			SellRatio float64 `json:"sell_ratio"`

			LockedRatio float64 `json:"locked_ratio,omitempty"`

			Ts int64 `json:"ts"`
		} `json:"list"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetEstimatedSettlementPriceResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		ContractCode string `json:"contract_code"`

		EstimatedSettlementPrice string `json:"estimated_settlement_price"`

		SettlementType string `json:"settlement_type"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetFundingRateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		FeeAsset string `json:"fee_asset"`

		FundingTime string `json:"funding_time"`

		FundingRate string `json:"funding_rate"`

		EstimatedRate string `json:"estimated_rate"`

		NextFundingTime string `json:"next_funding_time"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetHisFundingRateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Data []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			FeeAsset string `json:"fee_asset"`

			FundingTime string `json:"funding_time"`

			FundingRate string `json:"funding_rate"`

			RealizedRate string `json:"realized_rate"`

			AvgPremiumIndex string `json:"avg_premium_index"`
		} `json:"data"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetHisOpenInterestResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Tick []struct {
			Volume float64 `json:"volume"`

			AmountType int `json:"amount_type"`

			Value float64 `json:"value"`

			Ts int64 `json:"ts"`
		} `json:"tick"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetHisTradeResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	HisTrade []struct {
		Id int64 `json:"id"`

		Data []struct {
			Amount float64 `json:"amount"`

			Direction string `json:"direction"`

			Id int64 `json:"id"`

			Price float64 `json:"price"`

			Ts int64 `json:"ts"`
		} `json:"data"`

		Ts int64 `json:"ts"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetIndexResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		ContractCode string `json:"contract_code"`

		IndexPrice float64 `json:"index_price"`

		IndexTs int64 `json:"index_ts"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetInsuranceFundResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Tick []struct {
			InsuranceFund float64 `json:"insurance_fund"`

			Ts int64 `json:"ts"`
		} `json:"tick"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetKLineResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

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
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetLiquidationOrdersResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Order []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			CreatedAt int64 `json:"created_at"`

			Direction string `json:"direction"`

			Offset string `json:"offset"`

			Price float64 `json:"price"`

			Volume float64 `json:"volume"`
		} `json:"orders"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetMergedResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Tick struct {
		Id int64 `json:"id"`

		Amount string `json:"status"`

		Ask []float64 `json:"ask"`

		Bid []float64 `json:"bid"`

		Open string `json:"open"`

		Close string `json:"close"`

		Count float64 `json:"count"`

		High string `json:"high"`

		Low string `json:"low"`

		Vol string `json:"vol"`

		TradeTurnover string `json:"trade_turnover"`

		Ts int64 `json:"ts"`
	} `json:"tick,omitempty"`

	Ts int64 `json:"ts"`
}

type GetOpenInterestResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Amount float64 `json:"amount"`

		Volume float64 `json:"volume"`

		Value float64 `json:"value"`

		TradeAmount float64 `json:"trade_amount"`

		TradeVolume float64 `json:"trade_volume"`

		TradeTurnover float64 `json:"trade_turnover"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetPriceLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		HighLimit float64 `json:"high_limit"`

		LowLimit float64 `json:"low_limit"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetRiskInfoResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		EstimatedClawback float64 `json:"estimated_clawback"`

		InsuranceFund float64 `json:"insurance_fund"`

		ContractCode string `json:"contract_code"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetStrKLineResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Id int64 `json:"id"`

		Vol string `json:"vol"`

		Count string `json:"count"`

		Open string `json:"open"`

		Close string `json:"close"`

		Low string `json:"low"`

		High string `json:"high"`

		Amount string `json:"amount"`

		TradeTurnover string `json:"trade_turnover"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetTradeResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Tick struct {
		Id int64 `json:"id"`

		Data []struct {
			Amount string `json:"amount"`

			Direction string `json:"direction"`

			Id int64 `json:"id"`

			Price string `json:"price"`

			Ts int64 `json:"ts"`
		} `json:"data"`

		Ts int64 `json:"ts"`
	} `json:"tick,omitempty"`

	Ts int64 `json:"ts"`
}

type GetTradeStateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		Open int `json:"open"`

		Close int `json:"close"`

		Cancel int `json:"cancel"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetTransferStateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		TransferIn int `json:"transfer_in"`

		TransferOut int `json:"transfer_out"`

		MasterTransferSub int `json:"master_transfer_sub"`

		SubTransferMaster int `json:"sub_transfer_master"`

		MasterTransferSubInnerIn int `json:"master_transfer_sub_inner_in"`

		MasterTransferSubInnerOut int `json:"master_transfer_sub_inner_out"`

		SubTransferMasterInnerIn int `json:"sub_transfer_master_inner_in"`

		SubTransferMasterInnerOut int `json:"sub_transfer_master_inner_out"`

		TransferInnerIn int `json:"transfer_inner_in"`

		TransferInnerOut int `json:"transfer_inner_out"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

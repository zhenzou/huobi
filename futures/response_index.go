package futures

type ReqBasisResponse struct {
	Rep string `json:"rep"`

	Status string `json:"status"`

	Id string `json:"id"`

	Wsid int64 `json:"wsid"`

	Ts int64 `json:"ts"`

	Data []struct {
		Id int64 `json:"id"`

		ContractPrice string `json:"contract_price"`

		IndexPrice string `json:"index_price"`

		Basis string `json:"basis"`

		BasisRate string `json:"basis_rate"`
	} `json:"data"`
}

type ReqIndexKLineResponse struct {
	Rep string `json:"rep"`

	Status string `json:"status"`

	Id string `json:"id"`

	Wsid int64 `json:"wsid"`

	Ts int64 `json:"ts"`

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
	} `json:"data"`
}

type SubBasisResponse struct {
	Ch string `json:"rep"`

	Ts int64 `json:"ts"`

	Tick struct {
		Id int64 `json:"id"`

		ContractPrice string `json:"contract_price"`

		IndexPrice string `json:"index_price"`

		Basis string `json:"basis"`

		BasisRate string `json:"basis_rate"`
	} `json:"tick"`
}

type SubIndexKLineResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Id int64 `json:"id"`

		Vol string `json:"vol"`

		Count string `json:"count"`

		Open string `json:"open"`

		Close string `json:"close"`

		Low string `json:"low"`

		High string `json:"high"`

		Amount string `json:"amount"`

		TradeTurnover string `json:"trade_turnover"`
	} `json:"tick"`
}

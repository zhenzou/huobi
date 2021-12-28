package spot

import (
	"github.com/shopspring/decimal"
)

type GetCurrenciesResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

type GetMarketStatusResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    MarketStatus `json:"data"`
}

type MarketStatus struct {
	MarketStatus    int    `json:"marketStatus"`
	HaltStartTime   int64  `json:"haltStartTime"`
	HaltEndTime     int64  `json:"haltEndTime"`
	HaltReason      int    `json:"haltReason"`
	AffectedSymbols string `json:"affectedSymbols"`
}

type GetSymbolsResponse struct {
	Status string   `json:"status"`
	Data   []Symbol `json:"data"`
}

type Symbol struct {
	BaseCurrency           string          `json:"base-currency"`
	QuoteCurrency          string          `json:"quote-currency"`
	PricePrecision         int             `json:"price-precision"`
	AmountPrecision        int             `json:"amount-precision"`
	SymbolPartition        string          `json:"symbol-partition"`
	Symbol                 string          `json:"symbol"`
	State                  string          `json:"state"`
	ValuePrecision         int             `json:"value-precision"`
	LimitOrderMinOrderAmt  decimal.Decimal `json:"limit-order-min-order-amt"`
	LimitOrderMaxOrderAmt  decimal.Decimal `json:"limit-order-max-order-amt"`
	SellMarketMinOrderAmt  decimal.Decimal `json:"sell-market-min-order-amt"`
	SellMarketMaxOrderAmt  decimal.Decimal `json:"sell-market-max-order-amt"`
	BuyMarketMaxOrderValue decimal.Decimal `json:"buy-market-max-order-value"`
	MinOrderValue          decimal.Decimal `json:"min-order-value"`
	MaxOrderValue          decimal.Decimal `json:"max-order-value"`
	LeverageRatio          decimal.Decimal `json:"leverage-ratio"`
}

type GetTimestampResponse struct {
	Status string `json:"status"`
	Data   int    `json:"data"`
}

type GetV2ReferenceCurrencies struct {
	Currency       string
	AuthorizedUser string
}

type GetV2ReferenceCurrenciesResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []CurrencyChain `json:"data"`
}

type CurrencyChain struct {
	Currency   string   `json:"currency"`
	InstStatus string   `json:"instStatus"`
	Chains     []Chains `json:"chains"`
}

type Chains struct {
	Chain                   string `json:"chain"`
	BaseChain               string `json:"baseChain"`
	BaseChainProtocol       string `json:"baseChainProtocol"`
	IsDynamic               bool   `json:"isDynamic"`
	NumOfConfirmations      int    `json:"numOfConfirmations"`
	NumOfFastConfirmations  int    `json:"numOfFastConfirmations"`
	MinDepositAmt           string `json:"minDepositAmt"`
	DepositStatus           string `json:"depositStatus"`
	MinWithdrawAmt          string `json:"minWithdrawAmt"`
	MaxWithdrawAmt          string `json:"maxWithdrawAmt"`
	WithdrawQuotaPerDay     string `json:"withdrawQuotaPerDay"`
	WithdrawQuotaPerYear    string `json:"withdrawQuotaPerYear"`
	WithdrawQuotaTotal      string `json:"withdrawQuotaTotal"`
	WithdrawPrecision       int    `json:"withdrawPrecision"`
	WithdrawFeeType         string `json:"withdrawFeeType"`
	TransactFeeWithdraw     string `json:"transactFeeWithdraw"`
	MinTransactFeeWithdraw  string `json:"minTransactFeeWithdraw"`
	MaxTransactFeeWithdraw  string `json:"maxTransactFeeWithdraw"`
	TransactFeeRateWithdraw string `json:"transactFeeRateWithdraw"`
	WithdrawStatus          string `json:"withdrawStatus"`
}

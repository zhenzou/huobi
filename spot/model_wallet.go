package spot

import (
	"github.com/shopspring/decimal"
)

type CancelWithdrawResponse struct {
	Status string `json:"status"`
	Data   int64  `json:"data"`
}

type CreateWithdrawRequest struct {
	Address  string `json:"address"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Fee      string `json:"fee"`
	Chain    string `json:"chain,omitempty"`
	AddrTag  string `json:"addr-tag,omitempty"`
}

type CreateWithdrawResponse struct {
	Status string `json:"status"`
	Data   int64  `json:"data"`
}

type GetDepositAddressResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    []DepositAddress `json:"data"`
}
type DepositAddress struct {
	Currency   string `json:"currency"`
	Address    string `json:"address"`
	AddressTag string `json:"addressTag"`
	Chain      string `json:"chain"`
}

type GetWithdrawAddressResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Currency   string `json:"currency"`
		Chain      string `json:"chain"`
		Address    string `json:"address"`
		AddressTag string `json:"addressTag"`
		Note       string `json:"note"`
	}
	NextId int64 `json:"nextId"`
}

type GetWithdrawQuotaResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    *WithdrawQuota `json:"data"`
}
type WithdrawQuota struct {
	Currency string `json:"currency"`
	Chains   []struct {
		Chain                      string `json:"chain"`
		MaxWithdrawAmt             string `json:"maxWithdrawAmt"`
		WithdrawQuotaPerDay        string `json:"withdrawQuotaPerDay"`
		RemainWithdrawQuotaPerDay  string `json:"remainWithdrawQuotaPerDay"`
		WithdrawQuotaPerYear       string `json:"withdrawQuotaPerYear"`
		RemainWithdrawQuotaPerYear string `json:"remainWithdrawQuotaPerYear"`
		WithdrawQuotaTotal         string `json:"withdrawQuotaTotal"`
		RemainWithdrawQuotaTotal   string `json:"remainWithdrawQuotaTotal"`
	} `json:"chains"`
}

type QueryDepositWithdrawOptionalRequest struct {
	Currency string
	From     string
	Size     string
	Direct   string
}

type QueryDepositWithdrawResponse struct {
	Status string            `json:"status"`
	Data   []DepositWithdraw `json:"data"`
}
type DepositWithdraw struct {
	Id         int64           `json:"id"`
	Type       string          `json:"type"`
	Currency   string          `json:"currency"`
	TxHash     string          `json:"tx-hash"`
	Chain      string          `json:"chain"`
	Amount     decimal.Decimal `json:"amount"`
	Address    string          `json:"address"`
	AddressTag string          `json:"address-tag"`
	Fee        decimal.Decimal `json:"fee"`
	State      string          `json:"state"`
	CreatedAt  int64           `json:"created-at"`
	UpdatedAt  int64           `json:"updated-at"`
}

package spot

import (
	"github.com/shopspring/decimal"
)

type CreateSubUserRequest struct {
	UserList []Users `json:"userList"`
}

type Users struct {
	UserName string `json:"userName"`
	Note     string `json:"note"`
}

type CreateSubUserResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []UserData `json:"data"`
}

type UserData struct {
	UserName     string `json:"userName"`
	Note         string `json:"note"`
	Uid          int64  `json:"uid"`
	ErrorCode    string `json:"errCode"`
	ErrorMessage string `json:"errMessage"`
}

type QuerySubUserDepositHistoryOptionalRequest struct {
	Currency  string `json:"currency"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Sort      string `json:"sort"`
	Limit     string `json:"limit"`
	FromId    int64  `json:"fromId"`
}

type QuerySubUserDepositHistoryResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    []DepositHistory `json:"data"`
	NextId  int64            `json:"nextId"`
}

type DepositHistory struct {
	Id              int64           `json:"id"`
	Currency        string          `json:"currency"`
	TransactionHash string          `json:"txHash"`
	Chain           string          `json:"chain"`
	Amount          decimal.Decimal `json:"amount"`
	Address         string          `json:"address"`
	AddressTag      string          `json:"addressTag"`
	State           string          `json:"state"`
	CreateTime      int64           `json:"createTime"`
	UpdateTime      int64           `json:"updateTime"`
}

type SetSubUserTradableMarketResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    []TradableMarket `json:"data"`
}

type TradableMarket struct {
	SubUid      string `json:"subUid"`
	AccountType string `json:"accountType"`
	Activation  string `json:"activation"`
	ErrCode     int    `json:"errCode"`
	ErrMessage  string `json:"errMessage"`
}

type SetSubUserTradableMarketRequest struct {
	SubUids     string `json:"subUids"`
	AccountType string `json:"accountType"`
	Activation  string `json:"activation"`
}

type SetSubUserTransferabilityRequest struct {
	SubUids       string `json:"subUids"`
	AccountType   string `json:"accountType"`
	Transferrable bool   `json:"transferrable"`
}

type SetSubUserTransferabilityResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []Transferability `json:"data"`
}

type Transferability struct {
	SubUid        int64  `json:"subUid"`
	AccountType   string `json:"accountType"`
	Transferrable bool   `json:"transferrable"`
	ErrCode       int    `json:"errCode"`
	ErrMessage    string `json:"errMessage"`
}

type SubUserManagementRequest struct {
	SubUid int64  `json:"subUid"`
	Action string `json:"action"`
}

type SubUserManagementResponse struct {
	Code int `json:"code"`
	Data *SubUserManagement
}
type SubUserManagement struct {
	SubUid    int64  `json:"subUid"`
	UserState string `json:"userState"`
}

type SubUserTransferRequest struct {
	SubUid   int64           `json:"sub-uid"`
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
	Type     string          `json:"type"`
}

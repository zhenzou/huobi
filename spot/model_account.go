package spot

import (
	"github.com/shopspring/decimal"
)

type FuturesTransferRequest struct {
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
	Type     string          `json:"type"`
}

type FuturesTransferResponse struct {
	Status string `json:"status"`
	Data   int64  `json:"data"`
}

type GetAccountAssetValuationResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		Balance   string `json:"balance"`
		Timestamp int64  `json:"timestamp"`
	}
}
type GetAccountBalanceResponse struct {
	Status string          `json:"status"`
	Data   *AccountBalance `json:"data"`
}

type GetAccountHistoryOptionalRequest struct {
	Currency      string
	TransactTypes string
	StartTime     int64
	EndTime       int64
	Sort          string
	Size          int
}

type GetAccountHistoryResponse struct {
	Status string           `json:"status"`
	Data   []AccountHistory `json:"data"`
	NextId int64            `json:"next-id"`
}

type AccountHistory struct {
	AccountId    int64  `json:"account-id"`
	Currency     string `json:"currency"`
	TransactAmt  string `json:"transact-amt"`
	TransactType string `json:"transact-type"`
	RecordId     int64  `json:"record-id"`
	AvailBalance string `json:"avail-balance"`
	AcctBalance  string `json:"acct-balance"`
	TransactTime int64  `json:"transact-time"`
}

type GetAccountInfoResponse struct {
	Status string        `json:"status"`
	Data   []AccountInfo `json:"data"`
}
type AccountInfo struct {
	Id      int64  `json:"id"`
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	State   string `json:"state"`
}

type GetAccountLedgerOptionalRequest struct {
	Currency      string `json:"currency"`
	TransactTypes string `json:"transactTypes"`
	StartTime     int64  `json:"startTime"`
	EndTime       int64  `json:"endTime"`
	Sort          string `json:"sort"`
	Limit         int    `json:"limit"`
	FromId        int64  `json:"fromId"`
}

type GetAccountLedgerResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []Ledger `json:"data"`
	NextId  int64    `json:"nextId"`
}

type Ledger struct {
	AccountId    int64           `json:"accountId"`
	Currency     string          `json:"currency"`
	TransactAmt  decimal.Decimal `json:"transactAmt"`
	TransactType string          `json:"transactType"`
	TransferType string          `json:"transferType"`
	TransactId   int64           `json:"transactId"`
	TransactTime int64           `json:"transactTime"`
	Transferer   int64           `json:"transferer"`
	Transferee   int64           `json:"transferee"`
}

type GetPointBalanceResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		AccountId      string `json:"accountId"`
		AccountStatus  string `json:"accountStatus"`
		AccountBalance string `json:"acctBalance"`
		GroupIds       []struct {
			GroupId      int64  `json:"groupId"`
			ExpiryDate   int64  `json:"expiryDate"`
			RemainAmount string `json:"remainAmount"`
		}
	}
}

type GetSubUserAccountResponse struct {
	Status string           `json:"status"`
	Data   []SubUserAccount `json:"data"`
}
type SubUserAccount struct {
	Id   int        `json:"id"`
	Type string     `json:"type"`
	List []*Balance `json:"list"`
}

type GetSubUserAggregateBalanceResponse struct {
	Status string    `json:"status"`
	Data   []Balance `json:"data"`
}
type GetUidResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    int64  `json:"data"`
}

type SubscribeAccountV2Response struct {
	WebSocketV2ResponseBase
	Data *struct {
		Currency    string `json:"currency"`
		AccountId   int    `json:"accountId"`
		Balance     string `json:"balance"`
		Available   string `json:"available"`
		ChangeType  string `json:"changeType"`
		AccountType string `json:"accountType"`
		ChangeTime  int64  `json:"changeTime"`
	}
}

type TransferAccountRequest struct {
	FromUser        int64  `json:"from-user"`
	FromAccountType string `json:"from-account-type"`
	FromAccount     int64  `json:"from-account"`
	ToUser          int64  `json:"to-user"`
	ToAccountType   string `json:"to-account-type"`
	ToAccount       int64  `json:"to-account"`
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
}

type TransferAccountResponse struct {
	Status string              `json:"status"`
	Data   TransferAccountData `json:"data"`
}

type TransferAccountData struct {
	TransactId   int64 `json:"transact-id"`
	TransactTime int64 `json:"transact-time"`
}
type TransferPointRequest struct {
	FromUid string `json:"fromUid"`
	ToUid   string `json:"toUid"`
	GroupId int64  `json:"groupId"`
	Amount  string `json:"amount"`
}
type TransferPointResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		TransactId   string `json:"transactId"`
		TransactTime int64  `json:"transactTime"`
	}
}

type RequestAccountV1Response struct {
	Timestamp int64            `json:"ts"`
	Op        string           `json:"op"`
	Topic     string           `json:"topic"`
	ErrorCode int              `json:"err-code"`
	ClientId  string           `json:"cid"`
	Data      []AccountBalance `json:"data"`
}

type AccountBalance struct {
	Id    int       `json:"id"`
	Type  string    `json:"type"`
	State string    `json:"state"`
	List  []Balance `json:"list"`
}

type Balance struct {
	Currency string `json:"currency"`
	Type     string `json:"type"`
	Balance  string `json:"balance"`
}

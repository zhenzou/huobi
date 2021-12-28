package spot

type CrossMarginAccountsBalanceResponse struct {
	Status string                      `json:"status"`
	Data   *CrossMarginAccountsBalance `json:"data"`
}
type CrossMarginAccountsBalance struct {
	Id             int64  `json:"id"`
	Type           string `json:"type"`
	State          string `json:"state"`
	AcctBalanceSum string `json:"acct-balance-sum"`
	DebtBalanceSum string `json:"debt-balance-sum"`
	RiskRate       string `json:"risk-rate"`
	List           []struct {
		Currency string `json:"currency"`
		Type     string `json:"type"`
		Balance  string `json:"balance"`
	}
}

type CrossMarginGeneralReplayLoanOptionalRequest struct {
	AccountId  string `json:"accountId"`
	Currency   string `json:"currency"`
	Amount     string `json:"amount"`
	TransactId string `json:"transactId"`
}

type CrossMarginGeneralReplayLoanRecordsOptionalRequest struct {
	RepayId   string `json:"repayId"`
	AccountId string `json:"accountId"`
	Currency  string `json:"currency"`
	StartDate int64  `json:"startDate"`
	EndDate   int64  `json:"endDate"`
	Sort      string `json:"sort"`
	Limit     int    `json:"limit"`
	FromId    int64  `json:"fromId"`
}

type CrossMarginGeneralReplyLoanRecordsResponse struct {
	Code int                                  `json:"code"`
	Data []CrossMarginGeneraReplaylLoanRecord `json:"data"`
}
type CrossMarginGeneraReplaylLoanRecord struct {
	RepayId      int64     `json:"repayId"`
	RepayTime    int64     `json:"repayTime"`
	AccountId    int64     `json:"accountId"`
	Currency     string    `json:"currency"`
	RepaidAmount string    `json:"repaidAmount"`
	TransactIds  *Transact `json:"transactIds"`
}
type Transact struct {
	TransactId      int64  `json:"transactId"`
	RepaidPrincipal string `json:"repaidPrincipal"`
	RepaidInterest  string `json:"repaidInterest"`
	PaidHt          string `json:"paidHt"`
	PaidPoint       string `json:"paidPoint"`
}

type CrossMarginGeneralReplyLoanResponse struct {
	Code int                            `json:"code"`
	Data []CrossMarginGeneraReplaylLoan `json:"data"`
}
type CrossMarginGeneraReplaylLoan struct {
	RepayId   int64 `json:"repayId"`
	RepayTime int64 `json:"repayTime"`
}

type CrossMarginLoanOrdersOptionalRequest struct {
	StartDate string
	EndDate   string
	Currency  string
	State     string
	From      string
	Direct    string
	Size      string
	SubUid    string
}

type CrossMarginLoanOrdersResponse struct {
	Status string                 `json:"status"`
	Data   []CrossMarginLoanOrder `json:"data"`
}
type CrossMarginLoanOrder struct {
	Id              int64  `json:"id"`
	UserId          int64  `json:"user-id"`
	AccountId       int64  `json:"account-id"`
	Currency        string `json:"currency"`
	LoanAmount      string `json:"loan-amount"`
	LoanBalance     string `json:"loan-balance"`
	InterestAmount  string `json:"interest-amount"`
	InterestBalance string `json:"interest-balance"`
	CreatedAt       int64  `json:"created-at"`
	AccruedAt       int64  `json:"accrued-at"`
	State           string `json:"state"`
	FilledPoints    string `json:"filled-points"`
	FilledHt        string `json:"filled-ht"`
}

type CrossMarginOrdersRequest struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type CrossMarginTransferRequest struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type GetCrossMarginLoanInfoResponse struct {
	Status string                `json:"status"`
	Data   []CrossMarginLoanInfo `json:"data"`
}
type CrossMarginLoanInfo struct {
	Currency     string `json:"currency"`
	InterestRate string `json:"interest-rate"`
	MinLoanAmt   string `json:"min-loan-amt"`
	MaxLoanAmt   string `json:"max-loan-amt"`
	LoanableAmt  string `json:"loanable-amt"`
	ActualRate   string `json:"actual-rate"`
}

type GetIsolatedMarginLoanInfoResponse struct {
	Status string                   `json:"status"`
	Data   []IsolatedMarginLoanInfo `json:"data"`
}
type IsolatedMarginLoanInfo struct {
	Symbol     string `json:"symbol"`
	Currencies []struct {
		Currency     string `json:"currency"`
		InterestRate string `json:"interest-rate"`
		MinLoanAmt   string `json:"min-loan-amt"`
		MaxLoanAmt   string `json:"max-loan-amt"`
		LoanableAmt  string `json:"loanable-amt"`
		ActualRate   string `json:"actual-rate"`
	}
}

type GetMarginLoanInfoOptionalRequest struct {
	Symbols string
}

type IsolatedMarginAccountsBalanceResponse struct {
	Status string                          `json:"status"`
	Data   []IsolatedMarginAccountsBalance `json:"data"`
}
type IsolatedMarginAccountsBalance struct {
	Id       int64  `json:"id"`
	Type     string `json:"type"`
	State    string `json:"state"`
	Symbol   string `json:"symbol"`
	FlPrice  string `json:"fl-price"`
	FlType   string `json:"fl-type"`
	RiskRate string `json:"risk-rate"`
	List     []struct {
		Currency string `json:"currency"`
		Type     string `json:"type"`
		Balance  string `json:"balance"`
	}
}

type IsolatedMarginLoanOrdersOptionalRequest struct {
	StartDate string
	EndDate   string
	States    string
	From      string
	Direct    string
	Size      string
	SubUid    int
}

type IsolatedMarginLoanOrdersResponse struct {
	Status string                    `json:"status"`
	Data   []IsolatedMarginLoanOrder `json:"data"`
}
type IsolatedMarginLoanOrder struct {
	Id              int64  `json:"id"`
	UserId          int64  `json:"user-id"`
	AccountId       int64  `json:"account-id"`
	Symbol          string `json:"symbol"`
	Currency        string `json:"currency"`
	LoanAmount      string `json:"loan-amount"`
	LoanBalance     string `json:"loan-balance"`
	InterestRate    string `json:"interest-rate"`
	InterestAmount  string `json:"interest-amount"`
	InterestBalance string `json:"interest-balance"`
	CreatedAt       int64  `json:"created-at"`
	AccruedAt       int64  `json:"accrued-at"`
	State           string `json:"state"`
}

type IsolatedMarginOrdersRequest struct {
	Symbol   string `json:"symbol"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type IsolatedMarginTransferRequest struct {
	Symbol   string `json:"symbol"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type MarginAccountsBalanceOptionalRequest struct {
	Symbol string
	SubUid int
}

type MarginOrdersRepayRequest struct {
	Amount string `json:"amount"`
}

type MarginOrdersRepayResponse struct {
	Status string `json:"status"`
	Data   int    `json:"data"`
}

type MarginOrdersResponse struct {
	Status string `json:"status"`
	Data   int    `json:"data"`
}

type TransferResponse struct {
	Status string `json:"status"`
	Data   int    `json:"data"`
}

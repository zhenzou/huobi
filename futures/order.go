package futures

import (
	"context"
	"fmt"
	"strings"

	"github.com/zhenzou/huobi/client"
	"github.com/zhenzou/huobi/config"
)

func NewOrderClient(cfg config.Config) *OrderClient {
	return &OrderClient{
		wsClient: client.NewWebSocketClient(cfg, "/linear-swap-notification"),
		rsClient: client.NewRestfulClient(cfg),
	}
}

type OrderClient struct {
	wsClient *client.WebSocketClient
	rsClient *client.RestfulClient
}

func (cli *OrderClient) SubOrders(ctx context.Context, contractCode string) (<-chan SubOrdersResponse, error) {
	topic := fmt.Sprintf("orders.%s", contractCode)

	receiver := make(chan SubOrdersResponse, 1024)
	err := cli.wsClient.SubWithOp(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *OrderClient) UnsubOrders(ctx context.Context, contractCode string) error {
	topic := fmt.Sprintf("orders.%s", contractCode)
	return cli.wsClient.UnsubWithOP(ctx, topic)
}

func (cli *OrderClient) SubPositions(ctx context.Context, contractCode string) (<-chan SubPositionsResponse, error) {
	topic := fmt.Sprintf("positions.%s", contractCode)

	receiver := make(chan SubPositionsResponse, 1024)
	err := cli.wsClient.SubWithOp(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *OrderClient) UnsubPositions(ctx context.Context, contractCode string) error {

	topic := fmt.Sprintf("positions.%s", contractCode)

	return cli.wsClient.UnsubWithOP(ctx, topic)
}

func (cli *OrderClient) SubMatchOrders(ctx context.Context, contractCode string) (<-chan SubOrdersResponse, error) {
	topic := fmt.Sprintf("matchOrders.%s", strings.ToLower(contractCode))

	receiver := make(chan SubOrdersResponse, 1024)
	err := cli.wsClient.SubWithOp(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil

}

func (cli *OrderClient) UnsubMathOrders(ctx context.Context, contractCode string) error {
	contractCode = strings.ToLower(contractCode)
	topic := fmt.Sprintf("matchOrders.%s", contractCode)
	return cli.wsClient.UnsubWithOP(ctx, topic)
}

func (cli *OrderClient) SubLiquidationOrders(ctx context.Context, contractCode string) (<-chan SubLiquidationOrdersResponse, error) {
	topic := fmt.Sprintf("public.%s.liquidation_orders", contractCode)

	receiver := make(chan SubLiquidationOrdersResponse, 1024)
	err := cli.wsClient.SubWithOp(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *OrderClient) UnsubLiquidationOrders(ctx context.Context, contractCode string) error {

	topic := fmt.Sprintf("public.%s.liquidation_orders", contractCode)
	return cli.wsClient.UnsubWithOP(ctx, topic)
}

func (cli *OrderClient) SubFundingRate(ctx context.Context, contractCode string) (<-chan SubFundingRateResponse, error) {
	topic := fmt.Sprintf("public.%s.funding_rate", contractCode)

	receiver := make(chan SubFundingRateResponse, 1024)
	err := cli.wsClient.SubWithOp(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *OrderClient) UnsubFundingRate(ctx context.Context, contractCode string) error {

	topic := fmt.Sprintf("public.%s.funding_rate", contractCode)

	return cli.wsClient.UnsubWithOP(ctx, topic)
}

func (cli *OrderClient) SubContractInfo(ctx context.Context, contractCode string) (<-chan SubContractInfoResponse, error) {
	topic := fmt.Sprintf("public.%s.contract_info", contractCode)

	receiver := make(chan SubContractInfoResponse, 1024)
	err := cli.wsClient.SubWithOp(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *OrderClient) UnsubContractInfo(ctx context.Context, contractCode string) error {

	topic := fmt.Sprintf("public.%s.contract_info", contractCode)

	return cli.wsClient.UnsubWithOP(ctx, topic)
}

func (cli *OrderClient) SubTriggerOrder(ctx context.Context, contractCode string) (<-chan SubTriggerOrderResponse, error) {
	topic := fmt.Sprintf("trigger_order.%s", contractCode)

	receiver := make(chan SubTriggerOrderResponse, 1024)
	err := cli.wsClient.SubWithOp(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *OrderClient) UnsubTriggerOrder(ctx context.Context, contractCode string) error {
	topic := fmt.Sprintf("trigger_order.%s", contractCode)
	return cli.wsClient.UnsubWithOP(ctx, topic)
}

type CreateTriggerOrderRequest struct {
	ContractCode string `json:"contract_code"`
	// ??????????????? ge????????????(????????????????????????)???le??????(????????????????????????)
	TriggerType  string  `json:"trigger_type"`
	TriggerPrice float64 `json:"trigger_price"`
	Volume       int64   `json:"volume"`
	// buy:??? sell:???
	Direction string `json:"direction"`
	// open:??? close:???
	Offset string `json:"offset"`
	// ??????????????????????????????????????????????????????[????????????10???????????????????????????20?????????
	// ????????????????????????(>20???)???????????????????????????web??????????????????????????????????????????????????????????????????(>20???)]
	LeverRate int `json:"lever_rate,omitempty"`
	// ???????????????????????????????????????????????????
	OrderPrice float64 `json:"order_price,omitempty"`
	// ??????????????? ???????????????limit; ?????????limit ?????????5??????optimal_5?????????10??????optimal_10?????????20??????optimal_20
	OrderPriceType string `json:"order_price_type,omitempty"`
}

// ?????? order id
func (cli *OrderClient) CreateTriggerOrder(ctx context.Context, request CreateTriggerOrderRequest) (string, error) {
	var resp = struct {
		OrderID string `json:"order_id_str"`
	}{}
	err := cli.rsClient.Post(ctx,
		"/linear-swap-api/v1/swap_trigger_order", request, &resp,
	)
	if err != nil {
		return "", err
	}
	return resp.OrderID, nil
}

// ?????? order id
func (cli *OrderClient) CancelTriggerOrder(ctx context.Context, contractCode, orderID string) error {
	var resp = struct {
		Errors []struct {
			OrderID string `json:"order_id"`
			ErrCode int64  `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		} `json:"errors"`
		Successes string `json:"successes"`
	}{}

	var request = map[string]string{
		"contract_code": contractCode,
		"order_id":      orderID,
	}
	err := cli.rsClient.Post(ctx,
		"/linear-swap-api/v1/swap_trigger_cancel", request, &resp,
	)
	if err != nil {
		return err
	}
	if len(resp.Errors) > 0 {
		return fmt.Errorf("huobi api error. code %d msg %s", resp.Errors[0].ErrCode, resp.Errors[0].ErrMsg)
	}
	return nil
}

type BatchCreateOrderRequest []CreateOrderRequest

/**
????????????
?????????????????????(direction???buy???offset???open)

?????????????????????(direction???sell???offset???close)

?????????????????????(direction???sell???offset???open)

?????????????????????(direction???buy???offset???close)
*/
type CreateOrderRequest struct {
	ContractCode  string  `json:"contract_code"`
	ClientOrderID int64   `json:"client_order_id,omitempty"`
	Price         float64 `json:"price"`
	Volume        int64   `json:"volume"`
	// buy:??? sell:???
	Direction string `json:"direction"`
	// open:??? close:???
	Offset string `json:"offset"`
	// ??????????????????????????????????????????????????????[????????????10???????????????????????????20?????????
	// ????????????????????????(>20???)???????????????????????????web??????????????????????????????????????????????????????????????????(>20???)]
	LeverRate int64 `json:"lever_rate,omitempty"`
	// ??????????????? ???????????????limit; ?????????limit ?????????5??????optimal_5?????????10??????optimal_10?????????20??????optimal_20
	OrderPriceType string `json:"order_price_type,omitempty"`
	// ??????????????????
	TpTriggerPrice   float64 `json:"tp_trigger_price,omitempty"`
	TpOrderPrice     float64 `json:"tp_order_price,omitempty"`
	TpOrderPriceType string  `json:"tp_order_price_type,omitempty"`

	// ??????????????????
	SlTriggerPrice   float64 `json:"sl_trigger_price,omitempty"`
	SlOrderPrice     float64 `json:"sl_order_price,omitempty"`
	SlOrderPriceType string  `json:"sl_order_price_type,omitempty"`
}

// ?????? order id
func (cli *OrderClient) CreateOrder(ctx context.Context, request CreateOrderRequest) (string, error) {
	var resp = struct {
		OrderID string `json:"order_id_str"`
	}{}
	err := cli.rsClient.Post(ctx,
		"/linear-swap-api/v1/swap_order", request, &resp,
	)
	if err != nil {
		return "", err
	}
	return resp.OrderID, nil
}

// ?????? order id
func (cli *OrderClient) CancelOrder(ctx context.Context, contractCode, orderID string) error {
	var resp = struct {
		Errors []struct {
			OrderID string `json:"order_id"`
			ErrCode int64  `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		} `json:"errors"`
		Successes string `json:"successes"`
	}{}

	var request = map[string]string{
		"contract_code": contractCode,
		"order_id":      orderID,
	}
	err := cli.rsClient.Post(ctx,
		"/linear-swap-api/v1/swap_cancel", request, &resp,
	)
	if err != nil {
		return err
	}
	if len(resp.Errors) > 0 {
		return fmt.Errorf("huobi api error. code %d msg %s", resp.Errors[0].ErrCode, resp.Errors[0].ErrMsg)
	}
	return nil
}

type ListOrdersRequest struct {
	client.PageRequest
	ContractCode string `json:"contract_code"`
	TradeType    int64  `json:"trade_type"`
	Status       string `json:"status"`
	CreateDate   int64  `json:"create_date"`
	Type         int64  `json:"type,omitempty"` //  1:????????????,2:?????????????????????	?????? order ????????????
}

func (cli *OrderClient) ListOpenOrders(ctx context.Context, req ListOrdersRequest) ([]Order, client.PageResponse, error) {
	var resp = struct {
		client.PageResponse
		Orders []Order `json:"orders"`
	}{}

	err := cli.rsClient.Post(ctx, "/linear-swap-api/v1/swap_openorders", req, &resp)
	if err != nil {
		return nil, client.PageResponse{}, err
	}
	return resp.Orders, resp.PageResponse, nil
}

func (cli *OrderClient) ListHistoryOrders(ctx context.Context, req ListOrdersRequest) ([]Order, client.PageResponse, error) {
	if req.Status == "" {
		req.Status = "0"
	}
	if req.CreateDate == 0 {
		req.CreateDate = 1
	}
	resp := struct {
		client.PageResponse
		Orders []Order `json:"orders"`
	}{}

	err := cli.rsClient.Post(ctx, "/linear-swap-api/v1/swap_hisorders", req, &resp)
	if err != nil {
		return nil, client.PageResponse{}, err
	}
	return resp.Orders, resp.PageResponse, nil
}

func (cli *OrderClient) ListOpenTriggerOrders(ctx context.Context, req ListOrdersRequest) ([]TriggerOrder, client.PageResponse, error) {
	var resp = struct {
		client.PageResponse
		Orders []TriggerOrder `json:"orders"`
	}{}

	err := cli.rsClient.Post(ctx, "/linear-swap-api/v1/swap_trigger_openorders", req, &resp)
	if err != nil {
		return nil, client.PageResponse{}, err
	}
	return resp.Orders, resp.PageResponse, nil
}

func (cli *OrderClient) ListHistoryTriggerOrders(ctx context.Context, req ListOrdersRequest) ([]TriggerOrder, client.PageResponse, error) {
	if req.Status == "" {
		req.Status = "0"
	}
	if req.CreateDate == 0 {
		req.CreateDate = 1
	}
	resp := struct {
		client.PageResponse
		Orders []TriggerOrder `json:"orders"`
	}{}

	err := cli.rsClient.Post(ctx, "/linear-swap-api/v1/swap_trigger_hisorders", req, &resp)
	if err != nil {
		return nil, client.PageResponse{}, err
	}
	return resp.Orders, resp.PageResponse, nil
}

type TpslTriggerOrderRequest struct {
	ContractCode     string  `json:"contract_code"`
	Direction        string  `json:"direction"`
	Volume           int64   `json:"volume"`
	TpTriggerPrice   float64 `json:"tp_trigger_price,omitempty"`
	TpOrderPrice     float64 `json:"tp_order_price,omitempty"`
	TpOrderPriceType string  `json:"tp_order_price_type,omitempty"`
	SlTriggerPrice   float64 `json:"sl_trigger_price,omitempty"`
	SlOrderPrice     float64 `json:"sl_order_price,omitempty"`
	SlOrderPriceType string  `json:"sl_order_price_type,omitempty"`
}
type BatchTpslOrderRequest []TpslTriggerOrderRequest

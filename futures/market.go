package futures

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/zhenzou/huobi/client"
	"github.com/zhenzou/huobi/config"
)

func NewMarketClient(cfg config.Config) *MarketClient {
	return &MarketClient{
		wsClient: client.NewWebSocketClient(cfg, "/linear-swap-ws"),
		rsClient: client.NewRestfulClient(cfg),
	}
}

type MarketClient struct {
	wsClient *client.WebSocketClient
	rsClient *client.RestfulClient
}

// -------------------------------------------------------------
// kline start
func (cli *MarketClient) SubKLine(ctx context.Context, contractCode string, period string) (<-chan SubKLineResponse, error) {
	topic := fmt.Sprintf("market.%s.kline.%s", contractCode, period)

	receiver := make(chan SubKLineResponse, 1024)
	err := cli.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *MarketClient) UnsubKLine(ctx context.Context, contractCode string, period string) error {
	topic := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
	return cli.wsClient.Unsub(ctx, topic)
}

func (cli *MarketClient) ReqKLine(ctx context.Context, contractCode string, period string, from, to time.Time) (resp ReqKLineResponse, err error) {
	topic := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
	err = cli.wsClient.Req(ctx, topic, from, to, &resp)
	return
}

// kline end
//-------------------------------------------------------------

//-------------------------------------------------------------
// depth start
func (cli *MarketClient) SubDepth(ctx context.Context, contractCode string, fcType string) (<-chan SubDepthResponse, error) {
	topic := fmt.Sprintf("market.%s.depth.%s", contractCode, fcType)

	receiver := make(chan SubDepthResponse, 1024)
	err := cli.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *MarketClient) UnsubDepth(ctx context.Context, contractCode string, fcType string) error {
	topic := fmt.Sprintf("market.%s.depth.%s", contractCode, fcType)
	return cli.wsClient.Unsub(ctx, topic)
}

// depth end
//-------------------------------------------------------------

//-------------------------------------------------------------
// incrementa depth start

func (cli *MarketClient) SubIncrementalDepth(ctx context.Context, contractCode string, size string) (<-chan SubDepthResponse, error) {
	topic := fmt.Sprintf("market.%s.depth.size_%s.high_freq", contractCode, size)

	receiver := make(chan SubDepthResponse, 1024)
	err := cli.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *MarketClient) UnsubIncrementalDepth(ctx context.Context, contractCode string, size string) error {
	topic := fmt.Sprintf("market.%s.depth.size_%s.high_freq", contractCode, size)
	return cli.wsClient.Unsub(ctx, topic)
}

// incrementa depth end
//-------------------------------------------------------------

//-------------------------------------------------------------
// detail start
func (cli *MarketClient) SubDetail(ctx context.Context, contractCode string) (<-chan SubKLineResponse, error) {
	topic := fmt.Sprintf("market.%s.detail", contractCode)

	receiver := make(chan SubKLineResponse, 1024)
	err := cli.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *MarketClient) UnsubDetail(ctx context.Context, contractCode string) error {
	topic := fmt.Sprintf("market.%s.detail", contractCode)
	return cli.wsClient.Unsub(ctx, topic)
}

// detail end
//-------------------------------------------------------------

//-------------------------------------------------------------
// bbo start
func (cli *MarketClient) SubBBO(ctx context.Context, contractCode string) (<-chan SubBBOResponse, error) {
	topic := fmt.Sprintf("market.%s.bbo", contractCode)

	receiver := make(chan SubBBOResponse, 1024)
	err := cli.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *MarketClient) UnsubBBO(ctx context.Context, contractCode string) error {
	topic := fmt.Sprintf("market.%s.bbo", contractCode)
	return cli.wsClient.Unsub(ctx, topic)
}

// bbo end
//-------------------------------------------------------------

//-------------------------------------------------------------
// trade detail start
func (cli *MarketClient) SubTradeDetail(ctx context.Context, contractCode string) (<-chan SubTradeDetailResponse, error) {
	topic := fmt.Sprintf("market.%s.trade.detail", contractCode)

	receiver := make(chan SubTradeDetailResponse, 1024)
	err := cli.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (cli *MarketClient) UnsubTradeDetail(ctx context.Context, contractCode string) error {
	topic := fmt.Sprintf("market.%s.trade.detail", contractCode)
	return cli.wsClient.Unsub(ctx, topic)
}

func (cli *MarketClient) ReqTradeDetail(ctx context.Context, contractCode string) (resp ReqTradeDetailResponse, err error) {
	topic := fmt.Sprintf("market.%s.trade.detail", contractCode)
	err = cli.wsClient.Req(ctx, topic, time.Time{}, time.Time{}, &resp)
	return
}

func (cli *MarketClient) ListContractInfos(ctx context.Context, contractCode string) ([]ContractInfo, error) {
	params := url.Values{}
	params.Set("contract_code", contractCode)
	contracts := []ContractInfo{}
	err := cli.rsClient.Get(ctx, "/linear-swap-api/v1/swap_contract_info", params, &contracts)
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

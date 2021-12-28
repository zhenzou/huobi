package futures

import (
	"context"
	"fmt"
	"time"

	"github.com/zhenzou/huobi/client"
	"github.com/zhenzou/huobi/config"
)

func NewIndexClient(config config.Config) *IndexClient {
	return &IndexClient{
		wsClient: client.NewWebSocketClient(config, "/ws_index"),
	}
}

type IndexClient struct {
	wsClient *client.WebSocketClient
}

// -------------------------------------------------------------
// premium index kline start
func (client *IndexClient) SubPremiumIndexKLine(ctx context.Context, contractCode string, interval string) (<-chan SubIndexKLineResponse, error) {
	topic := fmt.Sprintf("market.%s.premium_index.%s", contractCode, interval)
	receiver := make(chan SubIndexKLineResponse, 1024)
	err := client.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (client *IndexClient) UnsubPremiumIndexKLine(ctx context.Context, contractCode string, period string) error {
	topic := fmt.Sprintf("market.%s.premium_index.%s", contractCode, period)
	return client.wsClient.Unsub(ctx, topic)
}

func (client *IndexClient) ReqPremiumIndexKLine(ctx context.Context, contractCode string, period string, from, to time.Time) (resp ReqIndexKLineResponse, err error) {
	topic := fmt.Sprintf("market.%s.premium_index.%s", contractCode, period)
	err = client.wsClient.Req(ctx, topic, from, to, &resp)
	return
}

// premium index kline end
//-------------------------------------------------------------

// -------------------------------------------------------------
// estimated rate kline start
func (client *IndexClient) SubEstimatedRateKLine(ctx context.Context, contractCode string, period string) (<-chan SubIndexKLineResponse, error) {
	topic := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)
	receiver := make(chan SubIndexKLineResponse, 1024)
	err := client.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (client *IndexClient) UnsubEstimatedRateKLine(ctx context.Context, contractCode string, period string) error {
	topic := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)

	return client.wsClient.Unsub(ctx, topic)
}

func (client *IndexClient) ReqEstimatedRateKLine(ctx context.Context, contractCode string, period string,
	from, to time.Time) (resp ReqIndexKLineResponse, err error) {
	topic := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)
	err = client.wsClient.Req(ctx, topic, from, to, &resp)
	return
}

// estimated rate kline end
//-------------------------------------------------------------

// -------------------------------------------------------------
// basis start
func (client *IndexClient) SubBasis(ctx context.Context,
	contractCode string, period string, basisPriceType string) (<-chan SubBasisResponse, error) {

	if basisPriceType == "" {
		basisPriceType = "open"
	}
	topic := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, basisPriceType)

	receiver := make(chan SubBasisResponse, 1024)
	err := client.wsClient.Sub(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (client *IndexClient) UnsubBasis(ctx context.Context, contractCode string, period string, basisPriceType string) error {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	topic := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, basisPriceType)

	return client.wsClient.Unsub(ctx, topic)
}

func (client *IndexClient) ReqBasis(ctx context.Context,
	contractCode string, period string, from, to time.Time, priceType string) (resp ReqBasisResponse, err error) {
	if priceType == "" {
		priceType = "open"
	}
	topic := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, priceType)
	err = client.wsClient.Req(ctx, topic, from, to, &resp)
	return
}

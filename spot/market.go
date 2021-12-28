package spot

import (
	"context"
	"fmt"

	"github.com/zhenzou/huobi/client"
	"github.com/zhenzou/huobi/config"
)

func NewMarketClient(cfg config.Config) *MarketClient {
	return &MarketClient{
		wsClient: client.NewWebSocketClient(cfg, "/ws"),
	}
}

type MarketClient struct {
	wsClient *client.WebSocketClient
}

func (client *MarketClient) SubTradeDetail(ctx context.Context, symbol string) (<-chan SubscribeTradeResponse, error) {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	ch := make(chan SubscribeTradeResponse, 1024)
	err := client.wsClient.Sub(ctx, topic, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (client *MarketClient) UnsubTradeDetail(ctx context.Context, symbol string) error {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	return client.wsClient.Unsub(ctx, topic)
}

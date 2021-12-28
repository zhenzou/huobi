package futures

import (
	"context"
	"fmt"

	"github.com/zhenzou/huobi/client"
	"github.com/zhenzou/huobi/config"
)

func NewAccountClient(cfg config.Config) *AccountClient {
	return &AccountClient{
		wsClient: client.NewWebSocketClient(cfg, "/linear-swap-notification"),
	}
}

type AccountClient struct {
	wsClient *client.WebSocketClient
}

func (client *AccountClient) SubAccounts(ctx context.Context, contractCode string) (<-chan SubAccountsResponse, error) {
	topic := fmt.Sprintf("accounts.%s", contractCode)

	receiver := make(chan SubAccountsResponse, 1024)
	err := client.wsClient.SubWithOp(ctx, topic, receiver)
	if err != nil {
		return nil, err
	}
	return receiver, nil
}

func (client *AccountClient) UnsubAccounts(ctx context.Context, contractCode string) error {

	topic := fmt.Sprintf("accounts.%s", contractCode)

	return client.wsClient.UnsubWithOP(ctx, topic)
}

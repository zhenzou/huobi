package huobi

import (
	"github.com/zhenzou/huobi/config"
	"github.com/zhenzou/huobi/futures"
	"github.com/zhenzou/huobi/spot"
)

const (
	DecimalPrecise = 6
)

func NewFutureAPi(cfg config.Config) *FuturesApi {
	if cfg.Host == "" {
		cfg.Host = config.DefaultFuturesHost
	}
	return &FuturesApi{
		Market:  futures.NewMarketClient(cfg),
		Order:   futures.NewOrderClient(cfg),
		Account: futures.NewAccountClient(cfg),
		Index:   futures.NewIndexClient(cfg),
	}
}

// USDT 永续合约
type FuturesApi struct {
	Market  *futures.MarketClient
	Order   *futures.OrderClient
	Account *futures.AccountClient
	Index   *futures.IndexClient
}

func NewSpotAPI(cfg config.Config) *SpotAPI {
	if cfg.Host == "" {
		cfg.Host = config.DefaultSpotHost
	}
	return &SpotAPI{
		Market: spot.NewMarketClient(cfg),
	}
}

type SpotAPI struct {
	Market *spot.MarketClient
}

type Logger interface {
}

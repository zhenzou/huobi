package main

import (
	"context"
	"log"
	"os"

	"github.com/zhenzou/huobi"
	"github.com/zhenzou/huobi/common"
	"github.com/zhenzou/huobi/config"
	"github.com/zhenzou/huobi/futures"
)

var (
	futureAPi = huobi.NewFutureAPi(config.Config{
		AccessKey: os.Getenv("HUOBI_ACCESS_KEY"),
		SecretKey: os.Getenv("HUOBI_SECRET_KEY"),
		Debug:     true,
	})
	spotAPI = huobi.NewSpotAPI(config.Config{
		Host: "api.huobi.pro",
	})
)

func GetContracts() {
	contracts, err := futureAPi.Market.ListContractInfos(context.Background(), "DOGE-USDT")
	if err != nil {
		panic(err)
	}
	for _, c := range contracts {
		log.Println("contract ", c)
	}
}

func SubscribeFutureTrade(id string) {
	ch, err := futureAPi.Market.SubTradeDetail(context.Background(), "DOGE-USDT")
	if err != nil {
		panic(err)
	}

	for item := range ch {
		log.Println("future ", item)
	}
}

func SubscribeSpotTrade() {
	ch, err := spotAPI.Market.SubTradeDetail(context.Background(), "dogeusdt")
	if err != nil {
		panic(err)
	}
	for resp := range ch {
		log.Println("spot ", resp)
	}
}

func SubscribeKLine() {
	ch, err := futureAPi.Market.SubKLine(context.Background(), "DOGE-USDT", "1min")
	if err != nil {
		panic(err)
	}

	for item := range ch {
		log.Println("kline ", item)
	}
}

func SubscribeBasis() {
	ch, err := futureAPi.Index.SubBasis(context.Background(), "DOGE-USDT", "1min", "")
	if err != nil {
		panic(err)
	}

	for item := range ch {
		log.Println("kline ", item)
	}
}

func RequestTrade() {
	resp, err := futureAPi.Market.ReqTradeDetail(context.Background(), "DOGE-USDT")
	if err != nil {
		panic(err)
	}
	for _, datum := range resp.Data {
		log.Println("trade ", datum)
	}
}

func SubscribeOrder() {
	ch, err := futureAPi.Order.SubOrders(context.Background(), "DOGE-USDT")
	if err != nil {
		panic(err)
	}
	for order := range ch {
		log.Println("order ", order)
	}
}

func ListOpenTriggerOrders() {
	orders, _, err := futureAPi.Order.ListOpenTriggerOrders(context.Background(), futures.ListOrdersRequest{
		ContractCode: "DOGE-USDT",
	})
	if err != nil {
		panic(err)
	}
	for _, order := range orders {
		log.Println("order ", order)
	}
}

func ListOrders() {
	orders, _, err := futureAPi.Order.ListHistoryOrders(context.Background(), futures.ListOrdersRequest{
		ContractCode: "DOGE-USDT",
		Type:         1,
	})
	if err != nil {
		panic(err)
	}
	for _, order := range orders {
		log.Println("order ", order)
	}
}

func CreateOrder(ctx context.Context) string {
	id, err := futureAPi.Order.CreateTriggerOrder(ctx, futures.CreateTriggerOrderRequest{
		ContractCode:   "DOGE-USDT",
		TriggerType:    "ge",
		TriggerPrice:   0.08602,
		Volume:         10,
		Direction:      "buy",
		Offset:         "open",
		LeverRate:      2,
		OrderPriceType: "optimal_5",
	})
	if err != nil {
		panic(err)
	}
	return id
}

func CancelOrder(id string) {
	err := futureAPi.Order.CancelTriggerOrder(context.Background(), "DOGE-USDT", id)
	if err != nil {
		panic(err)
	}
	println("canceled:" + id)
}

func main() {
	//log.SetLevel(log.DebugLevel)

	go SubscribeFutureTrade("1")
	go SubscribeSpotTrade()
	//go SubscribeFutureTrade("2")
	//RequestTrade()
	//SubscribeOrder()
	//SubscribeKLine()
	//SubscribeBasis()
	//ListOpenTriggerOrders()
	//ListOrders()

	//orderID := CreateOrder(context.Background())

	common.WaitStopSignal()
}

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
)

//client := binance.NewClient(APIKEY, SECRITKEY)

func main() {
	client := binance.NewFuturesClient(APIKEY, SECRITKEY) // USDT-M Futures

	res, err := client.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	wsDepthHandler := func(event *binance.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := binance.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC
}

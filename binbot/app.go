package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

const APIKEY = "2k82QayDqeL36edeO3kkUPFTQrqGyICzPYRmw2cC3BHsSjteaIE0k2dcuGoyY4dU"
const SECRITKEY = "tpwei5t7NOdvYJx1qhKlQ3VqZincxuo6NVg4jdGmo2kMrmhkwcvYoDrdeWp1mpJz"

//client := binance.NewClient(APIKEY, SECRITKEY)

func main() {
	client := binance.NewFuturesClient(APIKEY, SECRITKEY) // USDT-M Futures
    
    orders, err := client.NewCreateOrderService().Symbol("ONEUSDT").Size().Side(futures.SideTypeBuy).Type(futures.OrderTypeMarket).Do(context.Background())
    if err != nil {fmt.Println(err);return}

    fmt.Println(orders)
}

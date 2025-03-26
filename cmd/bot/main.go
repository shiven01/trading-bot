package main

import (
	"fmt"
	"time"
)

import "github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"

func main() {
	fmt.Println("Running Trading Bot")

	client := marketdata.NewClient(marketdata.ClientOpts{})

	request := marketdata.GetCryptoBarsRequest{
		TimeFrame: marketdata.OneDay,
		Start:     time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC),
		End:       time.Date(2022, 9, 7, 0, 0, 0, 0, time.UTC),
	}

	bars, err := client.GetCryptoBars("BTC/USD", request)
	if err != nil {
		panic(err)
	}
	for _, bar := range bars {
		fmt.Printf("%+v\n", bar)
	}

	fmt.Println("Shutting Down Trading Bot")
}

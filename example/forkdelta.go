package main

import (
	"log"

	ed "github.com/alvinlaw/go-etherdelta"
)

func main() {
	service := ed.NewForkDelta(&ed.Options{
		ProviderURI: "https://cloudflare-eth.com",
	})

	orders, err := service.GetOrderBook(&ed.GetOrderBookOpts{
		TokenAddress: "0x8f3470a7388c05ee4e7af3d01d8c722b0ff52374",
	})

	if err != nil {
		panic(err)
	}

	log.Println(orders)
}

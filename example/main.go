package main

import (
	"log"

	go_stockx_client "github.com/bogdanfinn/go-stockx-client"
)

func main() {
	client, err := go_stockx_client.ProvideClient("USD", go_stockx_client.NewNoopLogger())

	if err != nil {
		log.Println(err.Error())
		return
	}

	stockSearchResults, err := client.SearchProducts("yeezy zebra", 10)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(stockSearchResults)

	productDetails, err := client.GetProduct(stockSearchResults[0].Urlkey)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(productDetails)
}

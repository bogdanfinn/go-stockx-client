package main

import (
	"fmt"
	"log"

	go_stockx_client "github.com/bogdanfinn/go-stockx-client"
)

func main() {
	// NewClient() returns each time a new instance
	// Provide() is creating one client instance and returning the same instance on every Provide() call
	client, err := go_stockx_client.ProvideClient("USD", go_stockx_client.NewNoopLogger())
	// client, err := go_stockx_client.NewClient("USD", go_stockx_client.NewNoopLogger())

	if err != nil {
		log.Println(err.Error())
		return
	}

	query := "adidas yeezy foam rnnr onyx"
	stockSearchResults, err := client.SearchProducts(query, 10)

	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(stockSearchResults) == 0 {
		log.Println(fmt.Sprintf("did not find any product for search query %s", query))
		return
	}

	log.Println(fmt.Sprintf("received %d search results for query %s", len(stockSearchResults), query))

	productDetails, err := client.GetProduct(stockSearchResults[0].ProductIdentifier)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(fmt.Sprintf("successfully loaded product details for %s", productDetails.ProductIdentifier))
}

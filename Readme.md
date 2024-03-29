# go-stockx-client

### Preface
This is a go library which should help you scraping stockx product data. The library does not use official stockx apis and might break due to unexpected changes from stockx.

Update: As stated above currently some endpoints are not working anymore. Pull Requests are welcome.

### Installation

```go
go get -u github.com/bogdanfinn/go-stockx-client

// or specific version:
// go get -u github.com/bogdanfinn/go-stockx-client@v0.1.1
```

### Quick Example
```go
package main

import (
	"fmt"
	"log"

	go_stockx_client "github.com/bogdanfinn/go-stockx-client"
)

func main() {
	// NewClient() returns each time a new instance
	// Provide() is creating one client instance and returning the same instance on every Provide() call
	client, err := go_stockx_client.ProvideClient("USD", go_stockx_client.NewNoopLogger(), false)
    // client, err := go_stockx_client.NewClient("USD", go_stockx_client.NewNoopLogger())

	if err != nil {
		log.Println(err.Error())
		return
	}

	query := "yeezy zebra"
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

	log.Println(fmt.Sprintf("successfully loaded product details for %s", productDetails))
}
```

### Supported Methods
Method | Description                                                                                                                                                        | Arguments                         | Return Value
--- |--------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------|--------------
NewClient | Creates a new client instance. Takes a currency string (for example `"USD"`) and a logger which implements the logger interface as parameters. Or returns an error | `currency string`, `logger Logger`, `vatAccount bool` | `Client`, `error` 
SearchProducts | Search Stockx for products based on the given search query and returns search results up to the provided limit argument or less. Or returns an error               | `searchQuery string`, `limit: int` | `[]SearchResultProduct`, `error` 
GetProduct | Scrapes product details for a given product identifier which you get from the search results. Or returns an error                                                  | `productIdentifier: string`       | `*ProductDetails`, `error`       
SetProxy | Sets a Proxy to be used for the Stockx Client. Use this pattern `http://username:password@host:port`                                                               | `proxyUrl: string`                | `error`       
GetProxy | Returns the current active proxy for the stockx client.   |      | `string`       

### Types

#### Client
```go
type Client interface {
	SearchProducts(query string, limit int) ([]SearchResultProduct, error)
	GetProduct(productIdentifier string) (*ProductDetails, error)
}
```

#### SearchResultProduct
```go
type SearchResultProduct struct {
	Brand             string `json:"brand"`
	Colorway          string `json:"colorway"`
	ImageUrl          string `json:"imageUrl"`
	Category          string `json:"category"`
	Description       string `json:"description"`
	Title             string `json:"title"`
	ProductIdentifier string `json:"productIdentifier"`
}
```

#### ProductDetails & ProductDetailsVariant
```go
type ProductDetails struct {
    ID                string                  `json:"id"`
    UUID              string                  `json:"uuid"`
    Brand             string                  `json:"brand"`
    Colorway          string                  `json:"colorway"`
    Minimumbid        int                     `json:"minimumBid"`
    Name              string                  `json:"name"`
    Releasedate       string                  `json:"releaseDate"`
    Retailprice       int                     `json:"retailPrice"`
    Shoe              string                  `json:"shoe"`
    SizeLocale        string                  `json:"sizeLocale"`
    SizeTitle         string                  `json:"sizeTitle"`
    Shortdescription  string                  `json:"shortDescription"`
    Styleid           string                  `json:"styleId"`
    Title             string                  `json:"title"`
    ProductIdentifier string                  `json:"productIdentifier"`
    Description       string                  `json:"description"`
    Imageurl          string                  `json:"imageUrl"`
    Smallimageurl     string                  `json:"smallImageUrl"`
    Thumburl          string                  `json:"thumbUrl"`
    Lowestask         int                     `json:"lowestAsk"`
    Highestbid        int                     `json:"highestBid"`
    Lowestaskfloat    float64                 `json:"lowestAskFloat"`
    Highestbidfloat   float64                 `json:"highestBidFloat"`
    Variants          []ProductDetailsVariant `json:"variants"`
}

type ProductDetailsVariant struct {
    UUID             string    `json:"UUID"`
    Size             string    `json:"size"`
    Lowestask        int       `json:"lowestAsk"`
    Highestbid       int       `json:"highestBid"`
    Annualhigh       int       `json:"annualHigh"`
    Annuallow        int       `json:"annualLow"`
    Lastsale         int       `json:"lastSale"`
    Saleslast72Hours int       `json:"salesLast72Hours"`
    Lastsaledate     time.Time `json:"lastSaleDate"`
    Lowestaskfloat   float64   `json:"lowestAskFloat"`
    Highestbidfloat  float64   `json:"highestBidFloat"`
}

```

### Stockx Currency & Country
You can use these values for `currency` and `locale` when creating a new client.

```json
{
  "currencies": {
    "AUD": {
      "countries": [
        "AU",
        "CX",
        "NF",
        "HM",
        "NR",
        "CC"
      ]
    },
    "CAD": {
      "countries": [
        "CA"
      ]
    },
    "CHF": {
      "countries": [
        "CH",
        "LI"
      ]
    },
    "EUR": {
      "countries": [
        "YT",
        "PM",
        "DE",
        "AX",
        "BL",
        "ES",
        "TF",
        "FX",
        "VA",
        "GF",
        "AT",
        "MC",
        "GP",
        "LU",
        "GR",
        "RE",
        "MF",
        "IT",
        "NL",
        "ME",
        "XK",
        "SI",
        "CY",
        "LV",
        "PT",
        "FR",
        "FI",
        "MQ",
        "SM",
        "AD",
        "LT",
        "BE",
        "SK",
        "IE",
        "EE",
        "MT",
        "PL"
      ]
    },
    "GBP": {
      "countries": [
        "PN",
        "GS",
        "AC",
        "GB"
      ]
    },
    "HKD": {
      "countries": [
        "HK"
      ]
    },
    "JPY": {
      "countries": [
        "JP"
      ]
    },
    "KRW": {
      "countries": [
        "KR"
      ]
    },
    "MXN": {
      "countries": [
        "MX"
      ]
    },
    "NZD": {
      "countries": [
        "NZ",
        "NU"
      ]
    },
    "SGD": {
      "countries": [
        "SG"
      ]
    },
    "USD": {
      "countries": [
        "US"
      ]
    }
  }
}
```

### Frequently Asked Questions / Errors
TBD

### Questions?
Join my discord support server: https: // discord.gg / 7Ej9eJvHqk 
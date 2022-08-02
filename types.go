package go_stockx_client

import "time"

type SearchResultProduct struct {
	Brand             string `json:"brand"`
	Colorway          string `json:"colorway"`
	ImageUrl          string `json:"imageUrl"`
	Category          string `json:"category"`
	Description       string `json:"description"`
	Title             string `json:"title"`
	ProductIdentifier string `json:"productIdentifier"`
}

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

package go_stockx_client

import "time"

type SearchResultProductResponse struct {
	Brand    string `json:"brand"`
	Colorway string `json:"colorway"`
	Media    struct {
		Thumburl string `json:"thumbUrl"`
	} `json:"media"`
	Productcategory  string `json:"productCategory"`
	Shortdescription string `json:"shortDescription"`
	Title            string `json:"title"`
	Urlkey           string `json:"urlKey"`
	Objectid         string `json:"objectID"`
}

type ProductSearchResultResponse struct {
	Pagination struct {
		Query        string      `json:"query"`
		Queryid      string      `json:"queryID"`
		Index        string      `json:"index"`
		Limit        string      `json:"limit"`
		Page         int         `json:"page"`
		Total        int         `json:"total"`
		Lastpage     string      `json:"lastPage"`
		Sort         []string    `json:"sort"`
		Order        []string    `json:"order"`
		Currentpage  string      `json:"currentPage"`
		Nextpage     interface{} `json:"nextPage"`
		Previouspage interface{} `json:"previousPage"`
	} `json:"Pagination"`
	Facets struct {
		Browseverticals struct {
			Sneakers int `json:"sneakers"`
		} `json:"browseVerticals"`
	} `json:"Facets"`
	Products []SearchResultProductResponse `json:"Products"`
}

type ProductResponse struct {
	Product Product `json:"Product"`
}

type Product struct {
	ID                   string        `json:"id"`
	UUID                 string        `json:"uuid"`
	Brand                string        `json:"brand"`
	Colorway             string        `json:"colorway"`
	Condition            string        `json:"condition"`
	Countryofmanufacture string        `json:"countryOfManufacture"`
	Gender               string        `json:"gender"`
	Contentgroup         string        `json:"contentGroup"`
	Minimumbid           int           `json:"minimumBid"`
	Name                 string        `json:"name"`
	Primarycategory      string        `json:"primaryCategory"`
	Secondarycategory    string        `json:"secondaryCategory"`
	Ushtscode            string        `json:"usHtsCode"`
	Ushtsdescription     string        `json:"usHtsDescription"`
	Productcategory      string        `json:"productCategory"`
	Releasedate          string        `json:"releaseDate"`
	Retailprice          int           `json:"retailPrice"`
	Shoe                 string        `json:"shoe"`
	Shortdescription     string        `json:"shortDescription"`
	Styleid              string        `json:"styleId"`
	Tickersymbol         string        `json:"tickerSymbol"`
	Title                string        `json:"title"`
	Datatype             string        `json:"dataType"`
	Urlkey               string        `json:"urlKey"`
	Sizelocale           string        `json:"sizeLocale"`
	Sizetitle            string        `json:"sizeTitle"`
	Sizedescriptor       string        `json:"sizeDescriptor"`
	Sizealldescriptor    string        `json:"sizeAllDescriptor"`
	Description          string        `json:"description"`
	Lithiumionbattery    bool          `json:"lithiumIonBattery"`
	Hazardousmaterial    interface{}   `json:"hazardousMaterial"`
	Type                 bool          `json:"type"`
	Alim                 int           `json:"aLim"`
	Year                 int           `json:"year"`
	Shippinggroup        string        `json:"shippingGroup"`
	Portfolioitems       []interface{} `json:"PortfolioItems"`
	Shipping             struct {
		Totaldaystoship         int  `json:"totalDaysToShip"`
		Hasadditionaldaystoship bool `json:"hasAdditionalDaysToShip"`
		Deliverydayslowerbound  int  `json:"deliveryDaysLowerBound"`
		Deliverydaysupperbound  int  `json:"deliveryDaysUpperBound"`
	} `json:"shipping"`
	Enhancedimage struct {
		Productuuid string `json:"productUuid"`
		Imagekey    string `json:"imageKey"`
		Imagecount  int    `json:"imageCount"`
	} `json:"enhancedImage"`
	Media struct {
		Num360        []string      `json:"360"`
		Imageurl      string        `json:"imageUrl"`
		Smallimageurl string        `json:"smallImageUrl"`
		Thumburl      string        `json:"thumbUrl"`
		Has360        bool          `json:"has360"`
		Gallery       []interface{} `json:"gallery"`
	} `json:"media"`
	Charitycondition int `json:"charityCondition"`
	Breadcrumbs      []struct {
		Level int    `json:"level"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"breadcrumbs"`
	Market struct {
		Productid                 int         `json:"productId"`
		Skuuuid                   interface{} `json:"skuUuid"`
		Productuuid               string      `json:"productUuid"`
		Lowestask                 int         `json:"lowestAsk"`
		Lowestasksize             interface{} `json:"lowestAskSize"`
		Parentlowestask           int         `json:"parentLowestAsk"`
		Numberofasks              int         `json:"numberOfAsks"`
		Hasasks                   int         `json:"hasAsks"`
		Salesthisperiod           int         `json:"salesThisPeriod"`
		Saleslastperiod           int         `json:"salesLastPeriod"`
		Highestbid                int         `json:"highestBid"`
		Highestbidsize            interface{} `json:"highestBidSize"`
		Numberofbids              int         `json:"numberOfBids"`
		Hasbids                   int         `json:"hasBids"`
		Annualhigh                int         `json:"annualHigh"`
		Annuallow                 int         `json:"annualLow"`
		Deadstockrangelow         int         `json:"deadstockRangeLow"`
		Deadstockrangehigh        int         `json:"deadstockRangeHigh"`
		Volatility                float64     `json:"volatility"`
		Deadstocksold             int         `json:"deadstockSold"`
		Pricepremium              float64     `json:"pricePremium"`
		Averagedeadstockprice     int         `json:"averageDeadstockPrice"`
		Lastsale                  int         `json:"lastSale"`
		Lastsalesize              string      `json:"lastSaleSize"`
		Saleslast72Hours          int         `json:"salesLast72Hours"`
		Changevalue               int         `json:"changeValue"`
		Changepercentage          float64     `json:"changePercentage"`
		Abschangepercentage       float64     `json:"absChangePercentage"`
		Totaldollars              int         `json:"totalDollars"`
		Updatedat                 int         `json:"updatedAt"`
		Lastlowestasktime         int         `json:"lastLowestAskTime"`
		Lasthighestbidtime        int         `json:"lastHighestBidTime"`
		Lastsaledate              time.Time   `json:"lastSaleDate"`
		Createdat                 time.Time   `json:"createdAt"`
		Deadstocksoldrank         int         `json:"deadstockSoldRank"`
		Pricepremiumrank          int         `json:"pricePremiumRank"`
		Averagedeadstockpricerank int         `json:"averageDeadstockPriceRank"`
		Featured                  interface{} `json:"featured"`
		Lowestaskfloat            float64     `json:"lowestAskFloat"`
		Highestbidfloat           float64     `json:"highestBidFloat"`
	} `json:"market"`
	Children map[string]ProductWithoutChildren `json:"children"`
}

type ProductWithoutChildren struct {
	ID                   string        `json:"id"`
	UUID                 string        `json:"uuid"`
	Brand                string        `json:"brand"`
	Colorway             string        `json:"colorway"`
	Condition            string        `json:"condition"`
	Countryofmanufacture string        `json:"countryOfManufacture"`
	Gender               string        `json:"gender"`
	Contentgroup         string        `json:"contentGroup"`
	Minimumbid           int           `json:"minimumBid"`
	Name                 string        `json:"name"`
	Primarycategory      string        `json:"primaryCategory"`
	Secondarycategory    string        `json:"secondaryCategory"`
	Ushtscode            string        `json:"usHtsCode"`
	Ushtsdescription     string        `json:"usHtsDescription"`
	Productcategory      string        `json:"productCategory"`
	Releasedate          string        `json:"releaseDate"`
	Retailprice          int           `json:"retailPrice"`
	Shoe                 string        `json:"shoe"`
	Shortdescription     string        `json:"shortDescription"`
	Styleid              string        `json:"styleId"`
	Tickersymbol         string        `json:"tickerSymbol"`
	Title                string        `json:"title"`
	Datatype             string        `json:"dataType"`
	Urlkey               string        `json:"urlKey"`
	Sizelocale           string        `json:"sizeLocale"`
	Sizetitle            string        `json:"sizeTitle"`
	Sizedescriptor       string        `json:"sizeDescriptor"`
	Sizealldescriptor    string        `json:"sizeAllDescriptor"`
	Description          string        `json:"description"`
	Lithiumionbattery    bool          `json:"lithiumIonBattery"`
	Hazardousmaterial    interface{}   `json:"hazardousMaterial"`
	Type                 bool          `json:"type"`
	Alim                 int           `json:"aLim"`
	Year                 int           `json:"year"`
	Shippinggroup        string        `json:"shippingGroup"`
	Portfolioitems       []interface{} `json:"PortfolioItems"`
	Shipping             struct {
		Totaldaystoship         int  `json:"totalDaysToShip"`
		Hasadditionaldaystoship bool `json:"hasAdditionalDaysToShip"`
		Deliverydayslowerbound  int  `json:"deliveryDaysLowerBound"`
		Deliverydaysupperbound  int  `json:"deliveryDaysUpperBound"`
	} `json:"shipping"`
	Enhancedimage struct {
		Productuuid string `json:"productUuid"`
		Imagekey    string `json:"imageKey"`
		Imagecount  int    `json:"imageCount"`
	} `json:"enhancedImage"`
	Media struct {
		Num360        []string      `json:"360"`
		Imageurl      string        `json:"imageUrl"`
		Smallimageurl string        `json:"smallImageUrl"`
		Thumburl      string        `json:"thumbUrl"`
		Has360        bool          `json:"has360"`
		Gallery       []interface{} `json:"gallery"`
	} `json:"media"`
	Charitycondition int `json:"charityCondition"`
	Breadcrumbs      []struct {
		Level int    `json:"level"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"breadcrumbs"`
	Market struct {
		Productid                 int         `json:"productId"`
		Skuuuid                   interface{} `json:"skuUuid"`
		Productuuid               string      `json:"productUuid"`
		Lowestask                 int         `json:"lowestAsk"`
		Lowestasksize             interface{} `json:"lowestAskSize"`
		Parentlowestask           int         `json:"parentLowestAsk"`
		Numberofasks              int         `json:"numberOfAsks"`
		Hasasks                   int         `json:"hasAsks"`
		Salesthisperiod           int         `json:"salesThisPeriod"`
		Saleslastperiod           int         `json:"salesLastPeriod"`
		Highestbid                int         `json:"highestBid"`
		Highestbidsize            interface{} `json:"highestBidSize"`
		Numberofbids              int         `json:"numberOfBids"`
		Hasbids                   int         `json:"hasBids"`
		Annualhigh                int         `json:"annualHigh"`
		Annuallow                 int         `json:"annualLow"`
		Deadstockrangelow         int         `json:"deadstockRangeLow"`
		Deadstockrangehigh        int         `json:"deadstockRangeHigh"`
		Volatility                float64     `json:"volatility"`
		Deadstocksold             int         `json:"deadstockSold"`
		Pricepremium              float64     `json:"pricePremium"`
		Averagedeadstockprice     int         `json:"averageDeadstockPrice"`
		Lastsale                  int         `json:"lastSale"`
		Lastsalesize              string      `json:"lastSaleSize"`
		Saleslast72Hours          int         `json:"salesLast72Hours"`
		Changevalue               int         `json:"changeValue"`
		Changepercentage          float64     `json:"changePercentage"`
		Abschangepercentage       float64     `json:"absChangePercentage"`
		Totaldollars              int         `json:"totalDollars"`
		Updatedat                 int         `json:"updatedAt"`
		Lastlowestasktime         int         `json:"lastLowestAskTime"`
		Lasthighestbidtime        int         `json:"lastHighestBidTime"`
		Lastsaledate              time.Time   `json:"lastSaleDate"`
		Createdat                 time.Time   `json:"createdAt"`
		Deadstocksoldrank         int         `json:"deadstockSoldRank"`
		Pricepremiumrank          int         `json:"pricePremiumRank"`
		Averagedeadstockpricerank int         `json:"averageDeadstockPriceRank"`
		Featured                  interface{} `json:"featured"`
		Lowestaskfloat            float64     `json:"lowestAskFloat"`
		Highestbidfloat           float64     `json:"highestBidFloat"`
	} `json:"market"`
}

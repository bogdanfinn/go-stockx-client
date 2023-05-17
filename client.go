package go_stockx_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"sync"

	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/fhttp/cookiejar"
	tls_client "github.com/bogdanfinn/tls-client"
)

const stockxBaseUrl = "https://stockx.com/"
const stockxSearchEndpointTemplate = "https://stockx.com/api/browse?_search=%s&page=1&resultsPerPage=%d&dataType=product&facetsToRetrieve[]=browseVerticals&propsToRetrieve[][]=brand&propsToRetrieve[][]=colorway&propsToRetrieve[][]=media.thumbUrl&propsToRetrieve[][]=title&propsToRetrieve[][]=productCategory&propsToRetrieve[][]=shortDescription&propsToRetrieve[][]=urlKey"
const stockxProductDetailsEndpointTemplate = "https://stockx.com/api/products/%s?includes=market&currency=%s&country=%s&market=%s"

var stockxHeader = http.Header{
	"accept":             {"application/json"},
	"accept-language":    {"de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7"},
	"app-platform":       {"Iron"},
	"app-version":        {"2022.07.17.01"},
	"cache-control":      {"no-cache"},
	"pragma":             {"no-cache"},
	"referer":            {"https://stockx.com/de-de"},
	"sec-ch-ua":          {`".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103"`},
	"sec-ch-ua-mobile":   {"?0"},
	"sec-ch-ua-platform": {`"macOS"`},
	"sec-fetch-dest":     {"empty"},
	"sec-fetch-mode":     {"cors"},
	"sec-fetch-site":     {"same-origin"},
	"user-agent":         {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"},
	"x-requested-with":   {"XMLHttpRequest"},
	http.HeaderOrderKey: {
		"accept",
		"accept-language",
		"app-platform",
		"app-version",
		"cache-control",
		"pragma",
		"referer",
		"sec-ch-ua",
		"sec-ch-ua-mobile",
		"sec-ch-ua-platform",
		"sec-fetch-dest",
		"sec-fetch-mode",
		"sec-fetch-site",
		"user-agent",
		"x-requested-with",
	},
}

type Client interface {
	SearchProducts(query string, limit int) ([]SearchResultProduct, error)
	GetProduct(productIdentifier string) (*ProductDetails, error)
	SetProxy(proxyUrl string) error
	GetProxy() string
}

type client struct {
	initialized bool
	logger      Logger
	currency    string
	locale      string
	httpClient  tls_client.HttpClient
	vatAccount  bool
}

var clientContainer = struct {
	sync.Mutex
	instance Client
}{}

func ProvideClient(currency string, locale string, logger Logger, vatAccount bool) (Client, error) {
	clientContainer.Lock()
	defer clientContainer.Unlock()

	if clientContainer.instance != nil {
		return clientContainer.instance, nil
	}

	instance, err := NewClient(currency, locale, logger, vatAccount)

	if err != nil {
		return nil, err
	}

	clientContainer.instance = instance

	return clientContainer.instance, nil
}

func NewClient(currency string, locale string, logger Logger, vatAccount bool) (Client, error) {
	jar, _ := cookiejar.New(nil)

	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(30),
		tls_client.WithClientProfile(tls_client.Chrome_105),
		tls_client.WithCookieJar(jar),
		// tls_client.WithNotFollowRedirects(),
	}

	httpClient, err := tls_client.NewHttpClient(logger, options...)

	if err != nil {
		return nil, fmt.Errorf("failed to construct http client: %w", err)
	}

	return &client{
		initialized: false,
		logger:      logger,
		currency:    strings.ToUpper(currency),
		locale:      strings.ToUpper(locale),
		httpClient:  httpClient,
		vatAccount:  vatAccount,
	}, nil
}

func (c *client) initialize() error {
	if c.initialized {
		return nil
	}

	statusCode, _, err := c.doRequest(stockxBaseUrl, stockxHeader)

	if err != nil {
		return fmt.Errorf("failed to initialize client: %w", err)
	}

	if statusCode == http.StatusOK {
		c.initialized = true
		return nil
	}

	return fmt.Errorf("received wrong status code during client initialization: %d", statusCode)
}

func (c *client) SetProxy(proxyUrl string) error {
	return c.httpClient.SetProxy(proxyUrl)
}

func (c *client) GetProxy() string {
	return c.httpClient.GetProxy()
}

func (c *client) SearchProducts(query string, limit int) ([]SearchResultProduct, error) {
	err := c.initialize()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize client: %w", err)
	}

	preparedQuery := query
	if !strings.Contains(query, "+") && strings.Contains(query, " ") {
		queryParts := strings.Split(query, " ")
		preparedQuery = strings.Join(queryParts, "+")
	}

	searchUrl := fmt.Sprintf(stockxSearchEndpointTemplate, preparedQuery, limit)

	_, respBodyBytes, err := c.doRequest(searchUrl, stockxHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	response := ProductSearchResultResponse{}
	err = json.Unmarshal(respBodyBytes, &response)

	if err != nil {
		return nil, fmt.Errorf("failed to convert response json into response struct: %w", err)
	}

	searchResultProducts := parseSearchResults(response)

	return searchResultProducts, nil
}

func (c *client) GetProduct(productIdentifier string) (*ProductDetails, error) {
	err := c.initialize()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize client: %w", err)
	}

	productUrl := fmt.Sprintf(stockxProductDetailsEndpointTemplate, productIdentifier, c.currency, c.locale, c.locale)
	if c.vatAccount {
		productUrl = fmt.Sprintf(stockxProductDetailsEndpointTemplate, productIdentifier, c.currency, c.locale, fmt.Sprintf("%s.vat-registered", c.locale))
	}
	_, respBodyBytes, err := c.doRequest(productUrl, stockxHeader)

	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	response := ProductResponse{}
	err = json.Unmarshal(respBodyBytes, &response)

	if err != nil {
		return nil, fmt.Errorf("failed to convert response json into response struct: %w", err)
	}

	product := parseProduct(response)

	return product, nil
}

func (c *client) doRequest(url string, header http.Header) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create stockx search request: %w", err)
	}

	req.Header = header

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return 0, nil, fmt.Errorf("failed to search for stockx products: %w", err)
	}

	c.logger.Info("stockx api (%s) response status code: %d", url, resp.StatusCode)

	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)

	c.logger.Debug("stockx api (%s) response body: %s", url, string(respBodyBytes))

	return resp.StatusCode, respBodyBytes, err
}

func parseSearchResults(response ProductSearchResultResponse) []SearchResultProduct {
	var searchResultProducts []SearchResultProduct

	for _, responseProduct := range response.Products {
		searchResultProducts = append(searchResultProducts, SearchResultProduct{
			Brand:             responseProduct.Brand,
			Colorway:          responseProduct.Colorway,
			ImageUrl:          responseProduct.Media.Thumburl,
			Category:          responseProduct.Productcategory,
			Description:       responseProduct.Shortdescription,
			Title:             responseProduct.Title,
			ProductIdentifier: responseProduct.Urlkey,
		})
	}

	return searchResultProducts
}

func parseProduct(response ProductResponse) *ProductDetails {
	var variants []ProductDetailsVariant

	product := response.Product

	for key, responseVariant := range product.Children {
		if responseVariant.Market.Lastsalesize == "" {
			continue
		}

		variants = append(variants, ProductDetailsVariant{
			UUID:             key,
			Size:             responseVariant.Market.Lastsalesize,
			Lowestask:        responseVariant.Market.Lowestask,
			Highestbid:       responseVariant.Market.Highestbid,
			Annualhigh:       responseVariant.Market.Annualhigh,
			Annuallow:        responseVariant.Market.Annuallow,
			Lastsale:         responseVariant.Market.Lastsale,
			Saleslast72Hours: responseVariant.Market.Saleslast72Hours,
			Lastsaledate:     responseVariant.Market.Lastsaledate,
			Lowestaskfloat:   responseVariant.Market.Lowestaskfloat,
			Highestbidfloat:  responseVariant.Market.Highestbidfloat,
		})
	}

	sort.Slice(variants, func(i, j int) bool {
		sizeA, errA := strconv.ParseFloat(variants[i].Size, 32)
		sizeB, errB := strconv.ParseFloat(variants[j].Size, 32)

		if errA != nil || errB != nil {
			return false
		}

		return sizeA < sizeB
	})

	return &ProductDetails{
		ID:                product.ID,
		UUID:              product.UUID,
		Brand:             product.Brand,
		Colorway:          product.Colorway,
		Minimumbid:        product.Minimumbid,
		Name:              product.Name,
		Releasedate:       product.Releasedate,
		Retailprice:       product.Retailprice,
		Shoe:              product.Shoe,
		Shortdescription:  product.Shortdescription,
		Styleid:           product.Styleid,
		Title:             product.Title,
		SizeLocale:        product.Sizelocale,
		SizeTitle:         product.Sizetitle,
		ProductIdentifier: product.Urlkey,
		Description:       product.Description,
		Imageurl:          product.Media.Imageurl,
		Smallimageurl:     product.Media.Smallimageurl,
		Thumburl:          product.Media.Thumburl,
		Lowestaskfloat:    product.Market.Lowestaskfloat,
		Lowestask:         product.Market.Lowestask,
		Highestbid:        product.Market.Highestbid,
		Highestbidfloat:   product.Market.Highestbidfloat,
		Variants:          variants,
	}
}

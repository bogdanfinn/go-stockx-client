package go_stockx_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

const stockxSearchEndpointTemplate = "https://stockx.com/api/browse?_search=%s&page=1&resultsPerPage=%d&dataType=product&facetsToRetrieve[]=browseVerticals&propsToRetrieve[][]=brand&propsToRetrieve[][]=colorway&propsToRetrieve[][]=media.thumbUrl&propsToRetrieve[][]=title&propsToRetrieve[][]=productCategory&propsToRetrieve[][]=shortDescription&propsToRetrieve[][]=urlKey"
const stockxProductDetailsEndpointTemplate = "https://stockx.com/api/products/%s?includes=market&currency=%s&country=US"

type Client interface {
	SearchProducts(query string, limit int) ([]SearchResultProduct, error)
	GetProduct(productIdentifier string) (*ProductDetails, error)
}

type client struct {
	logger     Logger
	currency   string
	httpClient tls_client.HttpClient
}

var clientContainer = struct {
	sync.Mutex
	instance Client
}{}

func ProvideClient(currency string, logger Logger) (Client, error) {
	clientContainer.Lock()
	defer clientContainer.Unlock()

	if clientContainer.instance != nil {
		return clientContainer.instance, nil
	}

	instance, err := NewClient(currency, logger)

	if err != nil {
		return nil, err
	}

	clientContainer.instance = instance

	return clientContainer.instance, nil
}

func NewClient(currency string, logger Logger) (Client, error) {
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(30),
		tls_client.WithClientProfile(tls_client.Chrome_103),
		tls_client.WithNotFollowRedirects(),
	}

	httpClient, err := tls_client.NewHttpClient(logger, options...)

	if err != nil {
		return nil, fmt.Errorf("failed to construct http client: %w", err)
	}

	return &client{
		logger:     logger,
		currency:   currency,
		httpClient: httpClient,
	}, nil
}

func (c *client) SearchProducts(query string, limit int) ([]SearchResultProduct, error) {
	preparedQuery := query
	if !strings.Contains(query, "+") && strings.Contains(query, " ") {
		queryParts := strings.Split(query, " ")
		preparedQuery = strings.Join(queryParts, "+")
	}

	searchUrl := fmt.Sprintf(stockxSearchEndpointTemplate, preparedQuery, limit)

	header := http.Header{
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

	respBodyBytes, err := c.doRequest(searchUrl, header)
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
	productUrl := fmt.Sprintf(stockxProductDetailsEndpointTemplate, productIdentifier, c.currency)

	header := http.Header{
		"accept":             {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
		"accept-encoding":    {"gzip, deflate, br"},
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

	respBodyBytes, err := c.doRequest(productUrl, header)

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

func (c *client) doRequest(url string, header http.Header) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create stockx search request: %w", err)
	}

	req.Header = header

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to search for stockx products: %w", err)
	}

	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)

	return respBodyBytes, err
}

func parseSearchResults(response ProductSearchResultResponse) []SearchResultProduct {
	var searchResultProducts []SearchResultProduct

	return searchResultProducts
}

func parseProduct(response ProductResponse) *ProductDetails {

	return &ProductDetails{}
}

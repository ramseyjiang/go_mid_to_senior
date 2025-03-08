package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/labstack/echo"
)

// Company is a structure that contains the company's stock ticker from the client's HTTP request
type Company struct {
	Ticker string `json:"ticker" form:"ticker" query:"ticker"`
}

// PriceResponse is a structure that contains the stock price and other information
type PriceResponse struct {
	Ticker    string    `json:"ticker"`
	Price     string    `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

// GrabPrice - handler method for binding JSON body and scraping for stock price
func GrabPrice(c echo.Context) error {
	com := new(Company)
	if err := c.Bind(com); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	ticker := strings.ToUpper(com.Ticker)
	baseURL := "https://finance.yahoo.com/quote/"
	pricePath := "//fin-streamer[@data-symbol='" + ticker + "' and @data-field='regularMarketPrice']"
	// pricePath := "//fin-streamer[@data-field='regularMarketPrice']" // 更加精确的XPath

	doc, err := htmlquery.LoadURL(baseURL + ticker)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to load URL"})
	}

	// yahoo update html, make script cannot find html tag to get price
	htmlContent := htmlquery.OutputHTML(doc, true)
	fmt.Println(htmlContent)

	context := htmlquery.FindOne(doc, pricePath)
	if context == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Price not found"})
	}

	price := strings.TrimSpace(htmlquery.InnerText(context))

	response := PriceResponse{
		Ticker:    ticker,
		Price:     price,
		Timestamp: time.Now(),
	}

	return c.JSON(http.StatusOK, response)
}

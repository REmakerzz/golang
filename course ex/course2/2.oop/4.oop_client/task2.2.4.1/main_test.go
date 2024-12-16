package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExmo_GetTicker(t *testing.T) {
	mockResponse := `{
		"BTC_USD": {
			"buy": 30000.0,
			"sell": 29950.0,
			"last": 29975.0
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewExmo(WithClient(server.Client()), WithURL(server.URL))

	ticker, err := client.GetTicker()
	assert.NoError(t, err)
	assert.Equal(t, 30000.0, ticker["BTC_USD"].Buy)
}

func TestExmo_GetTrades(t *testing.T) {
	mockResponse := `{
		"BTC_USD": [
			{"price": "30000.0", "quantity": "0.1"}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewExmo(WithClient(server.Client()), WithURL(server.URL))

	trades, err := client.GetTrades("BTC_USD")
	assert.NoError(t, err)
	assert.Equal(t, 30000.0, trades["BTC_USD"][0].Price)
}

func TestExmo_GetOrderBook(t *testing.T) {
	mockResponse := `{
		"BTC_USD": {
			"bid": [["30000.0", "0.5"]],
			"ask": [["30050.0", "0.5"]]
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewExmo(WithClient(server.Client()), WithURL(server.URL))

	orderBook, err := client.GetOrderBook(10, "BTC_USD")
	assert.NoError(t, err)
	assert.Equal(t, 30000.0, orderBook["BTC_USD"].Bid[0].Price)
}

func TestExmo_GetCurrencies(t *testing.T) {
	mockResponse := `["BTC", "USD", "ETH"]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewExmo(WithClient(server.Client()), WithURL(server.URL))

	currencies, err := client.GetCurrencies()
	assert.NoError(t, err)
	assert.Equal(t, []string{"BTC", "USD", "ETH"}, currencies)
}

func TestExmo_GetCandlesHistory(t *testing.T) {
	mockResponse := `{
		"candles": [
			{"o": "30000.0", "c": "30100.0", "h": "30200.0", "l": "29900.0", "t": 1609459200}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewExmo(WithClient(server.Client()), WithURL(server.URL))

	start := time.Unix(1609459200, 0)
	end := time.Unix(1609545600, 0)
	candles, err := client.GetCandlesHistory("BTC_USD", 10, start, end)
	assert.NoError(t, err)
	assert.Equal(t, 30000.0, candles.Candles[0].Open)
}

func TestExmo_GetClosePrice(t *testing.T) {
	mockResponse := `{
		"candles": [
			{"o": "30000.0", "c": "30100.0", "h": "30200.0", "l": "29900.0", "t": 1609459200}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewExmo(WithClient(server.Client()), WithURL(server.URL))

	start := time.Unix(1609459200, 0)
	end := time.Unix(1609545600, 0)
	closePrices, err := client.GetClosePrice("BTC_USD", 10, start, end)
	assert.NoError(t, err)
	assert.Equal(t, 30100.0, closePrices[0])
}

func TestExmo_ErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewExmo(WithClient(server.Client()), WithURL(server.URL))

	_, err := client.GetTicker()
	assert.Error(t, err)

	_, err = client.GetTrades("BTC_USD")
	assert.Error(t, err)

	_, err = client.GetOrderBook(10, "BTC_USD")
	assert.Error(t, err)

	_, err = client.GetCurrencies()
	assert.Error(t, err)

	start := time.Unix(1609459200, 0)
	end := time.Unix(1609545600, 0)
	_, err = client.GetCandlesHistory("BTC_USD", 10, start, end)
	assert.Error(t, err)
}

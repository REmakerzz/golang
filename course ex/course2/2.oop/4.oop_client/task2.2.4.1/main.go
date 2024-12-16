package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Константы для API путей
const (
	tickerPath         = "/ticker"
	tradesPath         = "/trades"
	orderBookPath      = "/order_book"
	currencyPath       = "/currency"
	candlesHistoryPath = "/candles_history"
)

// Структуры ответа API
type TickerValue struct {
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
	Last float64 `json:"last"`
}

type Pair struct {
	Price    float64 `json:"price,string"`
	Quantity float64 `json:"quantity,string"`
}

type OrderBookPair struct {
	Bid [][]float64 `json:"bid"`
	Ask [][]float64 `json:"ask"`
}

type Candle struct {
	Open  float64 `json:"o,string"`
	Close float64 `json:"c,string"`
	High  float64 `json:"h,string"`
	Low   float64 `json:"l,string"`
	Time  int64   `json:"t"`
}

type Exchanger interface {
	GetTicker() (map[string]TickerValue, error)
	GetTrades(pairs ...string) (map[string][]Pair, error)
	GetOrderBook(limit int, pairs ...string) (map[string]OrderBookPair, error)
	GetCurrencies() ([]string, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) ([]Candle, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

// Структура клиента Exmo
type Exmo struct {
	client *http.Client
	url    string
}

// Опции конструктора
func NewExmo(opts ...func(*Exmo)) *Exmo {
	ex := &Exmo{
		client: &http.Client{Timeout: 10 * time.Second},
		url:    "https://api.exmo.com/v1.1",
	}
	for _, opt := range opts {
		opt(ex)
	}
	return ex
}

func WithClient(client *http.Client) func(*Exmo) {
	return func(e *Exmo) {
		e.client = client
	}
}

func WithURL(url string) func(*Exmo) {
	return func(e *Exmo) {
		e.url = url
	}
}

// Вспомогательная функция для выполнения запросов
func (e *Exmo) request(path string, params map[string]string) ([]byte, error) {
	url := e.url + path
	if params != nil {
		query := make([]string, 0)
		for k, v := range params {
			query = append(query, fmt.Sprintf("%s=%s", k, v))
		}
		url += "?" + strings.Join(query, "&")
	}

	resp, err := e.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected server response")
	}

	return ioutil.ReadAll(resp.Body)
}

// Реализация методов Exchanger
func (e *Exmo) GetTicker() (map[string]TickerValue, error) {
	data, err := e.request(tickerPath, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]TickerValue
	err = json.Unmarshal(data, &result)
	return result, err
}

func (e *Exmo) GetTrades(pairs ...string) (map[string][]Pair, error) {
	params := map[string]string{"pair": strings.Join(pairs, ",")}
	data, err := e.request(tradesPath, params)
	if err != nil {
		return nil, err
	}

	var result map[string][]Pair
	err = json.Unmarshal(data, &result)
	return result, err
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (map[string]OrderBookPair, error) {
	params := map[string]string{
		"pair":  strings.Join(pairs, ","),
		"limit": fmt.Sprintf("%d", limit),
	}
	data, err := e.request(orderBookPath, params)
	if err != nil {
		return nil, err
	}

	var result map[string]OrderBookPair
	err = json.Unmarshal(data, &result)
	return result, err
}

func (e *Exmo) GetCurrencies() ([]string, error) {
	data, err := e.request(currencyPath, nil)
	if err != nil {
		return nil, err
	}

	var result []string
	err = json.Unmarshal(data, &result)
	return result, err
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) ([]Candle, error) {
	params := map[string]string{
		"pair":  pair,
		"limit": fmt.Sprintf("%d", limit),
		"from":  fmt.Sprintf("%d", start.Unix()),
		"to":    fmt.Sprintf("%d", end.Unix()),
	}
	data, err := e.request(candlesHistoryPath, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Candles []Candle `json:"candles"`
	}
	err = json.Unmarshal(data, &result)
	return result.Candles, err
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	candles, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		return nil, err
	}

	closePrices := make([]float64, len(candles))
	for i, candle := range candles {
		closePrices[i] = candle.Close
	}
	return closePrices, nil
}

func main() {
	client := NewExmo()
	ticker, err := client.GetTicker()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Ticker:", ticker)
}

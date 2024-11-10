package main

import (
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
	"github.com/guptarohit/asciigraph"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type BinanceTrade struct {
	Price string `json:"price"`
}

var symbols = map[string]struct {
	Name string
	Pair string
}{
	"1": {"BTC_USD", "BTCUSDT"},
	"2": {"LTC_USD", "LTCUSDT"},
	"3": {"ETH_USD", "ETHUSDT"},
}

const (
	binanceAPIURL  = "https://api.binance.com/api/v3/ticker/price?symbol="
	graphWidth     = 100
	graphHeight    = 10
	updateInterval = 1 * time.Second
)

func main() {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	controlChan := make(chan string)
	exitChan := make(chan string)
	go keyboardHandler(controlChan)

	go func() {
		sig := <-sigs
		log.Println("Получен сигнал:", sig)
		exitChan <- "EXIT"
	}()

	state := "menu"
	currentPair := ""
	priceData := make([]float64, 0, graphWidth)
	var mu sync.Mutex

	displayMenu := func() {
		fmt.Fprintf(writer, "\033[H\033[2J")
		output := "Меню:\n"

		keys := make([]string, 0, len(symbols))
		for key := range symbols {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			symbol := symbols[key]
			output += fmt.Sprintf("%s. %s\n", key, symbol.Name)
		}

		output += "\nPress 1-3 to change symbol, press q to exit.\n"
		fmt.Fprintf(writer, output)
	}

	displayMenu()

	ticker := time.NewTicker(updateInterval)
	defer ticker.Stop()

	for {
		select {
		case cmd := <-controlChan:
			if state == "menu" {
				if cmd == "EXIT" || cmd == "q" {
					log.Println("Выход из программы по команде пользователя.")
					return
				}
				if symbol, exists := symbols[cmd]; exists {
					currentPair = symbol.Pair
					state = "monitor"
					priceData = make([]float64, 0, graphWidth)
					fmt.Fprintf(writer, "\033[H\033[2J")
				}
			} else if state == "monitor" {
				if cmd == "BACKSPACE" {
					state = "menu"
					displayMenu()
				} else if cmd == "EXIT" || cmd == "q" {
					log.Println("Выход из программы по команде пользователя.")
					return
				} else if symbol, exists := symbols[cmd]; exists {
					currentPair = symbol.Pair
					priceData = make([]float64, 0, graphWidth)
					fmt.Fprintf(writer, "\033[H\033[2J")
				}
			}
		case cmd := <-exitChan:
			if cmd == "EXIT" {
				log.Println("Выход из программы по системному сигналу.")
				return
			}
		case <-ticker.C:
			if state == "monitor" && currentPair != "" {
				price, err := fetchPrices(currentPair)
				if err != nil {
					log.Println("Ошибка при получении цены:", err)
					continue
				}
				mu.Lock()

				if len(priceData) >= graphWidth {
					priceData = priceData[1:]
				}
				priceData = append(priceData, price)
				graphData := make([]float64, len(priceData))
				copy(graphData, priceData)
				mu.Unlock()

				mu.Lock()
				graph := asciigraph.Plot(
					graphData,
					asciigraph.Height(graphHeight),
					asciigraph.Width(graphWidth),
					asciigraph.SeriesColors(asciigraph.Red),
				)
				currentTime := time.Now().Format("15:04:05")
				currentDate := time.Now().Format("2006-01-02")
				graph = fmt.Sprintf("%v\nТекущая дата: %v\nТекущее время: %v\n", graph, currentDate, currentTime)
				mu.Unlock()

				fmt.Fprintf(writer, "\033[H\033[2J")
				fmt.Fprintf(writer, "%s:%.2f\n\n%s\n\nPress BACKSPACE to return to menu, q to exit.\n", currentPair, price, graph)
			}
		}
	}
}

func fetchPrices(pair string) (float64, error) {
	url := binanceAPIURL + pair
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("HTTP запрос не удался: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("неверный статус ответа: %s", resp.Status)
	}

	var data BinanceTrade
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return 0, fmt.Errorf("Ошибка парсинга JSON:%v", err)
	}
	price, err := strconv.ParseFloat(data.Price, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка конвертации цены для %s: %v", pair, err)
	}
	return price, nil
}

func keyboardHandler(controlChan chan<- string) {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Println("Ошибка при чтении клавиши:", err)
			continue
		}

		if char == 'q' || char == 'Q' {
			controlChan <- "EXIT"
			return
		}

		if key == keyboard.KeyBackspace || key == keyboard.KeyBackspace2 {
			controlChan <- "BACKSPACE"
			continue
		}

		if char >= '1' && char <= '3' {
			controlChan <- string(char)
		}
	}
}

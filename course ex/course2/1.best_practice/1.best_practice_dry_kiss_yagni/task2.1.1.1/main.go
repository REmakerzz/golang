package main

import "fmt"

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

func NewStatisticProfit(opts ...func(*StatisticProfit)) Profitable {
	sp := &StatisticProfit{}
	for _, opt := range opts {
		opt(sp)
	}
	return sp
}

func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}

func (s *StatisticProfit) GetAverageProfit() float64 {
	if s.getAverageProfit != nil {
		return s.getAverageProfit()
	}
	return 0.0
}

func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	if s.getAverageProfitPercent != nil {
		return s.getAverageProfitPercent()
	}
	return 0.0
}

func (s *StatisticProfit) GetCurrentProfit() float64 {
	if s.getCurrentProfit != nil {
		return s.getCurrentProfit()
	}
	return 0.0
}

func (s *StatisticProfit) GetDifferenceProfit() float64 {
	if s.getDifferenceProfit != nil {
		return s.getDifferenceProfit()
	}
	return 0.0
}

func (s *StatisticProfit) GetAllData() []float64 {
	if s.getAllData != nil {
		return s.getAllData()
	}
	return nil
}

func (s *StatisticProfit) Average(prices []float64) float64 {
	if len(prices) == 0 {
		return 0.0
	}
	total := s.Sum(prices)
	return total / float64(len(prices))
}

func (s *StatisticProfit) Sum(prices []float64) float64 {
	total := 0.0
	for _, price := range prices {
		total += price
	}
	return total
}

func WithAverageProfit(s *StatisticProfit) {
	s.getAverageProfit = func() float64 {
		if s.product == nil {
			return 0.0
		}
		return s.Average(s.product.Sells) - s.Average(s.product.Buys)
	}
}

func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = func() float64 {
		if s.product == nil {
			return 0.0
		}
		averageProfit := s.getAverageProfit()
		if len(s.product.Buys) == 0 || s.Average(s.product.Buys) == 0 {
			return 0.0
		}
		return (averageProfit / s.Average(s.product.Buys)) * 100
	}
}

func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = func() float64 {
		if s.product == nil {
			return 0.0
		}
		return s.product.CurrentPrice - (s.product.CurrentPrice * (100 - s.product.ProfitPercent) / 100)
	}
}

func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = func() float64 {
		if s.product == nil {
			return 0.0
		}
		return s.product.CurrentPrice - s.Average(s.product.Sells)
	}
}

func WithAllData(s *StatisticProfit) {
	s.getAllData = func() []float64 {
		res := make([]float64, 0, 4)
		res = append(res, s.getAverageProfit())
		res = append(res, s.getAverageProfitPercent())
		res = append(res, s.getCurrentProfit())
		res = append(res, s.getDifferenceProfit())
		return res
	}
}

func main() {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{1.5, 2.0, 2.5},
		Buys:          []float64{1.0, 1.2, 1.3},
		CurrentPrice:  2.5,
		ProfitPercent: 20,
	}
	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	)

	statProfit.SetProduct(product)

	fmt.Println("Average Profit:", statProfit.GetAverageProfit())
	fmt.Println("Average Profit Percent:", statProfit.GetAverageProfitPercent())
	fmt.Println("Current Profit:", statProfit.GetCurrentProfit())
	fmt.Println("Difference Profit:", statProfit.GetDifferenceProfit())
	fmt.Println("All Data:", statProfit.GetAllData())
}

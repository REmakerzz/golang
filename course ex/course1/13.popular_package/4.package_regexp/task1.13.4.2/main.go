package main

import (
	"fmt"
	"regexp"
)

type Ad struct {
	Title    string
	Описание string
}

func main() {

	ads := []Ad{
		{
			Title:    "Куплю велосипед MeRiDa",
			Описание: "Куплю велосипед merdiDA в хорошем состоянии.",
		},
		{
			Title:    "Продам ВаЗ 2101",
			Описание: "Продам ваз 2101 в хорошем состоянии.",
		},
		{
			Title:    "Продам БМВ",
			Описание: "Продам бМв в хорошем состоянии.",
		},
		{
			Title:    "Продам macBook pro",
			Описание: "Продам macBook PRO в хорошем состоянии.",
		},
	}

	ads = censorAds(ads, map[string]string{
		"велосипед merida": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macbook pro":      "Macbook Pro",
	})

	for _, ad := range ads {
		fmt.Println(ad.Title)
		fmt.Println(ad.Описание)
		fmt.Println()
	}
}

func censorAds(ads []Ad, censor map[string]string) []Ad {

	for i := range ads {
		for pattern, replacement := range censor {
			re := regexp.MustCompile("(?i)" + pattern)
			ads[i].Title = re.ReplaceAllString(ads[i].Title, replacement)
			ads[i].Описание = re.ReplaceAllString(ads[i].Описание, replacement)
		}
	}
	return ads
}

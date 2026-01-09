package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var result []BrainrotMeme
	for _, m := range memes {
		if m.Views > minViews {
			result = append(result, m)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].TrendLevel > result[j].TrendLevel
	})
	return result
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	result := make(map[string]float64)
	for _, m := range memes {
		result[m.Category] += m.Views
	}
	return result
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var result []string
	for _, m := range memes {
		if m.TrendLevel >= 7 || (m.Views > 50 && m.Category == "Sigma") {
			result = append(result, m.Name)
		}
	}
	return result
}

func main() {
	memes := []BrainrotMeme{
		{"Subo Rise", 8, "Subo Bratik", 45},
		{"Night Sigma", 9, "Sigma", 60},
		{"Skibidi Boom", 6, "Skibidi", 80},
		{"Jawline Max", 7, "Mewing", 30},
		{"Tunnel Chant", 5, "TUNTUNTUNSAHUR", 25},
		{"Sigma Walk", 10, "Sigma", 120},
		{"Random Clip", 4, "Other", 15},
	}

	fmt.Println("Топ трендовые мемы:")
	for _, m := range FindTopTrending(memes, 40) {
		fmt.Println(m.Name, m.TrendLevel, m.Views)
	}

	fmt.Println("\nВлияние категорий:")
	for k, v := range CalculateCategoryImpact(memes) {
		fmt.Println(k, v)
	}

	fmt.Println("\nМемы по сложному условию:")
	for _, name := range FilterByComplexCondition(memes) {
		fmt.Println(name)
	}
}
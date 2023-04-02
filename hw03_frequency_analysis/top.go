package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordWithTotal struct {
	word  string
	total int
}

func Top10(text string) []string {
	words := strings.Fields(text)
	wordTotals := make(map[string]int)
	for _, word := range words {
		wordTotals[word]++
	}

	wordWithTotalSlice := []wordWithTotal{}
	for word, total := range wordTotals {
		wordWithTotalSlice = append(wordWithTotalSlice, wordWithTotal{word, total})
	}

	sort.Slice(wordWithTotalSlice, func(i, j int) bool {
		if wordWithTotalSlice[i].total == wordWithTotalSlice[j].total {
			return wordWithTotalSlice[i].word < wordWithTotalSlice[j].word
		}
		return wordWithTotalSlice[i].total > wordWithTotalSlice[j].total
	})

	words = nil
	for i := 0; i < 10 && i < len(wordWithTotalSlice); i++ {
		words = append(words, wordWithTotalSlice[i].word)
	}
	return words
}

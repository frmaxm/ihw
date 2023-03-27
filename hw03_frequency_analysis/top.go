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
	for word, count := range wordTotals {
		wordWithTotalSlice = append(wordWithTotalSlice, wordWithTotal{word, count})
	}

	sort.Slice(wordWithTotalSlice, func(i, j int) bool {
		if wordWithTotalSlice[i].total == wordWithTotalSlice[j].total {
			return wordWithTotalSlice[i].word < wordWithTotalSlice[j].word
		}
		return wordWithTotalSlice[i].total > wordWithTotalSlice[j].total
	})

	var topWords []string
	for i := 0; i < 10 && i < len(wordWithTotalSlice); i++ {
		topWords = append(topWords, wordWithTotalSlice[i].word)
	}
	return topWords
}

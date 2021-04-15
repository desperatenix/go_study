package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

func countWords(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	return wordCounts
}

func Top10(text string) []string {
	if text == "" {
		return []string{}
	}

	words := regexp.MustCompile(`\S+`)
	m := countWords(words.FindAllString(text, -1))

	type kv struct {
		Word  string
		Count int
	}

	var wc []kv //nolint
	for k, v := range m {
		wc = append(wc, kv{k, v})
	}

	sort.Slice(wc, func(i, j int) bool {
		return wc[i].Word < wc[j].Word
	})

	sort.SliceStable(wc, func(i, j int) bool {
		return wc[i].Count > wc[j].Count
	})

	var top10word []string //nolint
	var wordsSlice []kv

	if len(wc) < 10 {
		wordsSlice = wc
	} else {
		wordsSlice = wc[:10]
	}

	for _, kv := range wordsSlice {
		top10word = append(top10word, kv.Word)
	}

	return top10word
}

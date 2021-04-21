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
	uniqWord := countWords(words.FindAllString(text, -1))

	type kv struct {
		Word  string
		Count int
	}

	wc := make([]kv, len(uniqWord))
	i := 0
	for k, v := range uniqWord {
		wc[i] = kv{k, v}
		i++
	}

	sort.Slice(wc, func(i, j int) bool {
		if wc[i].Count == wc[j].Count {
			return wc[i].Word < wc[j].Word
		}
		return wc[i].Count > wc[j].Count
	})

	if len(wc) > 10 {
		wc = wc[:10]
	}

	top10word := make([]string, len(wc))
	for i, kv := range wc {
		top10word[i] = kv.Word
	}

	return top10word
}

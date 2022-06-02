package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {

	fileName := "file.txt"
	totalCounts := WordCount(fileName)
	fmt.Println(totalCounts)

}

func WordCount(fileName string) string {

	file, err := ioutil.ReadFile(fileName)

	if err != nil {

		log.Fatal(err)
	}
	// convert byteslice to string
	text := string(file)
	words := strings.Fields(text)

	// count same words
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	// create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]string, len(m))
	for key := range m {
		wordCounts = append(wordCounts, key)
	}

	// sort wordCount slice
	sort.Slice(wordCounts, func(i, j int) bool {
		return m[wordCounts[i]] > m[wordCounts[j]]
	})

	// get the ten most frequent words
	n := make(map[string]int)
	for index, key := range wordCounts {
		n[key] = m[key]
		fmt.Printf("%s %d\n", key, n[key])
		if index == 9 {
			break
		}
	}
	return "Running"

}

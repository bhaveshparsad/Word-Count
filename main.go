package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// WordCount holds word and count pair
type WordCount struct {
	word  string
	count int
}

func main() {

    fileName := "file.txt"

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
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	// create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{word: key, count: val})
	}

	// sort wordCount slice 
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	// get the ten most frequent words
	for i := 0; i < len(wordCounts) && i < 10; i++ {
		fmt.Println(wordCounts[i].word, ":", wordCounts[i].count)
	}
}
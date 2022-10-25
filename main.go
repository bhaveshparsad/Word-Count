package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Function for word count, returns map
func word(s string) string {
	count := map[string]int{}
	for _, word := range strings.Fields(s) { //Fields for words as a slice {
		count[word]++
	}

	word_count := make([]count_word, 0, len(count)) // slice
	for k, v := range count {
		word_count = append(word_count, count_word{word: k, count: v}) //counting frequency of words
	}

	//sorting (high to low frequency of words)
	sort.Slice(word_count, func(i, j int) bool {
		return word_count[i].count > word_count[j].count
	})

	for i := 0; i < 10 && i != len(word_count); i++ {
		fmt.Println("Word ", word_count[i].word, "\t count :  \t", word_count[i].count)

	}
	return "Success"
}

type count_word struct {
	word  string
	count int
}

func main() {
	// for input
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal()
	}

	res := word(str)

	fmt.Println(res)
}


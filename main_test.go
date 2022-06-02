package main

import "testing"

func TestWordCount(t *testing.T){
	input := "file.txt"
	Output := "Running"

	x := WordCount(input)

	if x!= Output {
		t.Errorf("got %q, expected %q", x,Output)
	}
}


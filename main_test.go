package main

import "testing"

func TestWordCount(t *testing.T) {
	File := "file.txt"
	Output := "Running"

	x := WordCount(File)

	if x != Output {
		t.Errorf("got %q, expected %q", x, Output)
	}
}

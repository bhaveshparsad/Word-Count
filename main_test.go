package main

import (
	"testing"
)

func TestWord(t *testing.T) {
	str := "Think Future Technoogies Pvt. Ltd."
	expected := "Success"

	actual := word(str)

	if actual != expected {
		t.Errorf("Failed...")
	}
}
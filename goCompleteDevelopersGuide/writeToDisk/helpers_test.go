package main

import "testing"

func TestBytesToStringList(t *testing.T) {
	b := []byte{99, 97, 114, 100, 49, 44, 99, 97, 114, 100, 50}
	s := BytesToStringList(b)
	// Create length based test
	if len(s) != 2 {
		t.Errorf("Expected length 2, but got: %v", len(s))
	}
	// Since go can't compare slices directly run iterative list comparison
	comp_str := []string{"card1", "card2"}
	for i := range comp_str {
		if comp_str[i] != s[i] {
			t.Errorf("Expeced element %v to be %v but got %v instead", i, comp_str[i], s[i])
		}
	}
}

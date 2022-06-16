package main

import "testing"

// a1 := []string{"apple", "pear", "lemon", "orange"}

func TestTakeEasy(t *testing.T) {
	inp1 := []string{"apple", "pear", "lemon", "orange"}
	n1 := 2
	exp1 := []string{"apple", "pear"}
	out1 := takeEasy(inp1, n1)
	for i, j := range out1 {
		if j != exp1[i] {
			t.Errorf("Expected %v, got %v", exp1[i], j)
		}
	}

	// test n > len array
	n2 := 5
	out2 := takeEasy(inp1, n2)
	for i, j := range out2 {
		if j != inp1[i] {
			t.Errorf("Expected %v, got %v", exp1[i], j)
		}
	}
}

func TestMed1(t *testing.T) {
	inp1 := []string{"apple", "pear", "lemon", "orange"}
	n1 := 2
	exp1 := []string{"apple", "pear"}
	out1 := takeMed1(inp1, n1)
	for i, val := range exp1 {
		if val != exp1[i] {
			t.Errorf("Expected %v, got %v", val, out1[i])
		}
	}

	n2 := 8
	if n2 != len(takeMed1(inp1, n2)) {
		t.Errorf("Expected array length %v, got array length %v", n2, len(takeMed1(inp1, n2)))
	}
}

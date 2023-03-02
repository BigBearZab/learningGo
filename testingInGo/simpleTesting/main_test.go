package main

import "testing"

func Test_IsPrime(t *testing.T) {
	type test struct {
		value           int
		expectedBoolean bool
	}

	tests := map[string]test{
		"Basic prime": {
			value:           2,
			expectedBoolean: true,
		},
		"Bigger prime": {
			value:           37,
			expectedBoolean: true,
		},
		"Zero expects not prime": {
			value:           0,
			expectedBoolean: false,
		},
		"Negative should not be prime": {
			value:           -10,
			expectedBoolean: false,
		},
	}

	for name, tt := range tests {
		testBool, _ := isPrime(tt.value)
		if testBool != tt.expectedBoolean {
			t.Errorf("%v: Expected prime check for %d to be %v, but got %v", name, tt.value, tt.expectedBoolean, testBool)
		}
	}
}

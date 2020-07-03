package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: 2, b: 2, want: 4},
		{a: 6, b: 7, want: 13},
		{a: 5, b: 10, want: 15},
		{a: 7, b: 13, want: 20},
		{a: 15, b: 15, want: 30},
		{a: 15, b: 25, want: 40},
		{a: 3, b: 11, want: 14},

	}
	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}


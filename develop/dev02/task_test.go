package main

import "testing"

func TestAddSymbolToStr(t *testing.T) {

	testTable := []struct {
		symbol      string
		countSymbol string
		expected    string
	}{
		{
			symbol:      "s",
			countSymbol: "4",
			expected:    "ssss",
		},
		{
			symbol:      "",
			countSymbol: "4",
			expected:    "",
		},
		{
			symbol:      "a",
			countSymbol: "1",
			expected:    "a",
		},
		{
			symbol:      "a",
			countSymbol: "-1",
			expected:    "",
		},
		{
			symbol:      "a",
			countSymbol: "0",
			expected:    "",
		},
		{
			symbol:      "aa",
			countSymbol: "2",
			expected:    "aaaa",
		},
	}

	for _, testCase := range testTable {
		result := addSymbolToStr(testCase.symbol, testCase.countSymbol)
		t.Logf("result: %s, expect: %s",
			result, testCase.expected)
		if result != testCase.expected {
			t.Errorf("Expect: %s, got: %s",
				testCase.expected, result)
		}
	}
}

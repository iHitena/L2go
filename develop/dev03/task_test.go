package main

import "testing"
import "reflect"

func TestSpecifyingColumnSort(t *testing.T) {
	testTable := []struct {
		inSliceString []string
		column        string
		normalSort    bool

		expected []string
	}{
		{
			inSliceString: []string{"b a", "c"},
			column:        "",
			normalSort:    true,

			expected: []string{"a b", "c"},
		},

		{
			inSliceString: []string{"b a a c", "a"},
			column:        "b",
			normalSort:    true,

			expected: []string{"b b a a c", "a"},
		},

		{
			inSliceString: []string{"b a a c", "a"},
			column:        "",
			normalSort:    false,

			expected: []string{"c b a a", "a"},
		},

		{
			inSliceString: []string{"b a a c", "b a a"},
			column:        "b",
			normalSort:    false,

			expected: []string{"b b c a a", "b b a a"},
		},
	}

	for _, testCase := range testTable {
		result := specifyingColumnSort(testCase.inSliceString, testCase.column, testCase.normalSort)
		t.Logf("result: %s, expect: %s",
			result, testCase.expected)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expect: %s, got: %s",
				testCase.expected, result)
		}
	}
}

package main

import "testing"
import "reflect"

func TestFindAnagrams(t *testing.T) {
	testTable := []struct {
		arrayString []string

		expected *map[string]*[]string
	}{
		{
			arrayString: []string{"apa", "aap", "paa", "d"},

			expected: &map[string]*[]string{
				"aap": {"aap", "apa", "paa"},
			},
		},

		{
			arrayString: []string{},

			expected: &map[string]*[]string{},
		},

		{
			arrayString: []string{"d"},

			expected: &map[string]*[]string{},
		},

		{
			arrayString: []string{"apa", "da"},

			expected: &map[string]*[]string{
				"aap": {"apa"}, "ad": {"da"},
			},
		},
	}

	for _, testCase := range testTable {
		result := FindAnagrams(&testCase.arrayString)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Error(testCase.expected, result)
		}
	}
}

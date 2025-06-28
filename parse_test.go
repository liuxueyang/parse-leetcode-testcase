package main

import (
	"testing"
)

func TestProcessLine(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{`nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2`, []string{"[1,1,1,0,0,0,1,1,1,1,0]", "2"}},
		{`s = "A man, a plan, a canal: Panama", k = 1, b = "Hello, world"`, []string{"A man, a plan, a canal: Panama", "1", "Hello, world"}},
		{`s = "A man, a plan, a canal: Panama", k = 1, b = "Hello, world"@leetcode`, []string{"A man, a plan, a canal: Panama", "1", "Hello, world"}},
		{`n = 4, k = 6`, []string{"4", "6"}},
		{`nums = [6,2,8,4]`, []string{"[6,2,8,4]"}},
		{`nums = [1,5,1,4,2]`, []string{"[1,5,1,4,2]"}},
		{`original = [1,2,3,4], bounds = [[1,2],[2,3],[3,4],[4,5]]`, []string{"[1,2,3,4]", "[[1,2],[2,3],[3,4],[4,5]]"}},
		{`nums = [2,7,11,15], target = 9`, []string{"[2,7,11,15]", "9"}},
		{`l1 = [2,4,3], l2 = [5,6,4]`, []string{"[2,4,3]", "[5,6,4]"}},
		{`s = "abcabcbb"`, []string{"abcabcbb"}},
		{`nums1 = [1,3], nums2 = [2]`, []string{"[1,3]", "[2]"}},
		{`s = "PAYPALISHIRING", numRows = 3`, []string{"PAYPALISHIRING", "3"}},
		{`x = -123`, []string{"-123"}},
		{`x = 0`, []string{"0"}},
		{`s = "42"`, []string{"42"}},
		{`s = "aa", p = "a*"`, []string{"aa", "a*"}},
		{`nums = [-1,0,1,2,-1,-4]`, []string{"[-1,0,1,2,-1,-4]"}},
		{`s = "()[]{}"`, []string{"()[]{}"}},
		{`nums =  [3,3,7,7,10,11,11]`, []string{"[3,3,7,7,10,11,11]"}},
	}
	for _, test := range tests {
		res := processLine(test.input)
		if len(res) != len(test.expected) {
			t.Errorf("processLine(%q) = %v; want %v", test.input, res, test.expected)
		}
		for i, v := range res {
			if v != test.expected[i] {
				t.Errorf("processLine(%q)[%d] = %q; want %q", test.input, i, v, test.expected[i])
			}
		}
	}
}

func TestConvertTo1DStringSlice(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"[\"Hello\",\"World\"]", []string{"Hello", "World"}},
		{"[1,1,1,0,0,0,1,1,1,1,0]", []string{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "0"}},
	}

	for _, test := range tests {
		result := convertTo1DStringSlice(test.input)
		if !equalStringSlices(result, test.expected) {
			t.Errorf("convertToString1DSlice(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestConvertTo2DStringSlice(t *testing.T) {
	tests := []struct {
		input    string
		expected [][]string
	}{
		{"[[\"Hello\",\"World\"],[\"Foo\",\"Bar\"]]", [][]string{{"Hello", "World"}, {"Foo", "Bar"}}},
		{"[[1,1,1,0,0,0,1,1,1,1,0]]", [][]string{{"1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "0"}}},
		{`[[-5]]`, [][]string{{"-5"}}},
	}

	for _, test := range tests {
		result := convertTo2DStringSlice(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("convertToString2DSlice(%q) length = %d; want %d", test.input, len(result), len(test.expected))
			continue
		}
		for i := range result {
			if !equalStringSlices(result[i], test.expected[i]) {
				t.Errorf("convertToString2DSlice(%q)[%d] = %v; want %v", test.input, i, result[i], test.expected[i])
			}
		}
	}
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

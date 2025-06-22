package main

import (
	"fmt"
	"testing"
)

func TestProcessLine(t *testing.T) {
	line := `s = "A man, a plan, a canal: Panama", k = 1, b = "Hello, world"`
	t.Log(line)
	res := processLine(line)
	t.Log(res)
	fmt.Printf("Hello")

	tests := []struct {
		input    string
		expected []string
	}{
		{`nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2`, []string{"[1,1,1,0,0,0,1,1,1,1,0]", "2"}},
		{`s = "A man, a plan, a canal: Panama", k = 1, b = "Hello, world"`, []string{"\"A man, a plan, a canal: Panama\"", "1", "\"Hello, world\""}},
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

		t.Logf("Input: %s, Processed: %v", test.input, res)
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

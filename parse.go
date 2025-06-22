package main

import (
	"strings"
)

func processLine(s string) []string {
	parts := strings.Split(s, " = ")

	var ans []string

	for i, part := range parts {
		if i == 0 {
			continue
		}
		part = removePostfix(part)
		ans = append(ans, part)
	}

	return ans
}

func removePostfix(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '"' {
			return s[:i+1]
		}
		if s[i] == ',' {
			return s[:i]
		}
	}
	return s
}

func convertTo1DStringSlice(s string) []string {
	s = strings.Trim(s, "[]")
	parts := strings.Split(s, ",")
	result := make([]string, len(parts))

	for i, part := range parts {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "\"") && strings.HasSuffix(part, "\"") {
			part = strings.Trim(part, "\"")
		}
		result[i] = part
	}

	return result
}

func convertTo2DStringSlice(s string) [][]string {
	s = strings.Trim(s, "[]")
	innerSlices := strings.Split(s, "],[")

	result := make([][]string, len(innerSlices))

	for i, inner := range innerSlices {
		result[i] = convertTo1DStringSlice(inner)
	}

	return result
}

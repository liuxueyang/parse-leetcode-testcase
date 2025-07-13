package main

import (
	"strings"
)

func doGetTokens(s string) (ans []string) {
	s = strings.TrimSuffix(s, "©leetcode")
	parts := strings.Split(s, " = ")

	if len(parts) == 1 {
		// If there is no '=' in the string, return the string as a single element
		s = strings.TrimSuffix(parts[0], "©leetcode")
		ans = append(ans, unquoteString(s))
		return ans
	}

	for i, part := range parts {
		if i == 0 {
			continue
		}
		part = unquoteString(removePostfix(part))
		part = strings.TrimSpace(part)
		ans = append(ans, part)
	}

	return ans
}

// remove all of the characters after the last ','
// If it is a string with '"', then return the string up to the last '"'
// If it is a slice with ']', then return the string up to the last ']'
func removePostfix(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '"' || s[i] == ']' {
			return s[:i+1]
		}
		if s[i] == ',' {
			return s[:i]
		}
	}
	return s
}

func rawStrTo1DStrSlice(s string) []string {
	s = strings.Trim(s, "[]")
	parts := strings.Split(s, ",")
	result := make([]string, len(parts))

	for i, part := range parts {
		part = unquoteString(strings.TrimSpace(part))
		result[i] = part
	}

	return result
}

func unquoteString(s string) string {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return strings.Trim(s, "\"")
	}
	return s
}

func rawStrTo2DStrSlice(s string) [][]string {
	s = strings.Trim(s, "[]")
	innerSlices := strings.Split(s, "],[")

	result := make([][]string, len(innerSlices))

	for i, inner := range innerSlices {
		result[i] = rawStrTo1DStrSlice(inner)
	}

	return result
}

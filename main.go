package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var name = flag.String("p", "", "The suffix name of the file to write to")

func main() {
	flag.Parse()

	readFh, err := os.Open("raw.txt")
	if err != nil {
		panic(err)
	}
	defer readFh.Close()

	// write the output to file "input.txt"
	writeFh, err := os.OpenFile("input.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer writeFh.Close()

	// Create a new writer
	writer := bufio.NewWriter(writeFh)

	input := bufio.NewScanner(readFh)
	for input.Scan() {
		line := input.Text()

		if strings.Contains(line, "输入：") || strings.Contains(line, "输入:") {
			ProcessInput(line, writer)
		}
	}

	writer.Flush() // Ensure all data is written to the file
	writeFh.Sync() // Ensure the file is synced to disk

	if *name != "" {
		// If a suffix name is provided, copy the "input.txt" to "input_<suffix>.txt"
		// Create a new file with the suffix name
		newName := fmt.Sprintf("input_%s.txt", *name)
		newFile, err := os.OpenFile(newName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		defer newFile.Close()

		newWriter := bufio.NewWriter(newFile)

		readFh, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer readFh.Close()

		scanner := bufio.NewScanner(readFh)
		for scanner.Scan() {
			newWriter.WriteString(scanner.Text() + "\n")
		}
		newWriter.Flush() // Ensure all data is written to the new file
		// Now we have created a new file with the suffix name
		// This will copy the content of "input.txt" to "input_<suffix>.txt"
		// Note: This is not a true copy operation, it just writes the content to a new file
		// copy the "input.txt" to "input_<suffix>.txt"
	}
}

func ProcessInput(line string, writer *bufio.Writer) {
	// Remove the prefix "输入：" or "输入:"
	line = strings.TrimPrefix(line, "输入：")
	line = strings.TrimPrefix(line, "输入:")
	// Remove any leading or trailing whitespace
	line = strings.TrimSpace(line)
	// Print the cleaned line
	// If the line is not empty, print it
	if line != "" {
		// split the line by ", "
		parts := strings.Split(line, ", ")
		for _, part := range parts {
			// Remove any leading or trailing whitespace from each part
			part = strings.TrimSpace(part)
			// Print the cleaned part
			if part != "" {
				parts := strings.Split(part, " = ")
				line := parts[1]

				if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
					// Check if it is a 2-dimensional slice
					if is2dSlice(line) {
						// Convert to a 2-dimensional slice
						slice2d := convertToSlice2d(line)
						fmt.Fprintf(writer, "%d ", len(slice2d))
						if len(slice2d) > 0 {
							fmt.Fprintf(writer, "%d\n", len(slice2d[0]))
							for _, innerSlice := range slice2d {
								for _, v := range innerSlice {
									fmt.Fprintf(writer, "%d ", v)
								}
								fmt.Fprintf(writer, "\n")
							}
						}
					} else {
						// Convert to a 1-dimensional slice
						slice1d := convertToIntSlice1d(line)
						fmt.Fprintf(writer, "%d\n", len(slice1d))
						for _, v := range slice1d {
							fmt.Fprintf(writer, "%d ", v)
						}
						fmt.Fprintf(writer, "\n")
					}
				} else {
					// If it is not a slice, just print the part
					if strings.HasPrefix(line, "\"") || strings.HasSuffix(line, "'") {
						// If it is a string, remove the quotes
						line = strings.Trim(line, "\"'")
					}
					fmt.Fprintf(writer, "%s\n", line)
				}
			}
		}
	}
}

// Given a string multi-dimensional slice
// Given a string representation of a multi-dimensional slice, convert it to an actual multi-dimensional slice of integers
// Example: convertToSlice("[[1,2,3],[4,5,6]]") should return [][]int{{1, 2, 3}, {4, 5, 6}}
func convertToSlice2d(s string) [][]int {
	// Remove the outer brackets
	s = strings.Trim(s, "[]")
	// Split by "],[" to get each inner slice
	innerSlices := strings.Split(s, "],[")
	result := make([][]int, len(innerSlices))

	for i, inner := range innerSlices {
		// Remove any leading or trailing whitespace
		inner = strings.TrimSpace(inner)
		// Split by "," to get individual integers
		nums := strings.Split(inner, ",")
		result[i] = make([]int, len(nums))
		for j, num := range nums {
			fmt.Sscanf(num, "%d", &result[i][j])
		}
	}

	return result
}

// Given a string representation of a slice, convert it to an actual slice of integers
func convertToIntSlice1d(s string) []int {
	// Remove the brackets
	s = strings.Trim(s, "[]")
	// Split by "," to get individual integers
	nums := strings.Split(s, ",")
	result := make([]int, len(nums))

	for i, num := range nums {
		// Remove any leading or trailing whitespace
		num = strings.TrimSpace(num)
		fmt.Sscanf(num, "%d", &result[i])
	}

	return result
}

// Given a string. Check Whether it is 1-dimensional or 2-dimensional slice
func is2dSlice(s string) bool {
	// Check if the string contains multiple inner slices
	return strings.Contains(s, "],[") || strings.Contains(s, "], [")
}

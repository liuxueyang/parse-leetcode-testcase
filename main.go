package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var name = flag.String("p", "", "The suffix name of the file to write to")
var inputFile = flag.String("i", "raw.txt", "The input file to read from")

func main() {
	flag.Parse()

	readFh, err := os.Open(*inputFile)
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

		if strings.Contains(line, "输入：") || strings.Contains(line, "输入:") || strings.Contains(line, "Input:") || strings.Contains(line, "Input: ") {
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
	}
}

// Trim prefix "输入:" or "输入：" or "Input:" or "Input: "
func trimPrefixInput(s string) string {
	s = strings.TrimPrefix(s, "输入：")
	s = strings.TrimPrefix(s, "输入:")
	s = strings.TrimPrefix(s, "Input: ")
	s = strings.TrimPrefix(s, "Input:")
	return strings.TrimSpace(s)
}

func ProcessInput(line string, writer *bufio.Writer) {
	line = trimPrefixInput(line)
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return
	}

	parts := processLine(line)
	for _, line := range parts {
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if is2dSlice(line) {
				slice2d := convertTo2DStringSlice(line)
				write2DSlice(writer, slice2d)
			} else {
				slice1d := convertTo1DStringSlice(line)
				writeSlice(writer, slice1d)
			}
		} else {
			fmt.Fprintf(writer, "%s\n", line)
		}
	}
}

func write2DSlice(writer *bufio.Writer, slice [][]string) {
	fmt.Fprintf(writer, "%d ", len(slice))
	fmt.Fprintf(writer, "%d\n", len(slice[0]))

	for _, innerSlice := range slice {
		for _, v := range innerSlice {
			fmt.Fprintf(writer, "%s ", v)
		}
		fmt.Fprintf(writer, "\n")
	}
}

func writeSlice(writer *bufio.Writer, slice []string) {
	fmt.Fprintf(writer, "%d\n", len(slice))
	for _, v := range slice {
		fmt.Fprintf(writer, "%s ", v)
	}
	fmt.Fprintf(writer, "\n")
}

// Given a string. Check Whether it is 1-dimensional or 2-dimensional slice
func is2dSlice(s string) bool {
	// Check if the string contains multiple inner slices
	return strings.Contains(s, "],[") || strings.Contains(s, "], [") || (strings.HasPrefix(s, "[[") && strings.HasSuffix(s, "]]"))
}

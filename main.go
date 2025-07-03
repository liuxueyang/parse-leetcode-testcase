package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	DefaultOutputFile = "input.txt"
	DefaultInputFile  = "raw.txt"
	FilePermission    = 0644
)

var name = flag.String("p", "", "The suffix name of the file to write to")
var inputFile = flag.String("i", "raw.txt", "The input file to read from")
var verbose = flag.Bool("v", false, "Enable verbose output")

var inputPrefixes = []string{"输入：", "输入:", "Input: ", "Input:"}

func main() {
	flag.Parse()

	err := validateFlags()
	if err != nil {
		log.Fatalf("Error validating flags: %v\n", err)
	}

	readFh, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer readFh.Close()

	writeFh, err := os.OpenFile(DefaultOutputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, FilePermission)
	if err != nil {
		log.Fatalf("Error opening output file: %v\n", err)
	}
	defer writeFh.Close()

	// Create a new writer
	writer := bufio.NewWriter(writeFh)

	input := bufio.NewScanner(readFh)
	var processNext bool

	for input.Scan() {
		line := input.Text()

		if containsAny(line, inputPrefixes) || processNext {
			ProcessInput(line, writer)
			processNext = false // Reset the flag after processing input
		} else if strings.HasSuffix(line, "=") {
			// If the line ends with '=', it indicates the start of a new input section
			processNext = true
		}
	}

	writer.Flush() // Ensure all data is written to the file
	writeFh.Sync() // Ensure the file is synced to disk

	if *name != "" {
		logf("Backing up file with suffix: %s\n", *name)
		err = backupFile(*name)
		if err != nil {
			log.Fatalf("Error backing up file: %v\n", err)
		}
	}
}

func logf(format string, args ...interface{}) {
	if *verbose {
		log.Printf(format, args...)
	}
}

func validateFlags() error {
	if *inputFile == "" {
		return fmt.Errorf("input file must be specified")
	}
	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", *inputFile)
	}
	return nil
}

func containsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

func backupFile(fileName string) error {
	if len(fileName) == 0 {
		return nil
	}

	// If a suffix name is provided, copy the output file to "input_<suffix>.txt"
	newName := fmt.Sprintf("input_%s.txt", fileName)
	newFile, err := os.Create(newName)
	if err != nil {
		return err
	}
	defer newFile.Close()

	readFh, err := os.Open(DefaultOutputFile)
	if err != nil {
		return err
	}
	defer readFh.Close()

	_, err = io.Copy(bufio.NewWriter(newFile), readFh)
	if err != nil {
		return err
	}
	return err
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

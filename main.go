package main

import (
	"bufio"
	"encoding/json"
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
var companion = flag.Bool("companion", false, "Export to competitive companion format")
var outputFile = flag.String("o", DefaultOutputFile, "The output file to write to")

var inputPrefixes = []string{"输入：", "输入:", "Input: ", "Input:"}

type TestCase struct {
	Test string `json:"test"`
}

type TestCases struct {
	TestCases []TestCase `json:"testcases"`
}

func main() {
	flag.Parse()

	err := validateFlags()
	if err != nil {
		log.Fatalf("Error validating flags: %v\n", err)
	}

	err = processFile(*inputFile, *outputFile)
	if err != nil {
		log.Fatalf("Error processing file: %v\n", err)
	}

	if *name != "" {
		logf("Backing up file with suffix: %s\n", *name)
		err = backupFile(*name)
		if err != nil {
			log.Fatalf("Error backing up file: %v\n", err)
		}
	}
}

func processFile(inputPath, outputPath string) error {
	readFh, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer readFh.Close()

	writeFh, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, FilePermission)
	if err != nil {
		log.Fatalf("Error opening output file: %v\n", err)
	}
	defer writeFh.Close()

	return processLines(readFh, writeFh)
}

func processLines(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)

	bufWriter := bufio.NewWriter(writer)
	defer bufWriter.Flush()

	var processNext bool
	var testCases []TestCase

	for scanner.Scan() {
		line := scanner.Text()

		// TODO: bug here.
		if containsAny(line, inputPrefixes) || processNext {
			tmp := processRawLine(line)

			if *companion {
				testCases = append(testCases, TestCase{Test: tmp})
			} else {
				writer.Write([]byte(tmp))
			}

			processNext = false // Reset the flag after processing input
		} else if strings.HasSuffix(line, "=") {
			// If the line ends with '=', it indicates the start of a new input section
			processNext = true
		}
	}

	if *companion && len(testCases) > 0 {
		// 将每个 TestCase 单独输出为 JSON 对象
		fmt.Fprint(writer, "[\n\t")
		for i, tc := range testCases {
			jsonData, err := json.MarshalIndent(tc, "\t", "\t")
			if err != nil {
				return err
			}

			if i > 0 {
				fmt.Fprint(writer, ",\n\t")
			}
			writer.Write(jsonData)
		}
		fmt.Fprint(writer, "\n]\n")
	}

	return scanner.Err()
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

var trimPrefixes = []string{"输入:", "输入：", "Input:", "Input: ", "input:", "input: "}

func trimPrefixInput(s string) string {
	for _, prefix := range trimPrefixes {
		s = strings.TrimPrefix(s, prefix)
	}
	return strings.TrimSpace(s)
}

func processRawLine(line string) string {
	parts := getTokens(line)
	var sb strings.Builder

	for _, line := range parts {
		if is2DSlice(line) {
			s := rawStrTo2DStrSlice(line)
			tmp := twoDimSliceToStr(s)
			sb.WriteString(tmp)
		} else if is1DSlice(line) {
			s := rawStrTo1DStrSlice(line)
			tmp := oneDimSliceToStr(s)
			sb.WriteString(tmp)
		} else {
			sb.WriteString(line + "\n")
		}
	}

	return sb.String()
}

func is1DSlice(s string) bool {
	return strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") && !is2DSlice(s)
}

func getTokens(line string) []string {
	line = trimPrefixInput(line)
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return nil
	}

	return doGetTokens(line)
}

func twoDimSliceToStr(slice [][]string) string {
	var sb strings.Builder
	n := len(slice)
	m := len(slice[0])

	sb.WriteString(fmt.Sprintf("%d %d\n", n, m))

	for _, innerSlice := range slice {
		for i, v := range innerSlice {
			sb.WriteString(v)
			if i < m-1 {
				sb.WriteByte(' ')
			} else {
				sb.WriteByte('\n')
			}
		}
	}

	return sb.String()
}

func oneDimSliceToStr(slice []string) string {
	var sb strings.Builder

	n := len(slice)

	sb.WriteString(fmt.Sprintf("%d\n", n))

	for i, s := range slice {
		sb.WriteString(s)
		if i < n-1 {
			sb.WriteByte(' ')
		} else {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

// Given a string. Check Whether it is 1-dimensional or 2-dimensional slice
func is2DSlice(s string) bool {
	// Check if the string contains multiple inner slices
	return strings.Contains(s, "],[") || strings.Contains(s, "], [") || (strings.HasPrefix(s, "[[") && strings.HasSuffix(s, "]]"))
}

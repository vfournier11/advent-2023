package cmd

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var wordAreDigits bool
var filename string

const defaultFilename = "day_01.txt"

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use: "day1",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(filename) // data is a byte slice
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(data), "\n")
		if lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1] // remove last empty line
		}
		calibrationFunction := findCalibrationValue
		if wordAreDigits {
			calibrationFunction = findCalibrationValueWordOrDigit
		}
		calibrationValueSum, err := findCalibrationDocumentValue(lines, calibrationFunction)
		if err != nil {
			panic(err)
		}

		fmt.Println(calibrationValueSum)
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)

	day1Cmd.Flags().BoolVarP(&wordAreDigits, "word-are-digits", "w", false, "Treat words as digits (second star)")
	day1Cmd.Flags().StringVarP(&filename, "filename", "f", defaultFilename, "Filename to read from")
}

type CalibrationFunc func(string) (int, error)

// findCalibrationDocumentValue Returns the sum of all calibration values in the document.
func findCalibrationDocumentValue(lines []string, calibrationFunc CalibrationFunc) (int, error) {
	var calibrationValueSum int
	for _, line := range lines {
		currentCalibrationValue, err := calibrationFunc(line)
		if err != nil {
			return -1, err
		}
		calibrationValueSum += currentCalibrationValue
	}
	return calibrationValueSum, nil
}

// findCalibrationValue Returns the combination of the first and last digit of the line.
func findCalibrationValue(line string) (int, error) {
	var leftDigit, rightDigit int
	leftDigit = -1
	for _, c := range line {
		if unicode.IsDigit(c) {
			leftDigit = int(c - '0')
			break
		}
	}
	if leftDigit == -1 {
		return -1, fmt.Errorf("no digit found in line: %s", line)
	}

	rightDigit = -1
	for i := len(line) - 1; i >= 0; i-- {
		c := rune(line[i])
		if unicode.IsDigit(c) {
			rightDigit = int(c - '0')
			break
		}
	}
	if rightDigit == -1 {
		return -1, fmt.Errorf("no digit found in line: %s", line)
	}

	return leftDigit*10 + rightDigit, nil
}

// findCalibrationValueWordOrDigit Returns the combination of the first and last digit of the line. The digit can be expressed as a word (e.g.: one, two).
func findCalibrationValueWordOrDigit(line string) (int, error) {
	var leftDigit, rightDigit int
	leftDigit = -1
	for idx, c := range line {
		if unicode.IsDigit(c) {
			leftDigit = int(c - '0')
			break
		}
		if digit := digitFromStringIndex(line, idx); digit != -1 {
			leftDigit = digit
			break
		}
	}
	if leftDigit == -1 {
		return -1, fmt.Errorf("no digit found in line: %s", line)
	}

	rightDigit = -1
	for i := len(line) - 1; i >= 0; i-- {
		c := rune(line[i])
		if unicode.IsDigit(c) {
			rightDigit = int(c - '0')
			break
		}
		if digit := digitFromStringIndex(line, i); digit != -1 {
			rightDigit = digit
			break
		}
	}
	if rightDigit == -1 {
		return -1, fmt.Errorf("no digit found in line: %s", line)
	}

	return leftDigit*10 + rightDigit, nil
}

func digitFromStringIndex(line string, idx int) int {
	digitInWords := []struct {
		Word  string
		Digit int
	}{
		{"zero", 0},
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}
	for _, digitInWord := range digitInWords {
		if idx+len(digitInWord.Word) > len(line) {
			continue
		}
		if line[idx:idx+len(digitInWord.Word)] == digitInWord.Word {
			return digitInWord.Digit
		}
	}
	return -1
}

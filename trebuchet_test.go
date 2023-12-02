package main

import "testing"

func TestFindCalibrationValue(t *testing.T) {
	for _, test := range []struct {
		line string
		want int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"treb7uchet", 77},
	} {
		got, err := findCalibrationValue(test.line)
		if err != nil {
			t.Fatalf("findCalibrationValue(%s) returned error: %v", test.line, err)
		}
		if got != test.want {
			t.Errorf("findCalibrationValue(%s) = %d, want %d", test.line, got, test.want)
		}
	}
}

func TestFindCalibrationDocumentValue(t *testing.T) {
	for _, test := range []struct {
		lines []string
		want  int
	}{
		{[]string{"1abc2", "pqr3stu8vwx", "treb7uchet"}, 127},
	} {
		got, err := findCalibrationDocumentValue(test.lines, findCalibrationValue)
		if err != nil {
			t.Fatalf("findCalibrationDocumentValue(%v) returned error: %v", test.lines, err)
		}
		if got != test.want {
			t.Errorf("findCalibrationDocumentValue(%v) = %d, want %d", test.lines, got, test.want)
		}
	}
}

func TestDigitFromStringIndex(t *testing.T) {
	for _, test := range []struct {
		line string
		idx  int
		want int
	}{
		{"1abc2", 0, -1},
		{"pqr3s", 0, -1},
		{"onet7two", 0, 1},
		{"otwo", 1, 2},
	} {
		got := digitFromStringIndex(test.line, test.idx)
		if got != test.want {
			t.Errorf("digitFromStringIndex(%s, 0) = %d, want %d", test.line, got, test.want)
		}
	}
}

func TestFindCalibrationValueWordOrDigit(t *testing.T) {
	for _, test := range []struct {
		line string
		want int
	}{
		{"1abc2", 12},
		{"two1nine", 29},
		{"pqr3stu8vwx", 38},
		{"treb7uchet", 77},
		{"onetreb7uchet", 17},
		{"4nineeightseven2", 42},
		{"onetworeb7uchetwo", 12},
	} {
		got, err := findCalibrationValueWordOrDigit(test.line)
		if err != nil {
			t.Fatalf("findCalibrationValueWordOrDigit(%s) returned error: %v", test.line, err)
		}
		if got != test.want {
			t.Errorf("findCalibrationValueWordOrDigit(%s) = %d, want %d", test.line, got, test.want)
		}
	}
}

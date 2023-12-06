package cmd

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseAlmanac(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected *Almanac
	}{
		{
			input: `
seeds: 1 2 3 4 5 6 7 8 9 10

soil-to-sand map:
2 3 1
4 5 5

`,
			expected: &Almanac{
				Seeds: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				AlmanacMapping: map[string]*AlmanacMapping{
					"soil": {
						To: "sand",
						Translations: []RangeTranslation{
							{
								DestinationStart: 2,
								SourceStart:      3,
								Range:            1,
							},
							{
								DestinationStart: 4,
								SourceStart:      5,
								Range:            5,
							},
						},
					},
				},
			},
		},
		{
			input: `
seeds: 1972667147 405592018 1450194064 27782252

seed-to-soil map:
325190047 421798005 78544109
4034765382 1473940091 137996533

humidity-to-location map:
1305211417 3371927062 89487200
947159122 0 358052295
`,
			expected: &Almanac{
				Seeds: []int{1972667147, 405592018, 1450194064, 27782252},
				AlmanacMapping: map[string]*AlmanacMapping{
					"seed": {
						To: "soil",
						Translations: []RangeTranslation{
							{
								DestinationStart: 325190047,
								SourceStart:      421798005,
								Range:            78544109,
							},
							{
								DestinationStart: 4034765382,
								SourceStart:      1473940091,
								Range:            137996533,
							},
						},
					},
					"humidity": {
						To: "location",
						Translations: []RangeTranslation{
							{
								DestinationStart: 1305211417,
								SourceStart:      3371927062,
								Range:            89487200,
							},
							{
								DestinationStart: 947159122,
								SourceStart:      0,
								Range:            358052295,
							},
						},
					},
				},
			},
		},
	} {
		almanac, err := ParseAlmanac(strings.Split(test.input, "\n"), false)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(almanac.AlmanacMapping, test.expected.AlmanacMapping) {
			t.Fatalf("expected almanac mapping to be %v, got %v", test.expected.AlmanacMapping, almanac.AlmanacMapping)
		}
	}
}

var SeedToSoilAlmanacMapping = &AlmanacMapping{
	To: "soil",
	Translations: []RangeTranslation{
		{
			DestinationStart: 50,
			SourceStart:      98,
			Range:            2,
		},
		{
			DestinationStart: 52,
			SourceStart:      50,
			Range:            48,
		},
	},
}

func TestAlmanacMappingApply(t *testing.T) {
	for _, test := range []struct {
		input    *AlmanacMapping
		seed     int
		expected int
	}{
		{
			input:    SeedToSoilAlmanacMapping,
			seed:     79,
			expected: 81,
		},
		{
			input:    SeedToSoilAlmanacMapping,
			seed:     14,
			expected: 14,
		},
		{
			input:    SeedToSoilAlmanacMapping,
			seed:     55,
			expected: 57,
		},
		{
			input:    SeedToSoilAlmanacMapping,
			seed:     13,
			expected: 13,
		},
	} {
		actual := test.input.Apply(test.seed)
		if actual != test.expected {
			t.Fatalf("expected %d, got %d", test.expected, actual)
		}
	}
}

func TestAlmanacCalculateLowestSeedLocation(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected int
	}{
		{
			input: `
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`,
			expected: 35,
		},
	} {
		almanac, err := ParseAlmanac(strings.Split(test.input, "\n"), false)
		if err != nil {
			t.Fatal(err)
		}
		actual := almanac.CalculateLowestSeedLocation()
		if actual != test.expected {
			t.Fatalf("expected %d, got %d", test.expected, actual)
		}
	}
}

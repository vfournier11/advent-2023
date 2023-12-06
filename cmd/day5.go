package cmd

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var day5Filename string
var seedsLineIsTuple bool

const defaultDay5Filename = "day_05.txt"

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use: "day5",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(day5Filename)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(data), "\n")
		almanac, err := ParseAlmanac(lines, seedsLineIsTuple)
		if err != nil {
			panic(err)
		}

		fmt.Println(almanac.CalculateLowestSeedLocation())
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
	day5Cmd.Flags().StringVarP(&day5Filename, "filename", "f", defaultDay5Filename, "Filename to read from")
	day5Cmd.Flags().BoolVarP(&seedsLineIsTuple, "seeds-line-is-tuple", "t", false, "Whether the seeds line is a pair of seed number and amount of seeds")
}

type Almanac struct {
	Seeds          []int
	AlmanacMapping map[string]*AlmanacMapping
}

func (a *Almanac) CalculateLowestSeedLocation() int {
	lowestSeedLocation := math.MaxInt32

	for _, seed := range a.Seeds {
		currentSeed := seed
		currentMapping := "seed"
		for {
			currentAlmanacMapping := a.AlmanacMapping[currentMapping]
			currentSeed = currentAlmanacMapping.Apply(currentSeed)
			if currentAlmanacMapping.To == "location" {
				if currentSeed < lowestSeedLocation {
					lowestSeedLocation = currentSeed
				}
				break
			}
			currentMapping = currentAlmanacMapping.To
		}
	}

	return lowestSeedLocation
}

type AlmanacMapping struct {
	To           string
	Translations []RangeTranslation
}

func (a *AlmanacMapping) Apply(seed int) int {
	for _, translation := range a.Translations {
		if seed >= translation.SourceStart && seed < translation.SourceStart+translation.Range {
			return translation.DestinationStart + seed - translation.SourceStart
		}
	}
	return seed
}

type RangeTranslation struct {
	DestinationStart int
	SourceStart      int
	Range            int
}

func ParseAlmanac(lines []string, seedsRepresentedInPairs bool) (*Almanac, error) {
	almanac := Almanac{
		AlmanacMapping: make(map[string]*AlmanacMapping),
	}

	var almanacMapping *AlmanacMapping
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seedsNumbersAsString := strings.Fields(line)[0:]
			for i := 1; i < len(seedsNumbersAsString); i += 2 {
				firstSeedNumber, err := strconv.Atoi(seedsNumbersAsString[i])
				if err != nil {
					return nil, fmt.Errorf("failed to parse line '%s': %w", line, err)
				}
				secondSeedNumber, err := strconv.Atoi(seedsNumbersAsString[i+1])
				if err != nil {
					return nil, fmt.Errorf("failed to parse line '%s': %w", line, err)
				}
				if seedsRepresentedInPairs {
					for j := 0; j < secondSeedNumber; j++ {
						almanac.Seeds = append(almanac.Seeds, firstSeedNumber+j)
					}
				} else {
					almanac.Seeds = append(almanac.Seeds, firstSeedNumber)
					almanac.Seeds = append(almanac.Seeds, secondSeedNumber)
				}
			}
		} else if strings.HasSuffix(line, "map:") {
			mappingName := strings.Fields(line)[0]
			fromToParts := strings.Split(mappingName, "-")
			almanacMapping = &AlmanacMapping{
				To: fromToParts[2],
			}
			almanac.AlmanacMapping[fromToParts[0]] = almanacMapping
		} else {
			var rangeTranslation RangeTranslation
			_, err := fmt.Sscanf(line, "%d %d %d", &rangeTranslation.DestinationStart, &rangeTranslation.SourceStart, &rangeTranslation.Range)
			if err != nil {
				return nil, fmt.Errorf("failed to parse line '%s' as range line: %w", line, err)
			}
			almanacMapping.Translations = append(almanacMapping.Translations, rangeTranslation)
		}
	}
	return &almanac, nil
}

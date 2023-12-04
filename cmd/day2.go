package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var day2Filename string
var day2SumPowerOfMinimum bool

const defaultDay2Filename = "day_02.txt"

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use: "day2",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(day2Filename)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(data), "\n")
		if lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1] // remove last empty line
		}

		var games []Game
		for _, line := range lines {
			var game Game

			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				panic("invalid line")
			}
			_, err = fmt.Sscanf(parts[0], "Game %d", &game.Id)
			if err != nil {
				panic(err)
			}
			for _, set := range strings.Split(parts[1], ";") {
				var setOfCube SetOfCube
				for _, cubeDescription := range strings.Split(set, ",") {
					cubeDescription = strings.TrimSpace(cubeDescription)
					var color string
					var nb int
					_, err := fmt.Sscanf(cubeDescription, "%d %s", &nb, &color)
					if err != nil {
						panic(err)
					}

					switch color {
					case "red":
						setOfCube.nbRed = nb
					case "blue":
						setOfCube.nbBlue = nb
					case "green":
						setOfCube.nbGreen = nb
					}
				}
				game.SetsOfCube = append(game.SetsOfCube, setOfCube)
			}
			if err != nil {
				panic(err)
			}
			games = append(games, game)
		}

		allCubes := SetOfCube{nbRed: 12, nbBlue: 14, nbGreen: 13}
		var sum int
		for _, game := range games {
			if day2SumPowerOfMinimum {
				sum += FindPowerOfMinimum(game)
			} else {
				if IsGameValid(game, allCubes) {
					sum += game.Id
				}
			}
		}
		fmt.Println(sum)
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
	day2Cmd.Flags().StringVarP(&day2Filename, "filename", "f", defaultDay2Filename, "Filename to read from")
	day2Cmd.Flags().BoolVarP(&day2SumPowerOfMinimum, "sum-power-of-minimum", "s", false, "Sum the power of the minimum")
}

type SetOfCube struct {
	nbRed   int
	nbBlue  int
	nbGreen int
}

type Game struct {
	Id         int
	SetsOfCube []SetOfCube
}

func IsGameValid(game Game, allCubes SetOfCube) bool {
	for _, setOfCube := range game.SetsOfCube {
		if setOfCube.nbRed > allCubes.nbRed {
			return false
		}
		if setOfCube.nbBlue > allCubes.nbBlue {
			return false
		}
		if setOfCube.nbGreen > allCubes.nbGreen {
			return false
		}
	}
	return true
}

func FindPowerOfMinimum(game Game) int {
	var maxRed, maxBlue, maxGreen int // maximum found in set of each color
	for _, setOfCube := range game.SetsOfCube {
		if setOfCube.nbRed > maxRed {
			maxRed = setOfCube.nbRed
		}
		if setOfCube.nbBlue > maxBlue {
			maxBlue = setOfCube.nbBlue
		}
		if setOfCube.nbGreen > maxGreen {
			maxGreen = setOfCube.nbGreen
		}
	}
	return maxRed * maxBlue * maxGreen
}

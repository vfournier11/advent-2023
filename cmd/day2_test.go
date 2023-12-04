package cmd

import "testing"

func TestIsGameValid(t *testing.T) {
	for _, test := range []struct {
		game     Game
		allCubes SetOfCube
		expected bool
	}{
		{
			game: Game{
				Id: 1,
				SetsOfCube: []SetOfCube{
					{
						nbRed:   1,
						nbBlue:  1,
						nbGreen: 1,
					},
				},
			},
			allCubes: SetOfCube{
				nbRed:   1,
				nbBlue:  1,
				nbGreen: 1,
			},
			expected: true,
		},
		{
			game: Game{
				Id: 1,
				SetsOfCube: []SetOfCube{
					{
						nbRed:   1,
						nbBlue:  1,
						nbGreen: 1,
					},
					{
						nbRed:   1,
						nbBlue:  0,
						nbGreen: 2,
					},
				},
			},
			allCubes: SetOfCube{
				nbRed:   1,
				nbBlue:  1,
				nbGreen: 1,
			},
			expected: false,
		},
	} {
		got := IsGameValid(test.game, test.allCubes)
		if got != test.expected {
			t.Errorf("IsGameValid(%v, %v) = %v, want %v", test.game, test.allCubes, got, test.expected)
		}
	}
}

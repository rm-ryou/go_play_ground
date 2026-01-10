package problem

import "testing"

func Test_solveMaze(t *testing.T) {
	type testCase struct {
		name string
		H, W int
		maze []string
		want int
	}

	testCases := []testCase{
		{
			name: "SからGにたどり着ける時、最短の数値を返す",
			H:    8,
			W:    8,
			maze: []string{
				".#....#G",
				".#.#....",
				"...#.##.",
				"#.##...#",
				"...###.#",
				".#.....#",
				"...#.#..",
				"S.......",
			},
			want: 16,
		},
		{
			name: "SからGにたどり着けない時、-1を返す",
			H:    8,
			W:    8,
			maze: []string{
				".#....#G",
				".#.#...#",
				"...#.##.",
				"#.##...#",
				"...###.#",
				".#.....#",
				"...#.#..",
				"S.......",
			},
			want: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := solveMaze(tc.H, tc.W, tc.maze)

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

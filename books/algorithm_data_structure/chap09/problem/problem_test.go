package problem

import "testing"

func Test_RPN(t *testing.T) {
	type testCase struct {
		name  string
		expre string
		want  int
	}

	testCases := []testCase{
		{
			name:  "return -7",
			expre: "3 4 + 1 2 - *",
			want:  -7,
		},
		{
			name:  "return 0",
			expre: "3 1 2 + -",
			want:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := RPN(tc.expre)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if act != tc.want {
				t.Errorf("want: %d, act: %d", tc.want, act)
			}
		})
	}
}

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

func Test_CheckBrackets(t *testing.T) {
	type testCase struct {
		name  string
		expre string
		want  bool
	}

	testCases := []testCase{
		{
			name:  "対応する括弧が全て存在する時、trueを返す",
			expre: "(()(())())(()())",
			want:  true,
		},
		{
			name:  "対応する括弧が存在しないものがある時、falseを返す",
			expre: "(()(())())((())",
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := CheckBrackets(tc.expre)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if act != tc.want {
				t.Errorf("want: %t, act: %t", tc.want, act)
			}
		})
	}
}

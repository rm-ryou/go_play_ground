package problem

import "testing"

func Test_Tribo(t *testing.T) {
	t.Run("トリボナッチ数列を返す", func(t *testing.T) {
		want := 13

		act := Tribo(7)
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

func Test_TriboMemo(t *testing.T) {
	t.Run("トリボナッチ数列を返す", func(t *testing.T) {
		want := 7046319384

		idx := 40
		memo := make([]int, idx+1)
		for i, _ := range memo {
			memo[i] = -1
		}

		TriboMemo(memo, idx)
		act := memo[idx]
		if act != want {
			t.Errorf("want: %d, act: %d", want, act)
		}
	})
}

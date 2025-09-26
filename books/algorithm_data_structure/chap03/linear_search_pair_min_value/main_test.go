package main

import "testing"

func Test_searchPairMinValue(t *testing.T) {
	arryA := []int{8, 5, 4}
	arryB := []int{4, 1, 9}
	k := 10
	expected := 12

	t.Run("k未満の最大値のペアを出力する", func(t *testing.T) {
		actual := searchPairMinValue(arryA, arryB, k)

		if actual != expected {
			t.Errorf("expected is %v. but actual is %v", expected, actual)
		}
	})
}

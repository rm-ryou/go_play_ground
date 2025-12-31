package main

import (
	"testing"
)

func assertSlice(t *testing.T, want, act []int) {
	t.Helper()

	if len(want) != len(act) {
		t.Errorf("size is not valid. want: %d, act: %d", len(want), len(act))
	}

	for i := range len(want) {
		if want[i] != act[i] {
			t.Errorf("want: %d, act: %d", want[i], act[i])
		}
	}
}

func TestHeap_Push(t *testing.T) {
	type testCase struct {
		name         string
		initData     []int
		pushVal      int
		expectedData []int
	}

	testCases := []testCase{
		{
			name:         "データが空の時、単純に挿入される",
			initData:     []int{},
			pushVal:      5,
			expectedData: []int{5},
		},
		{
			name:         "データが空でない時、昇順になるように挿入される",
			initData:     []int{2, 1},
			pushVal:      5,
			expectedData: []int{5, 1, 2},
			// 		2			->		5
			// 	1				->	1		2
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var h Heap
			for _, v := range tc.initData {
				h.Push(v)
			}

			h.Push(tc.pushVal)

			assertSlice(t, tc.expectedData, h.data)
		})
	}
}

func TestHeap_Pop(t *testing.T) {
	type testCase struct {
		name         string
		initData     []int
		wantErr      error
		expectedData []int
	}

	testCases := []testCase{
		{
			name:         "データが空の時、エラーが返る",
			initData:     []int{},
			wantErr:      ErrDataIsEmpty,
			expectedData: []int{},
		},
		{
			name:         "データが空でない時、昇順になるように挿入される",
			initData:     []int{6, 5, 4, 3, 2, 1},
			wantErr:      nil,
			expectedData: []int{5, 3, 4, 1, 2},
			// 				6				->				5
			// 		5				4		->		3				4
			// 	3		2		1			->	1		2
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var h Heap
			for _, v := range tc.initData {
				h.Push(v)
			}

			err := h.Pop()
			if err != tc.wantErr {
				t.Errorf("want err: %v, act: %v", tc.wantErr, err)
			}

			assertSlice(t, tc.expectedData, h.data)
		})
	}
}

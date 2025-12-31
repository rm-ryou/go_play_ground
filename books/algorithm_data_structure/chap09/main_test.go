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

func TestStack_Push(t *testing.T) {
	type testCase struct {
		name     string
		initData []int
		pushData int

		wantData []int
	}

	testCases := []testCase{
		{
			name:     "末尾に値が挿入されること",
			initData: []int{1},
			pushData: 3,
			wantData: []int{1, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var s Stack
			for _, v := range tc.initData {
				s.Push(v)
			}

			s.Push(tc.pushData)

			assertSlice(t, tc.wantData, s.data)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type testCase struct {
		name     string
		initData []int

		want     int
		wantErr  error
		wantData []int
	}

	testCases := []testCase{
		{
			name:     "値が取り出されること",
			initData: []int{1},
			want:     1,
			wantErr:  nil,
			wantData: []int{},
		},
		{
			name:     "dataが空の時エラーを返すこと",
			initData: []int{},
			want:     0,
			wantErr:  ErrDataIsEmpty,
			wantData: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var s Stack
			for _, v := range tc.initData {
				s.Push(v)
			}

			act, err := s.Pop()
			if err != tc.wantErr {
				t.Errorf("want err: %v, act: %v", tc.wantErr, err)
			}

			if act != tc.want {
				t.Errorf("want res: %d, act: %d", tc.want, act)
			}

			assertSlice(t, tc.wantData, s.data)
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type testCase struct {
		name        string
		max         int
		initData    []int
		enqueueData int

		wantErr      error
		expectedData []int
	}

	testCases := []testCase{
		{
			name:        "末尾に値が挿入されること",
			max:         3,
			initData:    []int{1},
			enqueueData: 3,

			wantErr:      nil,
			expectedData: []int{1, 3, 0},
		},
		{
			name:        "queueが満杯の時、エラーが返ること",
			max:         3,
			initData:    []int{1, 2},
			enqueueData: 3,

			wantErr:      ErrDataIsFull,
			expectedData: []int{1, 2, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := NewQueue(tc.max)
			for _, v := range tc.initData {
				q.Enqueue(v)
			}

			err := q.Enqueue(tc.enqueueData)

			if err != tc.wantErr {
				t.Errorf("want err: %v, act: %v", tc.wantErr, err)
			}
			assertSlice(t, tc.expectedData, q.data)
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type testCase struct {
		name     string
		max      int
		initData []int

		wantVal int
		wantErr error
	}

	testCases := []testCase{
		{
			name:     "値が取り出されること",
			max:      2,
			initData: []int{1},

			wantVal: 1,
			wantErr: nil,
		},
		{
			name:     "dataが空の時エラーを返すこと",
			max:      2,
			initData: []int{},

			wantVal: 0,
			wantErr: ErrDataIsEmpty,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := NewQueue(tc.max)
			for _, v := range tc.initData {
				q.Enqueue(v)
			}

			act, err := q.Dequeue()
			if err != tc.wantErr {
				t.Errorf("want err: %v, act: %v", tc.wantErr, err)
			}

			if act != tc.wantVal {
				t.Errorf("want res: %d, act: %d", tc.wantVal, act)
			}
		})
	}
}

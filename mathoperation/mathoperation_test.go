package mathoperation

import (
	"fmt"
	"testing"
)

func TestRunFindMax(t *testing.T) {
	testcases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 4},
	}

	for _, tc := range testcases {
		rev, err := FindMax(&tc.in, tc.want)
		if err != nil {
			fmt.Println(err)
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

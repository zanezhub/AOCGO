package day02_test

import (
	"fmt"
	"testing"

	"github.com/zanezhub/AOCGO/2023/day02"
)

func TestSolvePart1(t *testing.T) {
	testcases := []struct {
		time int
		dis  int
		// ex   int
	}{
		/*
			Example 1
				{7, 9, 4},
				{15, 40, 8},
				{30, 200, 9},

			Input 1
				{40, 219},
				{81, 1012},
				{77, 1365},
				{72, 1089},

			Example 2
				{71530, 940200},
		*/

		// Input 2
		{40817772, 219101213651089},
	}

	var ans int = 1

	/*
		for _, test := range testcases {
			res := day02.SolvePart1(test.time, test.dis)
			if res != test.ex {
				t.Errorf("Expected %d got %d", test.ex, res)
			}

			ans = res * ans
		}
	*/

	for _, test := range testcases {
		res := day02.SolvePart1(test.time, test.dis)

		ans = res * ans
	}

	fmt.Println(ans)

}

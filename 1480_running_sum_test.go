package l33t_test

import (
	"testing"

	"github.com/wintersparkle/l33t"

	"github.com/matryer/is"
)

func TestRunningSum(t *testing.T) {

	t.Run("Simple example", func(t *testing.T) {
		var (
			nums []int = []int{1, 2, 3}
			sums []int = []int{1, 3, 6}
		)
		is := is.New(t)
		result := l33t.RunningSum(nums)
		is.Equal(result, sums) // should equal results

	})

}

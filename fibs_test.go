package l33t_test

import (
	"testing"

	"github.com/matryer/is"
)

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func TestFib(t *testing.T) {
	is := is.New(t)

	is.Equal(fib(2), 1)
	is.Equal(fib(3), 2)
	is.Equal(fib(4), 3)

	is.Equal(fib(30), 832040)
}

package test

import "testing"
import "github.com/stretchr/testify/assert"

func TestChannel(t *testing.T) {

	res := channel()
	var expectOutput = []int{0, 1, 4}
	assert.Equal(t, res, expectOutput)

}

func channel() []int {
	src := make(chan int)
	dest := make(chan int)
	go func() {
		defer close(src)
		for i := 0; i < 3; i++ {
			src <- i
		}
	}()

	go func() {
		defer close(dest)
		for val := range src {
			dest <- val * val
		}

	}()
	res := make([]int, 0, 3)
	for i := range dest {
		res = append(res, i)
	}
	return res
}

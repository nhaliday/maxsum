package maxsum

import (
	"math"
)

func MaxSum(in <-chan int64, out chan<- int64) {
	var psum int64 = 0
	for v := range in {
		psum += v
		if psum < 0 {
			psum = 0
		}
		out <- psum
	}
}

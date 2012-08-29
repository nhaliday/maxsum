package main

import (
	"fmt"
	"io"
)

func MaxSum(in <-chan int64, out chan<- int64) {
	var psum, maxpsum int64 = 0, 0
	for v := range in {
		psum += v
		if psum < 0 {
			psum = 0
		}
		if psum > maxpsum {
			maxpsum = psum
		}
		out <- maxpsum
	}
}

func main() {
	in := make(chan int64)
	out := make(chan int64)
	var v int64
	go MaxSum(in, out)
	for {
		_, err := fmt.Scan(&v)
		if err == io.EOF {
			return
		}
		in <- v
		fmt.Println(<-out)
	}
}

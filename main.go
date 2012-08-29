package main

import (
	"fmt"
	"io"
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

func main() {
	in := make(chan int64)
	out := make(chan int64)
	go maxsum.MaxSum(in, out)
	for {
		var v int64
		_, err := fmt.Scan(&v)
		if err == io.EOF {
			return
		}
		in <- v
		fmt.Println(<-out)
	}
}

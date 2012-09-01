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
	in := make(chan int64, 1)
	defer close(in)
	out := make(chan int64, 1)
	go MaxSum(in, out)

	var v int64
	for {
		_, err := fmt.Scan(&v)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			return
		}
		in <- v
		fmt.Println(<-out)
	}
}

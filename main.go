package main

import (
	"fmt"
	"io"
	"maxsum"
)

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

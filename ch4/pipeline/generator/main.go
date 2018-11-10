package main

import (
	"fmt"
)

func main() {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valuesStream := make(chan interface{})
		go func() {
			defer close(valuesStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valuesStream <- v:
					}
				}
			}
		}()
		return valuesStream
	}

	take := func(
		done <-chan interface{},
		valuesStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valuesStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v \n", num)
	}

}

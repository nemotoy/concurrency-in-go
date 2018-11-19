package main

import "fmt"

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

	toString := func(
		done <-chan interface{},
		valuesStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valuesStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	done := make(chan interface{})
	defer close(done)

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}

	fmt.Printf("message: %s...\n", message)
}

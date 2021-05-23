package main

import "fmt"

func main() {
	done := make(chan struct{})
	count := 0

	visitOne := func() {
		count++
		go func() {
			done <- struct{}{}
		}()
	}

	visitOne()
	visitOne()

	for count > 0 {
		ok := <-done
		fmt.Println(ok)
		count--
	}
}

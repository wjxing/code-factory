package main

import "time"
import "fmt"

func main() {
	reqs := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		reqs <- i
	}
	close(reqs)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range reqs {
		<-limiter
		fmt.Println("req type1", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyReqs := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyReqs <- i
	}
	close(burstyReqs)
	for req := range burstyReqs {
		<-burstyLimiter
		fmt.Println("req type2", req, time.Now())
	}
}

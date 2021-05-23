package main

import (
	"fmt"
	"sync"
)

func doMe() {
	fmt.Println("Do me")
}

func main() {
	//once := sync.Once{}
	var once sync.Once
	once.Do(doMe)
	once.Do(doMe)
	fmt.Println("Do me done")
}

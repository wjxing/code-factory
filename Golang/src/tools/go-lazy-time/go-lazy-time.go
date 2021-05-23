package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	len := len(os.Args)
	f := "01-02 15:04:05"
	if len != 3 {
		fmt.Println("Usage:", os.Args[0], "from_time", "duration")
		fmt.Println("Ex:", os.Args[0], "\""+f+"\"", "+1000s")
		return
	}
	begin := os.Args[1]
	dur := os.Args[2]
	b, _ := time.Parse(f, begin)
	d, _ := time.ParseDuration(dur)
	fmt.Println("Begin\t:", b.Format(f))
	fmt.Println("To\t:", b.Add(d).Format(f))
}

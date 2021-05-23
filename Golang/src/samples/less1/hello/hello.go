package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Path interface {
	Base() string
	String() string
}

type PathContext interface {
	Base() string
	String() string
}

type OutputPath struct {
}

var _ Path = OutputPath{}
var _ Path = (PathContext)(nil)
var _ Path = PathContext(nil)

func (path OutputPath) Base() string {
	return "OutputPath Base"
}

func (path OutputPath) String() string {
	return "OutputPath String"
}

func max(num1, num2 int) int {
	var res int
	if num1 > num2 {
		res = num1
	} else {
		res = num2
	}
	return res
}

func main() {
	var a, b int = 10, 13
	println("Max = ", max(a, b))
	var x, y string = swap("Jie", "ba")
	println("Swap Jie ba => ", x, y)

	var name []int
	fmt.Println("Name", name)
	dev, ok := os.LookupEnv("HOME")
	if ok {
		println("Dev " + dev)
	} else {
		println("Dev nil")
	}
	cmd := exec.Command("which", "ggit")
	err := cmd.Run()
	fmt.Println("Err:", err)
}

func swap(x, y string) (string, string) {
	return y, x
}

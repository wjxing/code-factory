package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_sub(num int64, shift uint, mask int) int {
	return (int(num) >> shift) & mask
}

func show_banner() {
	fmt.Println("ION ioctl format:")
	fmt.Println("          |31    30|29    16|15     8|7      0")
	fmt.Println("----------|--------|--------|--------|--------")
	fmt.Println("Origin    |Dir     |Size    |Type    |Number")
	fmt.Println("----------|--------|--------|--------|--------")
}

func show_format(num int64) {
	n := get_sub(num, 0, 0xFF)
	t := get_sub(num, 8, 0xFF)
	s := get_sub(num, 16, 0x3FFF)
	d := get_sub(num, 30, 0x3)
	fmt.Printf("0x%-8X|%-8d|%-8d|%-8d|%-8d\n", num, d, s, t, n)
}

func main() {
	show_banner()
	for i := 1; i < len(os.Args); i++ {
		str := strings.TrimPrefix(os.Args[i], "0x")
		num, _ := strconv.ParseInt(str, 16, 64)
		show_format(num)
	}
}

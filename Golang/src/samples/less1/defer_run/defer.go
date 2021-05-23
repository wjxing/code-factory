package defer_run

import "os"
import "fmt"

func Test() {
	f := createFile("defer-test.txt")
	defer closeFile(f)
	writeFile(f)
	if f != nil {
		defer fmt.Println("Defer func")
	}
	fmt.Println("Test Done")
}

func createFile(p string) *os.File {
	fmt.Println("Create File...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("Close File...")
	f.Close()
}

func writeFile(f *os.File) {
	fmt.Println("Write File...")
	fmt.Fprintln(f, "Data")
}

func init() {
	fmt.Println("Defer init")
}

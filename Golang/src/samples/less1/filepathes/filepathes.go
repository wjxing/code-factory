package main

import (
	"fmt"
	"path/filepath"
	"sample/less1/defer_run"
)

func main() {
	defer_run.Test()
	fmt.Println("Dir", filepath.Dir("/root/filepathes.go"))
	relFile, err := filepath.Rel("./", "Android.bp")
	fmt.Println("Dir", relFile, err)
}

package main

import "fmt"
import "os/exec"
import "bytes"
import "os"
import "regexp"
import "strings"
import "sort"

var libs_cache map[string]string

func dis_dep_tree(so string, depth int) {
	_, ok := libs_cache[so]
	if ok {
		return
	}
	_, err := os.Stat(so)
	if err != nil {
		//fmt.Printf("File %s not exist\n", so)
		return
	}
	src := "readelf -d " + so + " | grep '(NEEDED)' | sed 's/^.*\\(\\[.*\\)$/\\1/' | sed 's/^\\[//; s/\\]$//'"
	cmd := exec.Command("/bin/bash", "-c", src)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		//fmt.Printf("File %s cmd err\n", so)
		return
	}
	libs_cache[so] = "Done"
	pre := strings.Repeat("\t", depth)
	list := out.String()
	rlist := strings.Split(list, "\n")
	re, _ := regexp.Compile("\n")
	list = re.ReplaceAllString(list, "\n"+pre)
	list = strings.TrimSpace(list)
	list = pre + list
	fmt.Println(pre + so + "--->")
	fmt.Println(list)
	for _, f := range rlist {
		dis_dep_tree(f, depth+1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " libname.so")
		os.Exit(1)
	}
	libs_cache = make(map[string]string)
	file_name := os.Args[1]
	dis_dep_tree(file_name, 0)
	var keys []string
	for k := range libs_cache {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("++++++Include libs:")
	fmt.Println(keys)
}

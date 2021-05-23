package main

import (
	"fmt"
	"os"
	"path/filepath"
	_ "strconv"
	"strings"
	"text/scanner"
)

func main() {
	var sc scanner.Scanner
	fileName := "Android.bp"
	r, _ := os.Open(fileName)
	defer r.Close()
	sc.Init(r)
	sc.Error = func(sc *scanner.Scanner, msg string) {
		fmt.Println("Scanner error : ", msg)
	}
	sc.Mode = scanner.ScanIdents | scanner.ScanStrings |
		scanner.ScanRawStrings | scanner.ScanComments

	for {
		tok := sc.Scan()
		if tok == scanner.EOF {
			break
		}
		/*if tok == scanner.Comment {
			fmt.Println("comment text : ", sc.TokenText())
		} else if tok == scanner.Ident {
			fmt.Println("ident text : ", sc.TokenText())
		} else if tok == scanner.String {
			fmt.Println("string text : ", sc.TokenText())
		} else if tok == scanner.RawString {
			fmt.Println("rawstring text : ", sc.TokenText())
		}*/
		//str, _ := strconv.Unquote(sc.TokenText())
		str := sc.TokenText()
		fmt.Println("text : ", str, ", tok : ", scanner.TokenString(tok))
		//pos = sc.Pos()
		//fmt.Printf("\t@ String %s - Offset : %d, Line : %d, Column : %d\n", pos.String(), pos.Offset, pos.Line, pos.Column)
	}

	fmt.Println("File Base : ",
		filepath.Base("build/blueprint/Blueprints"))
	fmt.Println("Contains : ",
		strings.ContainsAny("Blueprints*", "*?["))
	fmt.Println("Clean : ",
		filepath.Clean("build/*/blueprint"))
	files, _ := filepath.Glob("*")
	for _, file := range files {
		fmt.Println("In dir : ", file)
	}
}

package main

import "fmt"
import "flag"
import "runtime"
import "os"
import "path/filepath"

func main() {
	var output, mysrc, mybin, trimPath string

	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.StringVar(&output, "o", "", "Output file")
	flags.StringVar(&mysrc, "s", "", "Microfactory source directory (for rebuilding microfactory if necessary)")
	flags.StringVar(&mybin, "b", "", "Microfactory binary location")
	flags.StringVar(&trimPath, "trimpath", "", "remove prefix from recorded source file paths")
	flags.Parse(os.Args[1:])

	fmt.Println("output", output)
	fmt.Println("mysrc", mysrc)
	fmt.Println("mybin", mybin)
	fmt.Println("trimPath", trimPath)
	fmt.Println("Arg0", flags.Arg(0))
	fmt.Println("CPUS", runtime.NumCPU())

	srcDir, _ := filepath.Abs(".")
	fmt.Println("Abs", srcDir)

	trace, _ := os.LookupEnv("TRACE_BEGIN_SOONG")
	fmt.Println("Trace", trace)
}

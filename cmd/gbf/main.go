package main

import (
	"flag"
	"fmt"
	"gbf"
	"gbf/memory"
	"os"
)

var f string

func main() {
	flag.StringVar(&f, "f", "", "Filename to run")
	flag.Parse()

	if f == "" {
		fmt.Println("input file required")
		os.Exit(1)
	}

	data, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mem := memory.NewMem8(64 * 1024)
	gbf.Eval(mem, data, os.Stdin, os.Stdout)
}

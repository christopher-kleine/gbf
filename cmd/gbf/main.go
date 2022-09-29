package main

import (
	"embed"
	"flag"
	"fmt"
	"gbf"
	"gbf/memory"
	"io"
	"os"
	"strings"
)

var (
	exFile       string
	emFile       string
	listEmbedded bool
	bitwidth     int
	clamp        bool
)

//go:embed testing/*.b
var files embed.FS

func main() {
	flag.StringVar(&exFile, "f", "", "Filename to run")
	flag.StringVar(&emFile, "e", "", "Name of embedded file")
	flag.BoolVar(&listEmbedded, "l", false, "List embedded bf files")
	flag.IntVar(&bitwidth, "b", 8, "Bits for the interpreter (8, 16, 32")
	flag.BoolVar(&clamp, "c", false, "Clamp values instead of wrapping them")
	flag.Parse()

	var (
		err  error
		data []byte
	)

	if exFile == "-" {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if exFile != "" {
		data, err = os.ReadFile(exFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if emFile != "" {
		data, err = files.ReadFile("testing/" + emFile + ".b")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if listEmbedded {
		fileList, err := files.ReadDir("testing")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, f := range fileList {
			fmt.Println(strings.TrimSuffix(f.Name(), ".b"))
		}

		os.Exit(0)
	}

	if len(data) == 0 {
		fmt.Println("No bf file provided. Use -f to specify a filename, -e to use an embedded file or -l to list the embedded files.")
		os.Exit(1)
	}

	var mem memory.Memory
	switch bitwidth {
	case 8:
		mem = memory.NewMem8(64*1024, clamp)
	case 16:
		mem = memory.NewMem16(64*1024, clamp)
	case 32:
		mem = memory.NewMem32(64*1024, clamp)
	default:
		fmt.Println("Invalid bit value")
		os.Exit(1)
	}

	gbf.Eval(mem, data, os.Stdin, os.Stdout)
}

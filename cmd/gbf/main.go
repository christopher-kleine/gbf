package main

import (
	"embed"
	"flag"
	"fmt"
	"gbf"
	"gbf/memory"
	"os"
	"strings"
)

var (
	exFile       string
	emFile       string
	listEmbedded bool
)

//go:embed testing/*.b
var files embed.FS

func main() {
	flag.StringVar(&exFile, "f", "", "Filename to run")
	flag.StringVar(&emFile, "e", "", "Name of embedded file")
	flag.BoolVar(&listEmbedded, "l", false, "List embedded bf files")
	flag.Parse()

	var (
		err  error
		data []byte
	)

	if exFile != "" {
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

	mem := memory.NewMem8(64 * 1024)
	gbf.Eval(mem, data, os.Stdin, os.Stdout)
}

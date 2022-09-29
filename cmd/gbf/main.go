package main

import (
	"embed"
	"flag"
	"fmt"
	"gbf"
	"gbf/memory"
	"io"
	"os"
)

var (
	exFile       string
	emFile       bool
	listEmbedded bool
	bitwidth     int
	clamp        bool
	useEval      bool
	useString    string
)

//go:embed testing/*.b
var files embed.FS

func main() {
	flag.BoolVar(&emFile, "sample", false, "Use an included sample file")
	flag.BoolVar(&listEmbedded, "list", false, "List all included demo files")
	flag.IntVar(&bitwidth, "bit", 8, "Bits for the engine (8, 16, 32")
	flag.BoolVar(&clamp, "clamp", false, "Clamp values instead of wrapping them")
	flag.BoolVar(&useEval, "eval", false, "Use eval-engine instead of bytecode-engine")
	flag.StringVar(&useString, "code", "", "Direct input of bf code")
	flag.Parse()

	var (
		err  error
		data []byte
	)

	exFile = flag.Arg(0)
	if emFile {
		if exFile == "-" {
			fmt.Fprintln(os.Stderr, `You can't use pipe on sample files.
Remove the -sample option if you wish to use pipes.
Or specify one of the sample files.

To find out, what files are included, use the -list option.

Options:`)

			flag.PrintDefaults()
			os.Exit(1)
		} else if exFile != "" {
			data, err = files.ReadFile("testing/" + exFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintln(os.Stderr, `You need to specify the name of the sample.
Use the -list option to list all sample files.

Options:`)
			flag.PrintDefaults()
			os.Exit(1)
		}
	} else if exFile == "-" && useString == "" {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else if exFile != "" {
		data, err = os.ReadFile(exFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else if useString != "" {
		data = []byte(useString)
	} else if listEmbedded {
		fileList, err := files.ReadDir("testing")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, f := range fileList {
			fmt.Println(f.Name())
		}

		os.Exit(0)
	}

	if len(data) == 0 {
		flag.PrintDefaults()
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
		fmt.Fprintln(os.Stderr, `Invalid bit value. Valid are: 8, 16 or 24.

Options:`)
		flag.PrintDefaults()
		os.Exit(1)
	}

	if useEval {
		gbf.Eval(mem, data, os.Stdin, os.Stdout)
	} else {
		gbf.VM(mem, data, os.Stdin, os.Stdout)
	}
}

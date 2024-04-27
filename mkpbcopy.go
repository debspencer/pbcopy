package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/debspencer/pbcopy/pbcopier"
)

func main() {
	structName := flag.String("m", "", "Base protobuf message to create copy method")
	outFile := flag.String("o", "", "Output file (defaults to <Base function>.go")

	// process command line
	flag.Parse()

	args := flag.Args()
	if len(*structName) == 0 || len(args) != 1 {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage %s -m message [-o output] <file>.pb.go\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
	pbgo := args[0]

	pbreader, err := pbcopier.NewPbReader(pbgo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	var output string
	if len(*outFile) > 0 {
		output = *outFile
	} else {
		output = *structName + ".go"
	}

	packageName, err := pbreader.FindPackage()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	pbwriter, err := pbcopier.NewPbWriter(output, packageName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	pbCopier := pbcopier.NewPbCopier(pbreader, pbwriter)

	err = pbCopier.MkCopy(*structName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	pbwriter.Flush()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
	os.Exit(0)
}

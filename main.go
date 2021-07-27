package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/burpOverflow/goRecz/find"
	"github.com/burpOverflow/goRecz/pkg/banner"
	"github.com/burpOverflow/goRecz/pkg/fset"
	"github.com/burpOverflow/goRecz/rdl"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		banner.PrintBanner()
		fset.Usage()
		os.Exit(0)
	}
	var (
		findCmd = flag.NewFlagSet("find", flag.ExitOnError)
		rdlCmd  = flag.NewFlagSet("rdl", flag.ExitOnError)
	)

	switch os.Args[1] {
	case "find":
		findHandler(findCmd)
	case "rdl":
		rdlHandler(rdlCmd)
	default:

		os.Exit(1)
	}
	wg.Wait()
}

func findHandler(findCmd *flag.FlagSet) {
	var (
		domainPtr = findCmd.String("d", "", "domain name")
		modePtr   = findCmd.String("m", "passive", "passive mode or active mode")
		srcPtr    = findCmd.Bool("src", false, "show data source")
	)
	findCmd.Parse(os.Args[2:])

	if strings.TrimSpace(*domainPtr) == "" {
		banner.PrintBanner()
		findCmd.Usage()
		os.Exit(1)
	}

	if *modePtr == "passive" {
		find.PassiveHandler(domainPtr, srcPtr, wg)
	}
	if *modePtr == "active" {
		fmt.Println("Active Scan Not available for Now :)")
		os.Exit(0)
	}

}

func rdlHandler(rdlCmd *flag.FlagSet) {
	var (
		filePtr = rdlCmd.String("f", "", "Source file name")
		outPtr  = rdlCmd.String("o", "", "Output file name")
		slPtr   = rdlCmd.Bool("sl", false, "show output lines")
	)
	rdlCmd.Parse(os.Args[2:])

	if strings.TrimSpace(*filePtr) == "" {
		banner.PrintBanner()
		rdlCmd.Usage()
		os.Exit(1)
	}
	rdl.Handle(filePtr, outPtr, slPtr)
}

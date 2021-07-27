package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/burpOverflow/goRecz/find"
	"github.com/burpOverflow/goRecz/pkg/banner"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 2 {
		banner.PrintBanner()
		os.Exit(1)
	}
	var (
		findCmd = flag.NewFlagSet("find", flag.ExitOnError)
	)

	switch os.Args[1] {
	case "find":
		findHandler(findCmd)
	default:

		os.Exit(1)
	}

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
	wg.Wait()
}

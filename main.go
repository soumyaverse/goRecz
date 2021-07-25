package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/burpOverflow/goRecz/datasources/bufferover"
	"github.com/burpOverflow/goRecz/pkg/banner"
	"github.com/burpOverflow/goRecz/pkg/butil"
	checkerr "github.com/burpOverflow/goRecz/pkg/checkErr"
	"github.com/burpOverflow/goRecz/pkg/colors"
)

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
	// println(*domainPtr)
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

	// fmt.Println(*domainPtr)
	// fmt.Println(*modePtr)
	// fmt.Println(*srcPtr)

	if *modePtr == "passive" {
		domainSrc := make(map[string][]string)

		// fmt.Println("making passive")
		// bufferover domains
		bo := bufferover.New()
		domainList, err := bo.GetDomains(*domainPtr)
		checkerr.Check(err)
		domainSrc["BufferOver"] = domainList

		printOnConsole(domainSrc, *srcPtr)

	}
	if *modePtr == "active" {
		fmt.Println("Active Scan Not available for Now :)")
		os.Exit(0)
	}
}

func printOnConsole(domainSrc map[string][]string, srcPtr bool) {
	// fmt.Println(domainSrc["BufferOver"])
	if srcPtr {
		w := tabwriter.NewWriter(os.Stdout, 28, 1, 1, ' ', 0)
		for src, domainList := range domainSrc {
			for _, domain := range domainList {
				fmt.Fprintf(w, "%s[%s]%s\t%s%s%s\t\n", colors.Cyan, src, colors.Reset, colors.Yellow, domain, colors.Reset)
			}
		}
		w.Flush()
	} else {
		allDomain := []string{}
		for _, domainList := range domainSrc {
			allDomain = append(allDomain, domainList...)
		}
		allDomain = butil.RemoveDuplicateValuesStr(allDomain)
		for _, domain := range allDomain {
			fmt.Println(colors.Yellow + domain + colors.Reset)
		}
	}
}

package fset

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Usage() {
	fsu := make(map[string]string)

	fsu["find"] = "Find subdomains (Subdomain Enumeration)"
	fsu["rdl"] = "Remove Duplicate Lines"
	fsu["diff"] = "Different Lines bet two files"
	fsu["common"] = "Common Lines bet two files"

	fsupprint(fsu)

}

func fsupprint(fsu map[string]string) {
	w := tabwriter.NewWriter(os.Stdout, 15, 1, 1, ' ', 0)

	fmt.Println("")

	fmt.Println("Usage: " + os.Args[0] + " [Subcommand] [Options]")
	fmt.Println()
	fmt.Println()

	fmt.Println("Subcommands:")
	fmt.Println()

	for sc, usage := range fsu {
		fmt.Fprintf(w, "        %s\t%s\t\n", sc, usage)
		// fmt.Println(sc, usage)
	}
	w.Flush()

}

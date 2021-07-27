package butil

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/burpOverflow/goRecz/pkg/colors"
)

func RemoveDuplicateValuesStr(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func RemoveIndex(s []string, index int) []string {
	if index >= len(s) || index < 0 {
		log.Fatal("Index is out of range")
	}
	return append(s[:index], s[index+1:]...)
}

func PrintOnConsole(domainSrc map[string][]string, srcPtr bool) {
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
		// allDomain = RemoveDuplicateValuesStr(allDomain)
		for _, domain := range allDomain {
			fmt.Println(domain)
		}
	}
}

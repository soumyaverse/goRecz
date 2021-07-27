package rdl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	checkerr "github.com/burpOverflow/goRecz/pkg/checkErr"
)

func Handle(filePtr *string, outPtr *string, slPtr *bool) {
	domains := make(map[string]bool)

	of, err := os.Open(*filePtr)
	checkerr.Check(err)
	defer of.Close()

	scanner := bufio.NewScanner(of)

	for scanner.Scan() {
		domains[scanner.Text()] = true
	}
	if strings.TrimSpace(*outPtr) != "" {

		cf, err := os.Create(*outPtr)
		checkerr.Check(err)
		defer cf.Close()
		for domain, _ := range domains {
			_, err = cf.WriteString(domain + "\n")
			checkerr.Check(err)
			if *slPtr {
				fmt.Println(domain)
			}
		}
	} else {
		for domain, _ := range domains {
			fmt.Println(domain)
		}
	}

}

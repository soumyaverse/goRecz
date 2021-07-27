package diff

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/burpOverflow/goRecz/pkg/butil"
	checkerr "github.com/burpOverflow/goRecz/pkg/checkErr"
)

func Handle(f1Ptr *string, f2Ptr *string, outPtr *string) {
	f1List := getData(*f1Ptr)
	f2List := getData(*f2Ptr)
	diff := []string{}

	for _, f1 := range f1List {
		if !itemExists(f2List, f1) {
			diff = append(diff, f1)
		}
	}

	if strings.TrimSpace(*outPtr) != "" {
		butil.SaveArrayOnFile(*outPtr, diff)
	}

	for _, domain := range diff {
		fmt.Println(domain)
	}

}

func getData(fPtr string) []string {
	list := []string{}
	f, err := os.Open(fPtr)
	checkerr.Check(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list
}

func itemExists(arr []string, item string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return true
		}
	}
	return false
}

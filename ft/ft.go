package ft

import (
	"fmt"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/burpOverflow/goRecz/pkg/butil"
)

var w = tabwriter.NewWriter(os.Stdout, 2, 1, 1, ' ', 0)

func Handler(dlPtr *string, wg *sync.WaitGroup) {

	domainList := butil.GetListOfDataFromAFile(*dlPtr)
	// titleChan := make(chan string, 3)
	// defer close(titleChan)

	for _, d := range domainList {
		wg.Add(1)
		go gettitleCn(d, wg)
	}

	// for i := 0; i < len(domainList); {
	// 	fmt.Println(<-titleChan)
	// }
	wg.Wait()
	w.Flush()
}

func gettitleCn(d string, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Println(d+"    ", butil.GetTitle(butil.GetHTMLData(d, 2)))
	fmt.Fprintf(w, "[%s]\t%s\t\n", d, butil.GetTitle(butil.GetHTMLData(d, 2)))
}

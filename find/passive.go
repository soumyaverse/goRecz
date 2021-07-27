package find

import (
	"sync"

	"github.com/burpOverflow/goRecz/datasources/bufferover"
	"github.com/burpOverflow/goRecz/datasources/crtsh"
	"github.com/burpOverflow/goRecz/pkg/butil"
)

func PassiveHandler(domainPtr *string, srcPtr *bool, wg sync.WaitGroup) {
	allDomainListSrc := make(map[string][]string)
	domainChan := make(chan []string)
	defer close(domainChan)

	// CrtSh
	crt := crtsh.New()
	wg.Add(1)
	go crt.GetDomains(*domainPtr, domainChan, &wg)

	// BufferOver
	bo := bufferover.New()
	wg.Add(1)
	go bo.GetDomains(*domainPtr, domainChan, &wg)

	allDomainListSrc["CrtSh"] = <-domainChan
	allDomainListSrc["BufferOver"] = <-domainChan

	butil.PrintOnConsole(allDomainListSrc, *srcPtr)
}

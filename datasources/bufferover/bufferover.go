package bufferover

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/burpOverflow/goRecz/pkg/butil"
)

const BaseUrl = "https://dns.bufferover.run"

// https://dns.bufferover.run/dns?q=example.com

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (b *Client) GetDomains(domain string, domainChan chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	fetchURL := fmt.Sprintf("%s/dns?q=%s", BaseUrl, domain)

	var ret Domains
	if err := butil.FetchJSON(fetchURL, &ret); err != nil {
		log.Fatal(err)
	}
	var newDomainList []string
	var splitDomainList []string

	newDomainList = append(newDomainList, ret.FdnsA...)
	newDomainList = append(newDomainList, ret.Rdns...)

	for _, data := range newDomainList {
		splitDomainList = append(splitDomainList, strings.Split(data, ",")[1])
		// fmt.Println(strings.Split(data, ",")[1])

	}
	// fmt.Println(splitDomainList)

	// return butil.RemoveDuplicateValuesStr(splitDomainList), nil
	// srcKeyPair["BufferOver"] = butil.RemoveDuplicateValuesStr(splitDomainList)
	domainChan <- butil.RemoveDuplicateValuesStr(splitDomainList)

}

package bufferover

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/burpOverflow/goRecz/pkg/butil"
)

const BaseUrl = "https://dns.bufferover.run"

// https://dns.bufferover.run/dns?q=example.com

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (b *Client) GetDomains(domain string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/dns?q=%s", BaseUrl, domain))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var ret Domains
	var newDomainList []string
	var splitDomainList []string
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
		// log.Fatal(err)
	}

	newDomainList = append(newDomainList, ret.FdnsA...)
	newDomainList = append(newDomainList, ret.Rdns...)

	for _, data := range newDomainList {
		splitDomainList = append(splitDomainList, strings.Split(data, ",")[1])
		// fmt.Println(strings.Split(data, ",")[1])

	}
	// fmt.Println(splitDomainList)

	return butil.RemoveDuplicateValuesStr(splitDomainList), nil

}

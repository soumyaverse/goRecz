package crtsh

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/burpOverflow/goRecz/pkg/butil"
	checkerr "github.com/burpOverflow/goRecz/pkg/checkErr"
)

const BaseUrl = "https://crt.sh/"

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (crt *Client) GetDomains(domain string, domainChan chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	values := url.Values{}
	values.Add("output", "json")
	values.Add("q", domain)

	client := &http.Client{}

	req, err := http.NewRequest("GET", BaseUrl, nil)
	checkerr.Check(err)
	// defer req.Body.Close()

	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	checkerr.Check(err)

	var ret CrtShData
	var crtDomainList []string

	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		// return nil, err
		// log.Fatal(err)
		fmt.Println(err)
		domainChan <- crtDomainList
		return
	}

	for _, data := range ret {
		crtDomainList = append(crtDomainList, strings.TrimSpace(data.NameValue))
	}
	crtDomainList = butil.RemoveDuplicateValuesStr(crtDomainList)
	crtDomainListF := []string{}
	// crtDomainListD =
	for _, domain := range crtDomainList {
		if !strings.ContainsAny(domain, "*") {
			crtDomainListF = append(crtDomainListF, domain)
		}
	}
	// srcKeyPair["crtsh"] = crtDomainList
	domainChan <- crtDomainListF

	// return crtDomainList

}

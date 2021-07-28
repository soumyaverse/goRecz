package hackertarget

import (
	"bufio"
	"net/http"
	"strings"
	"sync"

	checkerr "github.com/burpOverflow/goRecz/pkg/checkErr"
)

const BaseUrl = "http://api.hackertarget.com"

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (crt *Client) GetDomains(domain string, domainChan chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	domainList := []string{}
	resp, err := http.Get(BaseUrl + "/hostsearch/?q=" + domain)

	checkerr.Check(err)
	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// checkerr.Check(err)
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		domainList = append(domainList, strings.Split(scanner.Text(), ",")[0])
	}
	// fmt.Println(domainList)
	domainChan <- domainList
}

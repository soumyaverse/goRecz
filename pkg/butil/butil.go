package butil

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	checkerr "github.com/burpOverflow/goRecz/pkg/checkErr"
	"github.com/burpOverflow/goRecz/pkg/colors"
	"golang.org/x/net/html"
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

func RemoveDuplicateValuesStr2(arr []string) []string {
	domains := make(map[string]bool)
	list := []string{}

	for _, d := range arr {
		domains[strings.TrimSpace(d)] = true
	}
	for d, _ := range domains {
		list = append(list, d)
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
	// fmt.Println(domainSrc)
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
		allDomain = RemoveDuplicateValuesStr2(allDomain)
		for _, domain := range allDomain {
			fmt.Println(domain)
		}
	}
}

func SaveArrayOnFile(filename string, arr []string) {
	file, err := os.Create(filename)
	checkerr.Check(err)
	defer file.Close()

	for _, domain := range arr {
		_, err = file.WriteString(domain + "\n")
		checkerr.Check(err)
	}
}

func FetchJSON(url string, wrapper interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	return dec.Decode(wrapper)
}

func GetTitle(HTMLString string) (title string) {
	// CREDIT: https://play.golang.org/p/0MRSefJ_-E

	r := strings.NewReader(HTMLString)
	z := html.NewTokenizer(r)

	var i int
	for {
		tt := z.Next()
		i++
		if i > 100 {
			return
		}
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data != "title" {
				continue
			}
			tt := z.Next()

			if tt == html.TextToken {
				t := z.Token()
				title = t.Data
				return
			}
		}
	}
}

func GetListOfDataFromAFile(fPtr string) []string {
	list := []string{}
	f, err := os.Open(fPtr)
	checkerr.Check(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list
}

func GetHTMLData(domain string, mode int) string {
	var HTMLData string

	if mode == 1 {

		resp, err := http.Get(domain)
		checkerr.Check(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		checkerr.Check(err)
		HTMLData = string(body)
	}
	if mode == 2 {

		resp, err := doReq("https://" + domain)
		if err != nil {
			resp, err = doReq("http://" + domain)
			// checkerr.Check(err)
			if err != nil {
				return ""
			}
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		checkerr.Check(err)
		HTMLData = string(body)
	}
	return HTMLData
}

func doReq(domain string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", domain, nil)
	checkerr.Check(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:90.0) Gecko/20100101 Firefox/90.0")

	return client.Do(req)
}

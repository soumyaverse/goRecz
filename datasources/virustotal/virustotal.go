package virustotal

import (
	"fmt"
)

var (
	BASEUrl = "https://www.virustotal.com"
	KEY     = ""
)

type VirusTotal struct {
	ApiKey string
}

func New(apiKey string) *VirusTotal {
	return &VirusTotal{ApiKey: apiKey}
}

func (vt *VirusTotal) GetDomain(domain string) {
	apiUrl := BASEUrl + "/vtapi/v2/domain/report?apikey=" + KEY + "&domain=" + domain
	fmt.Println(apiUrl)
	// resp, err := http.Get(apiUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(string(body))

}

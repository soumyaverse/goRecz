package riddler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	RIDDLER_BASE_URL = "https://riddler.io"
)

func GetAuthToken(user string, password string) {
	resource := "/auth/login"
	data := url.Values{}
	data.Set("email", user)
	data.Set("password", password)

	u, _ := url.ParseRequestURI(RIDDLER_BASE_URL)
	u.Path = resource
	urlStr := u.String() //https://riddler.io/auth/login

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
}

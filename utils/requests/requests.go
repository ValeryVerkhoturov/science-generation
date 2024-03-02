package requests

import (
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"sync"
)

func UrlWithParams(url string, params map[string]string) string {
	if len(params) == 0 {
		return url
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return url
	}
	q := req.URL.Query()
	for name, value := range params {
		q.Add(name, value)
	}
	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}

func DownloadPdfAndEncode(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error downloading:", url, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	results <- encoded
}

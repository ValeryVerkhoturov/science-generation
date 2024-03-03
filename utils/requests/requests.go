package requests

import (
	"encoding/base64"
	arxivModels "github.com/ValeryVerkhoturov/chat/utils/arxivApi/models"
	"github.com/ValeryVerkhoturov/chat/utils/formatting"
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

func DownloadPdfAndEncode(entry arxivModels.Entry, wg *sync.WaitGroup, results chan<- arxivModels.PdfFile) {
	defer wg.Done()

	resp, err := http.Get(entry.Link.Href)
	if err != nil {
		log.Println("Error downloading:", entry.Link.Href, err)
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

	published, err := formatting.FormatDate(entry.Published)
	if err != nil {
		log.Errorf("Error formatting date:", err)
		return
	}
	results <- arxivModels.PdfFile{
		Href:           entry.Link.Href,
		Name:           entry.Title,
		Published:      published,
		EncodedContent: encoded,
	}
}

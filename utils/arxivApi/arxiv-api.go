package arxivApi

import (
	"encoding/xml"
	"fmt"
	arxivModels "github.com/ValeryVerkhoturov/chat/utils/arxivApi/models"
	"github.com/ValeryVerkhoturov/chat/utils/requests"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func Query(query string, page int) (arxivModels.Feed, error) {
	url := "http://export.arxiv.org/api/query"
	params := map[string]string{"search_query": query, "max_results": "10", "start": fmt.Sprintf("%d", (page-1)*10)}

	resp, err := http.Get(requests.UrlWithParams(url, params))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %s\n", err)
		}
	}(resp.Body)

	if err != nil {
		log.Printf("Error making request: %s\n", err)
		return arxivModels.Feed{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %s\n", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s\n", err)
		return arxivModels.Feed{}, err
	}

	var feed arxivModels.Feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		log.Printf("Error parsing XML: %s\n", err)
		return arxivModels.Feed{}, err
	}

	return feed, nil
}

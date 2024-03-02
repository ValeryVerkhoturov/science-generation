package arxivApi

import (
	"encoding/xml"
	"github.com/ValeryVerkhoturov/chat/utils/requests"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// Entry defines the structure for each entry in the feed
type Entry struct {
	Title     string `xml:"title"`
	ID        string `xml:"id"`
	Published string `xml:"published"`
	Updated   string `xml:"updated"`
	Summary   string `xml:"summary"`
	Link      Link   `xml:"link"`
}

type Link struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

// Feed defines the structure of the feed received from the API
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
	ID      string   `xml:"id"`
	Link    Link     `xml:"link"`
	Updated string   `xml:"updated"`
	Entries []Entry  `xml:"entry"`
}

func Query(query string) (Feed, error) {
	url := "http://export.arxiv.org/api/query"
	params := map[string]string{"search_query": query}

	resp, err := http.Get(requests.UrlWithParams(url, params))
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Error making request: %s\n", err)
		return Feed{}, err
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
		return Feed{}, err
	}

	var feed Feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		log.Printf("Error parsing XML: %s\n", err)
		return Feed{}, err
	}
	return feed, nil
}

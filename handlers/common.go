package handlers

import (
	"fmt"
	"github.com/ValeryVerkhoturov/chat/utils/arxivApi"
	"github.com/ValeryVerkhoturov/chat/utils/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", TemplateData{Data: nil})
}

func ProcessQuery(c *gin.Context) {
	feed, err := arxivApi.Query("nlp")
	if err != nil {
		return
	}
	fmt.Println("Feed len", len(feed.Entries))

	var wg sync.WaitGroup
	results := make(chan string, len(feed.Entries))
	for _, entry := range feed.Entries {
		wg.Add(1)
		go requests.DownloadPdfAndEncode(entry.Link.Href, &wg, results)
	}
	wg.Wait()
	close(results)

	var finalOutput string
	for encoded := range results {
		fmt.Println("Encoded:", encoded)
		finalOutput += encoded // This can be modified depends on how you want to join the output
	}

	c.HTML(http.StatusOK, "index.html", TemplateData{})
}

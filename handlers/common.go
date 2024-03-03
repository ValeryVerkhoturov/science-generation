package handlers

import (
	"github.com/ValeryVerkhoturov/chat/utils/arxivApi"
	"github.com/ValeryVerkhoturov/chat/utils/formatting"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", TemplateData{
		HeaderTitle: "Index docs",
		NavbarTitle: formatting.NewNavbarTitle(formatting.Index),
	})
}

func Generate(c *gin.Context) {
	c.HTML(http.StatusOK, "generate.html", TemplateData{
		HeaderTitle: "Generate text",
		NavbarTitle: formatting.NewNavbarTitle(formatting.Generate),
	})
}

func ResetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "reset-index.html", TemplateData{
		HeaderTitle: "Reset index",
		NavbarTitle: formatting.NewNavbarTitle(formatting.ResetIndex),
	})
}

func Help(c *gin.Context) {
	c.HTML(http.StatusOK, "help.html", TemplateData{
		HeaderTitle: "Help",
		NavbarTitle: formatting.NewNavbarTitle(formatting.Help),
	})
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", TemplateData{
		HeaderTitle: "404",
	})
}

func ProcessQuery(c *gin.Context) {
	query := c.PostForm("query")
	page, err := strconv.Atoi(c.PostForm("page"))
	if page < 1 || err != nil {
		page = 1
	}
	if len(query) == 0 {
		c.HTML(http.StatusOK, "file-search-list.html", &TemplateData{Error: "Empty query"})
		return
	}

	feed, err := arxivApi.Query(query, page)
	if err != nil {
		c.HTML(http.StatusOK, "file-search-list.html", &TemplateData{Error: err.Error()})
		return
	}

	//var wg sync.WaitGroup
	//pdfFiles := make(chan arxivModels.PdfFile, len(feed.Entries))
	//for _, entry := range feed.Entries {
	//	wg.Add(1)
	//	go requests.DownloadPdfAndEncode(entry, &wg, pdfFiles)
	//}
	//wg.Wait()
	//close(pdfFiles)

	c.HTML(http.StatusOK, "file-search-list.html", TemplateData{Pagination: Pagination{Page: page, PreviousPage: page - 1, NextPage: page + 1}, Query: query, Data: feed})
}

func AppendToIndex(c *gin.Context) {
	href := c.PostForm("href")
	time.Sleep(10 * time.Second)
	if len(href) == 0 || !strings.HasPrefix(href, "http://arxiv.org/pdf/") {
		c.String(http.StatusOK, "<span class=\"text-red-500 hover:text-red-600\">Invalid link</span>")
		return
	}

	log.Info("Append to index:", href)

	c.String(http.StatusOK, "<span class=\"text-green-500 hover:text-green-600\">Successfully added</span>")
}

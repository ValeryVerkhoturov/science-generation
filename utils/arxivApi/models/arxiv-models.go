package models

import "encoding/xml"

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

type PdfFile struct {
	Name           string
	Published      string
	Href           string
	EncodedContent string
}

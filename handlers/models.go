package handlers

import "html/template"

type Pagination struct {
	PreviousPage        int
	PreviousPageLoading bool
	Page                int
	NextPage            int
	NextPageLoading     bool
}
type TemplateData struct {
	HeaderTitle string
	Pagination  Pagination
	Query       string
	Data        interface{}
	Error       string
	NavbarTitle template.HTML
}

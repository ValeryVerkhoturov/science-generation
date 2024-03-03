package requests

import (
	"net/http"
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

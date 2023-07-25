package process

import (
	"net/http"
	"time"
)

type ElmaAdapter struct {
	url    string
	token  string
	client *http.Client
}

func (e *ElmaAdapter) setHeader(req *http.Request) {
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.token)
}

func NewAdapter(url string, token string) *ElmaAdapter {
	return &ElmaAdapter{
		url:    url,
		token:  token,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

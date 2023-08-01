package elma

import (
	"net/http"
	"time"
)

type IElmaAdapter interface {
	ProcessRun(context ProcessContext) (err error)
}

type adapter struct {
	url    string
	token  string
	client *http.Client
}

func (e *adapter) setHeader(req *http.Request) {
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.token)
}

// Создание адаптера к Elma
func New(url string, token string) IElmaAdapter {
	return &adapter{
		url:    url,
		token:  token,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

package queue

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type rabbitHttpClient struct {
	client     *http.Client
	messageUrl string
}

func NewRabbitRestClient(messageUrl string) *rabbitHttpClient {
	return &rabbitHttpClient{
		client:     &http.Client{},
		messageUrl: messageUrl,
	}
}

func (rc *rabbitHttpClient) PullMessages() error {
	req, err := http.NewRequest("GET", rc.messageUrl, nil)
	if err != nil {
		return err
	}

	resp, err := rc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status not ok: %s", resp.Status)
	}

	var response struct {
		Message string `json:"message"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}

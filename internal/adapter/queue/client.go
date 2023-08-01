package queue

import (
	"action-worker/internal/action"
	"action-worker/internal/apperror"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	PullMessagePrefix = "PullMessages: "
)

type rabbitHttpClient struct {
	client    *http.Client
	apiUrl    string
	routeCode string
	routePass string
}

func NewRabbitRestClient(apiUrl string, routeCode string, routePass string) *rabbitHttpClient {
	return &rabbitHttpClient{
		client:    &http.Client{Timeout: 5 * time.Second},
		apiUrl:    apiUrl,
		routeCode: routeCode,
		routePass: routePass,
	}
}

// Получение сообщения из шины (action nil (http 204), action != nil (http 200))
func (rc *rabbitHttpClient) PullMessage() (*action.Action, error) {
	reqUrl := fmt.Sprintf("%s/pull", rc.apiUrl)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, apperror.New(apperror.ErrCreateRequest, PullMessagePrefix, http.StatusInternalServerError, err, nil)
	}

	rc.setHeader(req)

	resp, err := rc.client.Do(req)
	if err != nil {
		return nil, apperror.New(apperror.RequestNotSuccess, PullMessagePrefix, http.StatusInternalServerError, err, nil)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil, nil
	case http.StatusOK:
		var action action.Action
		err = json.NewDecoder(resp.Body).Decode(&action)
		if err != nil {
			return nil, apperror.New(apperror.ErrMarshal, PullMessagePrefix, http.StatusOK, err, nil)
		}
		return &action, nil
	case http.StatusUnauthorized:
		return nil, apperror.New(apperror.HttpStatusNotOK, PullMessagePrefix, http.StatusUnauthorized, errors.New("StatusUnauthorized"), nil)
	case http.StatusInternalServerError:
		var action action.PullMessageErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&action)
		if err != nil {
			return nil, apperror.New(apperror.HttpStatusNotOK, PullMessagePrefix, http.StatusInternalServerError, err, nil)
		}
		errText := fmt.Sprintf("message: %s, details: %s", action.Data.Message, action.Data.Details)
		return nil, apperror.New(apperror.HttpStatusNotOK, PullMessagePrefix, http.StatusInternalServerError, errors.New(errText), nil)
	default:
		errText := fmt.Sprintf("unexpected status: %s", resp.Status)
		return nil, apperror.New(apperror.HttpStatusNotOK, PullMessagePrefix, resp.StatusCode, errors.New(errText), nil)
	}
}

func (rc *rabbitHttpClient) setHeader(req *http.Request) {
	req.Header.Set("OriginRouteCode", rc.routeCode)
	req.Header.Set("OriginRoutePassword", rc.routePass)
}

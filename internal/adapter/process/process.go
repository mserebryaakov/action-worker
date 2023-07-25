package process

import (
	"action-worker/internal/apperror"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	processUrl = "bpm/template/t128.contract/createContract/run"
)

type ProcessRunRequestBody struct {
	Context ProcessContext `json:"context"`
}

type ProcessContext struct {
}

type ProcessRunResponseBody struct {
	ItemResponse
}

func (e *ElmaAdapter) ProcessRun(context ProcessContext) (err error) {
	prefix := "ProcessRun:"
	reqUrl := fmt.Sprintf("%s/%s", e.url, processUrl)

	processRequest := ProcessRunRequestBody{
		Context: context,
	}

	body, err := json.Marshal(processRequest)
	if err != nil {
		return apperror.New(apperror.ErrMarshal, fmt.Sprintf("%s ошибка перевода body в json", prefix), http.StatusInternalServerError, err, nil)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		reqUrl,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return apperror.New(apperror.ErrCreateRequest, fmt.Sprintf("%s ошибка создания запроса", prefix), http.StatusInternalServerError, err, nil)
	}

	e.setHeader(req)

	resp, err := e.client.Do(req)
	if err != nil {
		return apperror.New(apperror.ErrSendRequest, fmt.Sprintf("%s ошибка отправки запроса", prefix), http.StatusInternalServerError, err, nil)
	}
	defer resp.Body.Close()

	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return apperror.New(apperror.ErrReadBody, fmt.Sprintf("%s ошибка чтения body", prefix), http.StatusInternalServerError, err, nil)
	}

	if resp.StatusCode != http.StatusOK {
		return apperror.New(apperror.HttpStatusNotOK, fmt.Sprintf("%s код ответа - %d", prefix, resp.StatusCode), resp.StatusCode, nil, string(bts))
	}

	result := new(ProcessRunResponseBody)
	err = json.Unmarshal(bts, result)
	if err != nil {
		return apperror.New(apperror.ErrUnmarshal, fmt.Sprintf("%s ошибка парсинга json", prefix), http.StatusInternalServerError, err, nil)
	}

	if !result.Success {
		return apperror.New(apperror.RequestNotSuccess, fmt.Sprintf("%s неуспешный запрос - %s", prefix, result.Error), http.StatusInternalServerError, nil, nil)
	}

	return nil
}

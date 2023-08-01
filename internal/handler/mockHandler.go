package handler

import (
	"action-worker/internal/action"
	"action-worker/internal/adapter/elma"
	"encoding/json"
	"fmt"
)

type MockActionHandler struct {
}

func NewMock(elma elma.IElmaAdapter) IActionHandler {
	return &MockActionHandler{}
}

func (h *MockActionHandler) CreateLead(action action.CreateLeadAction) error {
	b, err := json.Marshal(action)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

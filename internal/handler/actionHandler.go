package handler

import (
	"action-worker/internal/action"
	"action-worker/internal/adapter/elma"
)

type IActionHandler interface {
	CreateLead(action action.CreateLeadAction) error
}

// Создание handler для обработки action
func New(elma elma.IElmaAdapter) IActionHandler {
	return &ActionHandler{
		elma: elma,
	}
}

type ActionHandler struct {
	elma elma.IElmaAdapter
}

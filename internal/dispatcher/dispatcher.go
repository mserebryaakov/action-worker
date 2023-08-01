package dispatcher

import (
	"action-worker/internal/action"
	"action-worker/internal/handler"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type IDispatcher interface {
	Dispatch(action action.Action) error
}

// Создание диспетчера, перенаправляющего action на обработчики
func New(handler handler.IActionHandler) IDispatcher {
	return &dispatcher{
		handler: handler,
	}
}

type dispatcher struct {
	handler handler.IActionHandler
}

func (d *dispatcher) Dispatch(actionItem action.Action) error {
	switch actionItem.Type {
	case action.CreateLeadRequestType:
		if data, ok := actionItem.Data.(map[string]interface{}); ok {
			var сreateLeadActionData action.CreateLeadActionData
			err := mapstructure.Decode(data, &сreateLeadActionData)
			if err != nil {
				textErr := fmt.Sprintf("CreateLeadActionData decode err: %s", err.Error())
				return errors.New(textErr)
			}

			createLeadAction := action.CreateLeadAction{
				Action: actionItem,
				Data:   сreateLeadActionData,
			}

			return d.handler.CreateLead(createLeadAction)
		}
		textErr := fmt.Sprintf("undexpected action data type: %s", actionItem.Data)
		return errors.New(textErr)
	default:
		textErr := fmt.Sprintf("undexpected action type: %s", actionItem.Type)
		return errors.New(textErr)
	}
}

package action

import "time"

// Структура action из rabbit REST потока
type Action struct {
	ID        string      `json:"Id"`
	Origin    string      `json:"Origin"`
	Status    string      `json:"Status"`
	CreatedOn time.Time   `json:"CreatedOn"`
	Type      string      `json:"Type"`
	Data      interface{} `json:"Data"`
}

package action

// Структура action из rabbit потока
type Action struct {
	Type       string
	Dictionary string
	Data       interface{}
}

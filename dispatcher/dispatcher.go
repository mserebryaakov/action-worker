package dispatcher

import "action-worker/action"

type IDispatcher interface {
	Dispatch(action action.Action)
}

// Конструктор диспетчера actions
func New() *Dispatcher {
	return &Dispatcher{}
}

// Диспетчер, перенаправляющий action на обработчики
type Dispatcher struct {
}

func (d *Dispatcher) Dispatch(action action.Action) {

}

package worker

import (
	"action-worker/dispatcher"
	"context"
	"time"

	"go.uber.org/zap"
)

type worker struct {
	d   dispatcher.IDispatcher
	log *zap.Logger
}

// Конструктор демон процесса worker
func New(log *zap.Logger, d dispatcher.IDispatcher) *worker {
	return &worker{
		d:   d,
		log: log,
	}
}

// Старт демон процесса worker
func (w *worker) Do(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			w.log.Warn("worker terminated by context")
			return
		case <-time.After(1 * time.Minute):
			// Отправка запроса
			// w.d.Dispatch()
		}
	}
}

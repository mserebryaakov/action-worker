package worker

import (
	"action-worker/internal/dispatcher"
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
func (w *worker) Do(ctx context.Context, frequency uint64) {
	for {
		select {
		case <-ctx.Done():
			w.log.Warn("worker terminated by context")
			return
		case <-time.After(time.Duration(frequency) * time.Second):
			w.log.Info("tick!")
		}
	}
}

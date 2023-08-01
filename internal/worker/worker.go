package worker

import (
	"action-worker/internal/adapter/queue"
	"action-worker/internal/apperror"
	"action-worker/internal/dispatcher"
	"context"
	"time"

	"go.uber.org/zap"
)

type worker struct {
	d     dispatcher.IDispatcher
	log   *zap.Logger
	queue queue.QueueClient
}

// Создание демон процесса worker (без запуска)
func New(log *zap.Logger, d dispatcher.IDispatcher, queue queue.QueueClient) *worker {
	return &worker{
		d:     d,
		log:   log,
		queue: queue,
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
			w.tick()
		}
	}
}

func (w *worker) tick() {
	action, err := w.queue.PullMessage()
	if err != nil {
		apperr, ok := err.(*apperror.AppError)
		if ok {
			w.log.Error(err.Error(), zap.Int("httpCode", apperr.HTTPCode), zap.Any("Code", apperr.Code))
		} else {
			w.log.Error(err.Error())
		}
	} else {
		if action != nil {
			err := w.d.Dispatch(*action)
			if err != nil {
				w.log.Error(err.Error())
			} else {
				w.log.Info("success dispatch")
			}
		} else {
			w.log.Info("queue is empty")
		}
	}
}

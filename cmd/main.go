package main

import (
	"action-worker/config"
	"action-worker/internal/adapter/queue"
	"action-worker/internal/dispatcher"
	"action-worker/internal/worker"
	"action-worker/pkg/zaplog"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Получение переменных окружения
	env, err := config.GetEnv()
	if err != nil {
		fmt.Println("get environment failed:", err)
		os.Exit(1)
	}

	// Создание объекта логгера
	log := zaplog.Default(env.Debug, os.Stdout)

	// Создание контекста
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создание диспетчера actions
	d := dispatcher.New()

	// Создание клиента к REST потоку RabbitMQ
	restRabbit := queue.NewRabbitRestClient(env.RestRabbitUrl)

	// Запуск демон горутины на выполнение
	doneCh := make(chan struct{})
	w := worker.New(log, d, restRabbit)
	go func() {
		w.Do(ctx, env.WorkerFrequency)
		doneCh <- struct{}{}
	}()

	// Gracefull shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sigCh:
			log.Warn("shutdown by signal")
			cancel()
		case <-doneCh:
			log.Warn("success shutdown")
			return
		}
	}
}

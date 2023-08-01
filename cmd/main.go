package main

import (
	"action-worker/config"
	"action-worker/internal/adapter/elma"
	"action-worker/internal/adapter/queue"
	"action-worker/internal/dispatcher"
	"action-worker/internal/handler"
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

	// Создание клиента к ELMA
	elmaAdapter := elma.New(env.ElmaUrl, env.ElmaToken)

	// Создание обработчика для actions
	//actionHandler := handler.New(elmaAdapter)
	actionHandler := handler.NewMock(elmaAdapter)

	// Создание диспетчера actions
	dispatch := dispatcher.New(actionHandler)

	// Создание клиента к REST потоку RabbitMQ
	restRabbit := queue.NewRabbitRestClient(env.RestRabbitUrl, env.RestRabbitRouteCode, env.RestRabbitRoutePass)

	// Запуск демон горутины на выполнение
	doneCh := make(chan struct{})
	w := worker.New(log, dispatch, restRabbit)
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
			log.Warn("service shutdown by signal")
			cancel()
		case <-doneCh:
			log.Warn("service success shutdown")
			return
		}
	}
}

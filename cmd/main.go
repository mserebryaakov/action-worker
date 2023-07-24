package main

import (
	"action-worker/config"
	"action-worker/dispatcher"
	"action-worker/pkg/zaplog"
	"action-worker/worker"
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

	// Запуск демон горутины на выполнение
	w := worker.New(log, d)
	w.Do(ctx)

	// Gracefull shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	cancel()
}

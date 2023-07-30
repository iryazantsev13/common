package service

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iryazantsev13/common/logger"
	"golang.org/x/exp/slog"
)

// awaitServiceTermination - helper функция по ожиданию работы сервиса
// Мониторит сигналы и ошибки от приложения, и помогает остановить приложение корректно
func AwaitTermination(gracefulShutdown func(), errs chan error) {
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigs:
		var interruptTimeout = 5 * time.Second
		switch sig {
		case syscall.SIGINT:
			interruptTimeout = 30 * time.Second
			slog.Info("SIGINT received. Graceful shutdown.", slog.Duration("timeout", time.Duration(interruptTimeout.Seconds())))
			gracefulShutdown()
		case syscall.SIGTERM:
			slog.Info("SIGTERM received. Trying to stop gracefully.", slog.Duration("timeout", time.Duration(interruptTimeout.Seconds())))
			gracefulShutdown()
		default:
			slog.Info("Unexpected signal received. Quiting.", "signal", sig)
			os.Exit(1)
		}

		select {
		case <-time.After(interruptTimeout):
			slog.Info("Interrupt timeout exceeded")
			os.Exit(1)
		case sig := <-sigs:
			slog.Info("Another signal received. Quiting.", "signal", sig)
			os.Exit(1)
		case err := <-errs:
			if err != nil {
				slog.Info("Service interrupted", "error", err)
				os.Exit(1)
			}
		}

	case err := <-errs:
		switch err {
		case nil:
			slog.Info("Service successfully finished it's work")
		case err:
			slog.Info("Service run failed.", logger.Err(err))
			os.Exit(1)
		}
	}
}

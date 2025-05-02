package main

import (
	"context"
	"load_balancer/backend"
	"load_balancer/balancer"
	configloading "load_balancer/config_loading"
	"load_balancer/internal/logger"
	"load_balancer/internal/messages"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func init() {
	logger.Init()

	err := configloading.LoadConfig()
	if err != nil {
		logger.Log.Error(messages.ErrLoadConfig, zap.Error(err))
	}
}

func main() {
	defer logger.Log.Sync()
	serverAddr, backendAddr, interval := configloading.SetParams()

	lb := balancer.NewBalancer()
	for _, addr := range backendAddr {
		lb.AddBack(backend.NewBackend(addr))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	go lb.HealthCheck(ctx, ticker.C)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: lb,
	}

	go func() {
		logger.Log.Info(messages.InfoBalancerON, zap.String(messages.Port, server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Error(messages.ErrLAS, zap.Error(err))
		}
	}()

	<-stop

	logger.Log.Info(messages.InfoGracefulStopStart)
	cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error(messages.ErrShutdown, zap.Error(err))
	}
	logger.Log.Info(messages.InfoGracefulStopFinish)
}

package balancer

import (
	"context"
	"load_balancer/backend"
	"net/http"
	"time"
)

type BalancerIface interface {
	AddBack(server backend.BackendIface)                    //добавить сервер в список доступных
	ServeHTTP(w http.ResponseWriter, r *http.Request)       //обработка запросов
	HealthCheck(ctx context.Context, tick <-chan time.Time) //проверка статуса серверов
}

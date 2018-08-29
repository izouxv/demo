package filter

import (
	"account-domain-rpc/metrics"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strings"
	"time"
)

func PrometheusInterceptor() grpc.UnaryServerInterceptor {
	log.Info("start Prometheus Interceptor")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod
		addCounter(method)
		metrics.AddCounter("total_count", 1) //total_count
		start := time.Now()
		resp, err := handler(ctx, req)
		addGauge(method, time.Now().Sub(start).Nanoseconds())
		return resp, err
	}
}

func addCounter(methed string) {
	metrics.AddCounter(getKey(methed, "count"), 1)
}

func addGauge(methed string, val int64) {
	metrics.AddGauge(getKey(methed, "gauge"), val)
}

func addHistograms(methed string, val int64) {
	switch methed {
	case "/pb.Sso/GetUserInfo":
		metrics.AddHistograms(getKey(methed, "histograms"), val)
	default:
	}
}

func getKey(methed string, value string) string {
	return strings.Replace(strings.ToLower(methed), "/", "_", -1)[1:] + "_" + value
}

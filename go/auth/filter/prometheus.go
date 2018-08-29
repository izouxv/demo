package filter

import (
	"strings"
	"time"

	"auth/metrics"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/cihub/seelog"
)

func PrometheusInterceptor() grpc.UnaryServerInterceptor {
	log.Info("start Prometheus Interceptor")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod
		addCounter(method)
		metrics.AddCounter("total_count", 1)
		start := time.Now()
		resp, err := handler(ctx, req)
		addGauge(method, time.Since(start).Nanoseconds())
		addHistograms(method, time.Since(start).Nanoseconds())
		return resp, err
	}
}

func addCounter(method string) {
	metrics.AddCounter(getKey(method, "count"), 1)
}

func addGauge(method string, val int64) {
	metrics.AddGauge(getKey(method, "gauge"), val)
}

func addHistograms(method string, val int64) {
	metrics.AddHistograms(getKey(method, "histograms"), val)
}

func getKey(method string, value string) string {
	return strings.Replace(strings.ToLower(method), "/", "_", -1)[1:] + "_" + value
}

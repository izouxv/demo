package metrics

import (
	"account-domain-rpc/core"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/rcrowley/go-metrics"
	"strconv"
	"strings"
	"time"
)

// PrometheusConfig provides a container with config parameters for the
// Prometheus Exporter

type PrometheusConfig struct {
	namespace     string
	Registry      metrics.Registry // Registry to be exported
	subsystem     string
	promRegistry  prometheus.Registerer //Prometheus registry
	FlushInterval time.Duration         //interval to update prom metrics
	gauges        map[string]prometheus.Gauge
	counter       map[string]prometheus.Counter
	histogram     map[string]prometheus.Histogram
	timer         map[string]prometheus.Timer
	summary       map[string]prometheus.Summary
}

// NewPrometheusProvider returns a Provider that produces Prometheus metrics.
// Namespace and subsystem are applied to all produced metrics.
func NewPrometheusProvider(r metrics.Registry, namespace string, subsystem string, promRegistry prometheus.Registerer, FlushInterval time.Duration) *PrometheusConfig {
	return &PrometheusConfig{
		namespace:     namespace,
		subsystem:     subsystem,
		Registry:      r,
		promRegistry:  promRegistry,
		FlushInterval: FlushInterval,
		gauges:        make(map[string]prometheus.Gauge),
		counter:       make(map[string]prometheus.Counter),
		histogram:     make(map[string]prometheus.Histogram),
		timer:         make(map[string]prometheus.Timer),
		summary:       make(map[string]prometheus.Summary),
	}
}

func (c *PrometheusConfig) flattenKey(key string) string {
	key = strings.Replace(key, " ", "_", -1)
	key = strings.Replace(key, ".", "_", -1)
	key = strings.Replace(key, "-", "_", -1)
	key = strings.Replace(key, "=", "_", -1)
	key = strings.Replace(key, "/", "_", -1)
	return key
}

func (c *PrometheusConfig) gaugeFromNameAndValue(name string, val float64) {
	key := fmt.Sprintf("%s_%s_%s", c.namespace, c.subsystem, name)
	g, _ := c.gauges[key]
	g = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: c.flattenKey(c.namespace),
		Subsystem: c.flattenKey(c.subsystem),
		Name:      c.flattenKey(name),
		Help:      name,
	})
	g.Set(val)
	url := "http://" + core.GetContext().Metrics.Host + ":" + strconv.Itoa(core.GetContext().Metrics.Port)
	if err := push.Collectors(
		c.flattenKey(key),
		map[string]string{},
		url,
		g,
	); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

func (c *PrometheusConfig) counterFromNameAndValue(name string, val float64) {
	key := fmt.Sprintf("%s_%s_%s", c.namespace, c.subsystem, name)
	g, _ := c.counter[key]
	g = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: c.flattenKey(c.namespace),
		Subsystem: c.flattenKey(c.subsystem),
		Name:      c.flattenKey(name),
		Help:      name,
	})
	g.Add(val)
	url := "http://" + core.GetContext().Metrics.Host + ":" + strconv.Itoa(core.GetContext().Metrics.Port)
	if err := push.Collectors(
		c.flattenKey(key),
		map[string]string{},
		url,
		g,
	); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

func (c *PrometheusConfig) histogramFromNameAndValue(name string, val float64) {
	key := fmt.Sprintf("%s_%s_%s", c.namespace, c.subsystem, name)
	g, _ := c.histogram[key]
	g = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: c.flattenKey(c.namespace),
		Subsystem: c.flattenKey(c.subsystem),
		Name:      c.flattenKey(name),
		Help:      name,
	})
	g.Observe(val)
	url := "http://" + core.GetContext().Metrics.Host + ":" + strconv.Itoa(core.GetContext().Metrics.Port)
	if err := push.Collectors(
		c.flattenKey(key),
		map[string]string{},
		url,
		g,
	); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

func (c *PrometheusConfig) summaryFromNameAndValue(name string, val float64) {
	key := fmt.Sprintf("%s_%s_%s", c.namespace, c.subsystem, name)
	g, _ := c.summary[key]
	g = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: c.flattenKey(c.namespace),
		Subsystem: c.flattenKey(c.subsystem),
		Name:      c.flattenKey(name),
		Help:      name,
	})
	g.Observe(val)
	url := "http://" + core.GetContext().Metrics.Host + ":" + strconv.Itoa(core.GetContext().Metrics.Port)
	if err := push.Collectors(
		c.flattenKey(key),
		map[string]string{},
		url,
		g,
	); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

func (c *PrometheusConfig) Write(snap metrics.Registry) error {
	snap.Each(func(key string, i interface{}) {
		switch metric := i.(type) {
		case metrics.Counter:
			c.counterFromNameAndValue(key, float64(metric.Count()))
		case metrics.Gauge:
			c.gaugeFromNameAndValue(key, float64(metric.Value()))
		case metrics.GaugeFloat64:
			c.gaugeFromNameAndValue(key, float64(metric.Value()))
		case metrics.Histogram:
			samples := metric.Snapshot().Sample().Values()
			if len(samples) > 0 {
				lastSample := samples[len(samples)-1]
				c.histogramFromNameAndValue(key, float64(lastSample))
			}
		case metrics.Meter:
		case metrics.Timer:
			lastSample := metric.Snapshot().Rate1()
			c.gaugeFromNameAndValue(key, float64(lastSample))
		default:
			return
		}
	})
	return nil
}

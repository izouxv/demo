package metrics

import (
	log "github.com/cihub/seelog"
	"github.com/rcrowley/go-metrics"
	"sync"
	"sync/atomic"
	"time"
)

const (
	eventCounter int32 = iota
	eventMeter
	eventTimer
	eventGauge
	eventHistograms
)

const (
	ElapseLess50ms  = "interval1"
	ElapseLess100ms = "interval2"
	ElapseLess200ms = "interval3"
	ElapseLess500ms = "interval4"
	ElapseMore500ms = "interval5"
	eventBufferSize = 1024 * 100
)

type event struct {
	event int32
	key   string
	value int64
}

type reporter struct {
	eventBus chan *event
	interval time.Duration
	stopping int32
	registry metrics.Registry
	writers  map[string]statWriter
	s        metrics.Sample
	evtBuf   *sync.Pool
}

var (
	sinkDuration = time.Second * 5
	reg          = &reporter{
		registry: metrics.NewRegistry(),
		stopping: 0,
		s:        metrics.NewExpDecaySample(1028, 0),
		eventBus: make(chan *event, eventBufferSize),
		writers:  make(map[string]statWriter),
		evtBuf:   &sync.Pool{New: func() interface{} { return new(event) }},
	}
)

/*Run 启动prometheus的监控服务*/
func Run() {
	//prometheusRegistry := prometheus.NewRegistry()
	//if core.GetContext().Metrics.Interval >= 1 {
	//	sinkDuration = time.Duration(core.GetContext().Metrics.Interval) * time.Second
	//}
	//
	//reg.writers[core.GetContext().Metrics.Name] = NewPrometheusProvider(reg.registry, "domain", "rpc", prometheusRegistry, 1*time.Second)
	//
	//go reg.eventLoop()
}

func (r *reporter) eventLoop() {

	if atomic.LoadInt32(&r.stopping) == 1 {
		return
	}

	ticker := time.NewTicker(sinkDuration)

	for {
		select {
		case evt := <-r.eventBus:
			r.processEvent(evt)
		case <-ticker.C:
			r.sink()
		}
	}
}

func (r *reporter) processEvent(evt *event) {
	switch evt.event {
	case eventCounter:
		metrics.GetOrRegisterCounter(evt.key, r.registry).Inc(evt.value)
	case eventMeter:
		getOrRegisterMeter(evt.key, r.registry).Mark(evt.value)
	case eventTimer:
		metrics.GetOrRegisterTimer(evt.key, r.registry).Update(time.Duration(evt.value))
	case eventHistograms:
		metrics.GetOrRegisterHistogram(evt.key, r.registry, r.s).Update(evt.value)
	case eventGauge:
		metrics.GetOrRegisterGauge(evt.key, r.registry).Update(evt.value)
	}
}

func (r *reporter) sink() {

	snap := r.snapshot()
	if snap == nil {
		return
	}

	for name, writer := range r.writers {
		if err := writer.Write(snap); err != nil {
			log.Error("metrics writer %s error : %v", name, err)
			break
		}
	}

}

func AddCounter(key string, value int64) {
	evt := reg.evtBuf.Get().(*event)
	evt.event = eventCounter
	evt.key = key
	evt.value = value
	select {
	case reg.eventBus <- evt:
	default:
		log.Info("metrics eventBus is full.")
	}
}
func AddMeter(key string, value int64) {
	evt := reg.evtBuf.Get().(*event)
	evt.event = eventMeter
	evt.key = key
	evt.value = value
	select {
	case reg.eventBus <- evt:
	default:
		log.Info("metrics eventBus is full.")
	}
}

func AddTimer(key string, duration int64) {
	evt := reg.evtBuf.Get().(*event)
	evt.event = eventTimer
	evt.key = key
	evt.value = duration
	select {
	case reg.eventBus <- evt:
	default:
		log.Info("metrics eventBus is full.")
	}
}

func AddHistograms(key string, duration int64) {
	evt := reg.evtBuf.Get().(*event)
	evt.event = eventHistograms
	evt.key = key
	evt.value = duration
	select {
	case reg.eventBus <- evt:
	default:
		log.Info("metrics eventBus is full.")
	}
}

func AddGauge(key string, duration int64) {
	evt := reg.evtBuf.Get().(*event)
	evt.event = eventGauge
	evt.key = key
	evt.value = duration
	select {
	case reg.eventBus <- evt:
	default:
		log.Info("metrics eventBus is full.")
	}
}

func (r *reporter) snapshot() metrics.Registry {
	hasElement := false
	snap := metrics.NewRegistry()

	r.registry.Each(func(key string, i interface{}) {
		switch m := i.(type) {
		case metrics.Counter:
			snap.Register(key, m.Snapshot())
			m.Clear()
		case metrics.Gauge:
			snap.Register(key, m.Snapshot())
		case metrics.GaugeFloat64:
			snap.Register(key, m.Snapshot())
		case metrics.Histogram:
			snap.Register(key, m.Snapshot())
			m.Clear()
		case metrics.Meter:
			snap.Register(key, m.Snapshot())
		case metrics.Timer:
			snap.Register(key, m.Snapshot())
		}
		hasElement = true
	})
	r.registry.UnregisterAll()
	if !hasElement {
		return nil
	}
	return snap
}

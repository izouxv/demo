package metrics

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/rcrowley/go-metrics"
	"net"
	"strconv"
	"strings"
)

const (
	messageMaxLen  = 65000
	graphiteWriter = "graphite"
)

type graphite struct {
	Host string
	Port int
	Name string
}

func newGraphite(ip, pool string, port int) *graphite {
	return &graphite{
		Host: ip,
		Port: port,
		Name: pool,
	}
}

func (g *graphite) Write(snap metrics.Registry) error {

	conn, err := net.Dial("udp", net.JoinHostPort(g.Host, strconv.Itoa(g.Port)))
	if err != nil {
		return err
	}
	ip := strings.SplitN(conn.LocalAddr().String(), ":", 2)[0]
	ip = strings.Replace(ip, ".", "_", -1)
	poolName := strings.Replace(g.Name, ".", "_", -1)
	messages := genGraphiteMessages(ip, poolName, snap)
	for _, message := range messages {
		fmt.Println(message)
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Error("graphite send message error: %v", err)
		}
	}

	return conn.Close()
}

func genGraphiteMessages(localIP string, poolName string, snap metrics.Registry) []string {
	messages := make([]string, 0)
	segments := make([]string, 0)
	segmentsLength := 0

	snap.Each(func(key string, i interface{}) {
		var segment string
		pni := strings.SplitN(key, ":", 2)
		pn := strings.Replace(pni[0], ".", "_", -1)

		switch m := i.(type) {

		case metrics.Counter:
			segment = fmt.Sprintf("grpc.%s.%s.byhost.%s.%s:%d|c\n",
				pn, poolName, localIP, pni[1], m.Count())
		case metrics.Meter:
			segment = fmt.Sprintf("grpc.%s.%s.byhost.%s.%s:%.2f|c\n",
				pn, poolName, localIP, pni[1], m.Count())
		case metrics.Timer:
			/*
				fmt.Println(m.Count())
				fmt.Println(m.Sum())
				ps := m.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999, 0.9999})
				fmt.Println(ps)

				segment = fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
					pn, poolName, localIP, pni[1], "p50", ps[0])
				segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
					pn, poolName, localIP, pni[1], "p75", ps[1])
				segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
					pn, poolName, localIP, pni[1], "p95", ps[2])
				segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
					pn, poolName, localIP, pni[1], "p99", ps[3])
				segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
					pn, poolName, localIP, pni[1], "p999", ps[4])
				segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
					pn, poolName, localIP, pni[1], "p9999", ps[5])
				segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
					pn, poolName, localIP, pni[1], "avg_time", m.Mean())
			*/
		case metrics.Gauge:
			segment = fmt.Sprintf("grpc.%s.%s.byhost.%s.%s:%d|kv\n",
				pn, poolName, localIP, pni[1], m.Value())
		//	case metrics.GaugeFloat64:
		case metrics.Histogram:
			ps := m.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999, 0.9999})
			segment = fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
				pn, poolName, localIP, pni[1], "p50", ps[0])
			segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
				pn, poolName, localIP, pni[1], "p75", ps[1])
			segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
				pn, poolName, localIP, pni[1], "p95", ps[2])
			segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
				pn, poolName, localIP, pni[1], "p99", ps[3])
			segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
				pn, poolName, localIP, pni[1], "p999", ps[4])
			segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|kv\n",
				pn, poolName, localIP, pni[1], "p9999", ps[5])
			segment += fmt.Sprintf("grpc.%s.%s.byhost.%s.%s.%s:%.2f|ms\n",
				pn, poolName, localIP, pni[1], "avg_time", m.Mean())

		default:
			return
		}
		if segmentsLength+len(segment) > messageMaxLen {
			message := strings.Join(segments, "")
			messages = append(messages, message)
			segments = make([]string, 0)
			segmentsLength = 0
		}
		segments = append(segments, segment)
		segmentsLength += len(segment)
	})

	message := strings.Join(segments, "")
	messages = append(messages, message)

	return messages
}

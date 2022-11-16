package c

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var metricC = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    `metricC`,
		Help:    `metricC`,
		Buckets: []float64{1, 2, 3},
	},
	[]string{"labelC"},
)

func init() {
	fmt.Println("Hello from init C")
}

func C() {
	fmt.Println("Hello from C")
}

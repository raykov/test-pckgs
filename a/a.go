package a

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var metricA = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    `metricA`,
		Help:    `metricA`,
		Buckets: []float64{1, 2, 3},
	},
	[]string{"labelA"},
)

func init() {
	fmt.Println("Hello from init A")
}

func A() {
	metricA.With(prometheus.Labels{"labelA": "label"}).Observe(1)

	fmt.Println("Hello from A")
}

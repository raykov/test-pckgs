package b

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var metricB = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    `metricB`,
		Help:    `metricB`,
		Buckets: []float64{1, 2, 3},
	},
	[]string{"labelB"},
)

func init() {
	fmt.Println("Hello from init B")
}

func B() {
	fmt.Println("Hello from B")
}

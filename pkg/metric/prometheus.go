package metric

import (
	"fmt"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

type Summaries map[string]prometheus.Summary

type Metric struct {
	Namespace string
	Summaries Summaries
}

func NewMetric(namespace string) *Metric {
	return &Metric{
		Namespace: namespace,
		Summaries: make(Summaries),
	}
}

func (m Metric) Observe(name string, begin time.Time) {
	summary, found := m.Summaries[name]
	if !found {
		summary = m.NewSummary(name)
	}
	summary.Observe(time.Since(begin).Seconds())
}

func (m Metric) NewSummary(name string) prometheus.Summary {
	m.Summaries[name] = promauto.NewSummary(prometheus.SummaryOpts{
		Namespace:   m.Namespace,
		Name:        name,
		Help:        fmt.Sprintf("gives execution time for %s", name),
		ConstLabels: nil,
		Objectives: map[float64]float64{
			0.5:   0.05,
			0.75:  0.025,
			0.99:  0.001,
			0.999: 0.0001,
		},
	})
	return m.Summaries[name]
}

func (m Metric) SetRoute(app *fiber.App) {
	metrics := adaptor.HTTPHandler(promhttp.Handler())
	app.Get("/metrics", metrics)
}

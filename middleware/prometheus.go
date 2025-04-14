// middleware/prometheus.go
package middleware

import (
	"building_management/metrics"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	// Register metrics
	prometheus.MustRegister(metrics.HttpRequestsTotal)
	// Register all metrics here
	prometheus.MustRegister(metrics.ApartmentCreatedCounter)
	prometheus.MustRegister(metrics.RequestDurationHistogram)
	prometheus.MustRegister(metrics.ErrorCounter)
}

func PrometheusRequestDurationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		duration := time.Since(start).Seconds()
		// Increment the histogram with labels: method, status, and path
		metrics.RequestDurationHistogram.WithLabelValues(c.Method(), strconv.Itoa(c.Response().StatusCode()), c.Path()).Observe(duration)

		// Check if the request had an error (e.g., non-2xx status code)
		if err != nil {
			metrics.ErrorCounter.WithLabelValues(c.Method(), c.Path(), "500").Inc() // Increment error counter
		} else if c.Response().StatusCode() >= 400 {
			metrics.ErrorCounter.WithLabelValues(c.Method(), c.Path(), strconv.Itoa(c.Response().StatusCode())).Inc()
		}

		return err
	}
}

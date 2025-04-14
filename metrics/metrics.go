package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	ApartmentCreatedCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "apartment_created_total",
			Help: "Total number of apartments created",
		},
	)
	RequestDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of HTTP request durations",
			Buckets: prometheus.DefBuckets, // Default bucket sizes (0.1s, 0.3s, 1s, etc.)
		},
		[]string{"method", "status", "path"}, // Labels: method, status, path
	)
	ErrorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors",
		},
		[]string{"method", "path", "status"},
	)
)

package main

// StringService provides operations on strings.
import (
	"net/http"
	"os"

	"github.com/asciiu/appa/gokit-server/middlewares"
	"github.com/asciiu/appa/gokit-server/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var svc service.StringService
	svc = service.GoKitService{}
	svc = middlewares.LogMiddleware{logger, svc}
	svc = middlewares.InstrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	var uppercase endpoint.Endpoint
	uppercase = service.MakeUppercaseEndpoint(svc)
	uppercase = middlewares.LoggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)
	uppercaseHandler := httptransport.NewServer(
		uppercase,
		service.DecodeUppercaseRequest,
		service.EncodeResponse,
	)

	var count endpoint.Endpoint
	count = service.MakeCountEndpoint(svc)
	count = middlewares.LoggingMiddleware(log.With(logger, "method", "count"))(count)
	countHandler := httptransport.NewServer(
		count,
		service.DecodeCountRequest,
		service.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("gokit-server", "HTTP", "address", "localhost:8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}

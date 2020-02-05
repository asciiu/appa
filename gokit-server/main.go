package main

// StringService provides operations on strings.
import (
	"net/http"
	"os"

	"github.com/asciiu/appa/gokit-server/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.StringService
	svc = service.GoKitService{}
	svc = service.LogMiddleware{logger, svc}

	var uppercase endpoint.Endpoint
	uppercase = makeUppercaseEndpoint(svc)
	uppercase = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)

	var count endpoint.Endpoint
	count = makeCountEndpoint(svc)
	count = loggingMiddleware(log.With(logger, "method", "count"))(count)

	uppercaseHandler := httptransport.NewServer(
		uppercase,
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		count,
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}

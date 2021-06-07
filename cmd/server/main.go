package main

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/generic"

	"github.com/sbutakov/layout-internal-api/internal/api"
)

func main() {
	server := http.Server{
		Addr: ":8080",
		Handler: api.Routes(
			log.NewNopLogger(),
			api.WithHistogramMetric(generic.NewHistogram("nop", 1)),
		),
	}

	fmt.Println(server.ListenAndServe())
}

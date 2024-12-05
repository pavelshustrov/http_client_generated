package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pavelshustrov/http_client_generated/internal/clients/service_a"
	"github.com/pavelshustrov/http_client_generated/pkg/http/client"
)

func main() {
	fmt.Println("Hello World")

	httpClient := client.New(
		client.WithAppName("test_app"),
		client.WithTimeout(time.Second*100),
		client.WithRetryFn(client.NewExponentialBackoff(
			100*time.Millisecond,
			1000*time.Millisecond,
			2.0,
			0*time.Millisecond),
		),
	)

	responses, err := service_a.NewClientWithResponses(os.Getenv("SERVICE_A_SERVER"), service_a.WithHTTPClient(httpClient))
	if err != nil {
		panic(err)
	}

	response, err := responses.PostAllInOneNameWithBodyWithResponse(
		context.Background(),
		"some_param",
		&service_a.PostAllInOneNameParams{},
		"application/json",
		strings.NewReader(`{"param": "some_params"}`),
	)
	if err != nil {
		panic(err)
	}

	_ = response
}

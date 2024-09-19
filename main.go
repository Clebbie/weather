package main

import (
	"net/http"

	"github.com/clebbie/weather/gen/weather/v1/weatherv1connect"
	"github.com/clebbie/weather/service"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	weatherService := service.NewWeatherService()
	mux := http.NewServeMux()
	route, handler := weatherv1connect.NewWeatherServiceHandler(weatherService)
	mux.Handle(route, handler)
	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
}

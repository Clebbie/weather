package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"connectrpc.com/connect"
	weatherv1 "github.com/clebbie/weather/gen/weather/v1"
	"github.com/clebbie/weather/internal"
)

// WeatherService implements the protobuf service.
type WeatherService struct {
	cache     internal.CacheInterface
	webClient internal.WebClientInterface
}

func NewWeatherService() *WeatherService {
	out := &WeatherService{}
	out.cache = internal.NewMapCache()
	out.webClient = http.DefaultClient
	return out
}

// CheckWeather checks the cache for the weather and if it fails, then it makes the request to weatherstack
func (service *WeatherService) CheckWeather(ctx context.Context, req *connect.Request[weatherv1.CheckWeatherRequest]) (*connect.Response[weatherv1.CheckWeatherResponse], error) {
	out := weatherv1.CheckWeatherResponse{}
	v, foundInCache := service.cache.RetrieveWeatherData(req.Msg.Zipcode)
	if foundInCache {
		out.CurrentTemp = fmt.Sprintf("%d C", v.Current.Temperature)
		out.IsCachedResponse = v.IsCachedResponse
		return connect.NewResponse(&out), nil
	}

	weatherURL, err := url.Parse("https://api.weatherstack.com/current")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return nil, err
	}
	weatherURL.ForceQuery = true
	weatherKey := os.Getenv("WEATHERSTACK_KEY")
	if weatherKey == "" {
		os.Stderr.WriteString("api key is missing! Please set the key as an environment variable WEATHERSTACK_KEY")
		return nil, err
	}
	q := weatherURL.Query()
	q.Set("access_key", weatherKey)
	q.Set("query", req.Msg.Zipcode)

	wreq, err := http.NewRequest(http.MethodGet, weatherURL.String()+q.Encode(), nil)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return nil, err
	}

	weatherResponse, err := service.webClient.Do(wreq)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return nil, err
	}
	defer weatherResponse.Body.Close()
	rawBytes, err := io.ReadAll(weatherResponse.Body)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return nil, err
	}
	var weatherData internal.WeatherData
	err = json.Unmarshal(rawBytes, &weatherData)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return nil, err
	}
	service.cache.AddWeatherData(req.Msg.Zipcode, weatherData)

	res := connect.NewResponse(&weatherv1.CheckWeatherResponse{
		CurrentTemp:      fmt.Sprintf("%d C", weatherData.Current.Temperature),
		IsCachedResponse: weatherData.IsCachedResponse,
	})
	return res, nil
}

package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	weatherv1 "github.com/clebbie/weather/gen/weather/v1"
	"github.com/clebbie/weather/internal"
	"github.com/magiconair/properties/assert"
)

type mockCache struct {
	shouldRetrieveErr bool
	mockData          map[string]internal.WeatherData
}

func (m *mockCache) RetrieveWeatherData(zip string) (internal.WeatherData, bool) {
	if m.shouldRetrieveErr {
		return internal.WeatherData{}, false
	}
	v, ok := m.mockData[zip]
	return v, ok
}

type mockWebClient struct {
	resp *http.Response
	err  error
}

func (m *mockWebClient) Do(req *http.Request) (*http.Response, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.resp, m.err
}

func (m *mockCache) AddWeatherData(string, internal.WeatherData) {}

type WeatherServiceTest struct {
	name         string
	service      WeatherService
	input        *connect.Request[weatherv1.CheckWeatherRequest]
	shouldErr    error
	cache        mockCache
	webclient    mockWebClient
	expectedTemp string
	wasCached    bool
}

func buildResponseBody(t *testing.T, data internal.WeatherData) io.ReadCloser {
	t.Helper()
	rawBytes, _ := json.Marshal(data)
	bb := io.NopCloser(bytes.NewReader(rawBytes))
	return bb
}

func TestWeatherService(t *testing.T) {
	testCases := []WeatherServiceTest{
		{
			name: "Happy Path not cached",
			service: WeatherService{
				cache: &mockCache{},
				webClient: &mockWebClient{
					resp: &http.Response{
						Status:     "ok",
						StatusCode: 200,
						Body:       buildResponseBody(t, internal.WeatherData{Current: internal.Current{Temperature: 999}}),
					},
				},
			},
			input: connect.NewRequest(&weatherv1.CheckWeatherRequest{
				Zipcode: "78660",
			}),
			expectedTemp: "999 C",
		},
		{
			name:      "Happy Path cached",
			wasCached: true,
			service: WeatherService{
				cache: &mockCache{
					mockData: map[string]internal.WeatherData{"78660": {Current: internal.Current{Temperature: 1000}, RequestTime: time.Now(), IsCachedResponse: true}},
				},
			},
			input: connect.NewRequest(&weatherv1.CheckWeatherRequest{
				Zipcode: "78660",
			}),
			expectedTemp: "1000 C",
		},
		{
			name: "Happy Path expired cache",
			service: WeatherService{
				cache: &mockCache{
					shouldRetrieveErr: true,
					mockData:          map[string]internal.WeatherData{"78660": {Current: internal.Current{Temperature: 1000}, RequestTime: time.Now().Add(time.Hour * -1), IsCachedResponse: true}},
				},
				webClient: &mockWebClient{
					resp: &http.Response{
						Status:     "ok",
						StatusCode: 200,
						Body:       buildResponseBody(t, internal.WeatherData{Current: internal.Current{Temperature: 800}}),
					},
				},
			},
			input: connect.NewRequest(&weatherv1.CheckWeatherRequest{
				Zipcode: "78660",
			}),
			expectedTemp: "800 C",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			r, e := c.service.CheckWeather(context.Background(), c.input)
			assert.Equal(t, e, c.shouldErr)
			assert.Equal(t, r.Msg.CurrentTemp, c.expectedTemp)
			assert.Equal(t, r.Msg.IsCachedResponse, c.wasCached)
		})
	}
}

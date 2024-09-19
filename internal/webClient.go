package internal

import (
	"net/http"
)

type WebClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

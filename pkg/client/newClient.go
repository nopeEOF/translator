package client

import (
	"net/http"
	"time"
)

func NewClient(timeout int) Client {
	request := Client{
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
	return request
}
package client

import "net/http"

type Client struct {
	Scheme string
	Host   string
	Port   string
	client *http.Client
}

func NewClient(scheme, host, port string) *Client {
	return &Client{
		Scheme: scheme,
		Host:   host,
		Port:   port,
		client: &http.Client{},
	}
}

package client

import (
	"github.com/go-resty/resty"
)

type URL struct {
	Schema string `schema`
	Host   string `host`
	path   string `path`
}

type Request struct {
	Url     URL               `url`
	Headers map[string]string `headers`
	Params  map[string]string `params`
	Body    string            `body`
	Method  string            `method`
}

type Response struct {
	Status int    `status`
	Body   string `body`
}

func (this *URL) Url2String() {
	return "test"
}

func Get(req Request) (*Response, error) {
	resp, err := resty.R().Get(req.Url.Url2String())
	if err != nil {

	}

	return nil, err
}

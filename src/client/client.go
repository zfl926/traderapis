package client

import (
	"gopkg.in/resty.v1"
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

func (this *URL) Url2String() string {
	return this.Schema + "/" + this.Host + "/" + this.path
}

func Get(req Request) (*Response, error) {
	_, err := resty.R().Get(req.Url.Url2String())
	if err != nil {
		return nil, err
	}
	return nil, err
}

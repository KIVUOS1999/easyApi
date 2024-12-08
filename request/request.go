package request

import (
	"net/http"
	"net/url"
)

type Req struct {
	Req        *http.Request
	PathParam  map[string]string
	QueryParam url.Values
}

func (h *Req) Request() *http.Request {
	return h.Req
}

func (h *Req) Method() string {
	return h.Req.Method
}

func (h *Req) GetHeader(key string) string {
	return h.Req.Header.Get(key)
}

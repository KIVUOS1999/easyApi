package app

import (
	"encoding/json"

	"github.com/KIVUOS1999/easyApi/request"
	"github.com/KIVUOS1999/easyApi/response"
)

type Context struct {
	Request  request.Req
	Response response.Resp
}

func (ctx *Context) PathParam(key string) string {
	return ctx.Request.PathParam[key]
}

func (ctx *Context) QueryParam(key string) string {
	return ctx.Request.QueryParam.Get(key)
}

func (ctx *Context) Bind(inp interface{}) error {
	body := ctx.Request.Req.Body

	err := json.NewDecoder(body).Decode(inp)
	if err != nil {
		return err
	}

	return nil
}

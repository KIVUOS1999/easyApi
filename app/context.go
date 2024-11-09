package app

import (
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

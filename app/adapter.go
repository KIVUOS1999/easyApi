package app

import (
	"net/http"

	"github.com/KIVUOS1999/easyApi/request"
	"github.com/KIVUOS1999/easyApi/response"
	"github.com/gorilla/mux"
)

func (a *app) adapterFunc(f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := response.Resp{Resp: w, Req: r}

		ctx := Context{
			Request:  request.Req{Req: r, PathParam: mux.Vars(r), QueryParam: r.URL.Query()},
			Response: response.Resp{Resp: w, Req: r},
		}

		resp.Resp.Header().Add("Content-Type", "application/json")

		output, err := f(&ctx)
		if err != nil {
			a.handleErrorResponse(resp, err)

			return
		}

		resp.WriteResponse(output)
	}
}

package response

import (
	"encoding/json"
	"net/http"
)

type model struct {
	Data interface{} `json:"data"`
}

type Resp struct {
	Resp http.ResponseWriter
	Req  *http.Request
}

func (r *Resp) SetHeaders(key, value string) {
	r.Resp.Header().Add(key, value)
}

func (r *Resp) Response() http.ResponseWriter {
	return r.Resp
}

func (r *Resp) WriteResponse(resp interface{}) {
	method := r.Req.Method
	statusCode := http.StatusOK

	switch method {
	case http.MethodPost:
		statusCode = http.StatusCreated

	case http.MethodDelete, http.MethodPut, http.MethodPatch:
		if resp == nil {
			statusCode = http.StatusNoContent

			return
		}
	}

	r.Resp.WriteHeader(statusCode)

	if resp == nil {
		return
	}

	err := json.NewEncoder(r.Resp).Encode(resp)
	if err != nil {
		r.Resp.WriteHeader(statusCode)

		jsonResp := model{
			Data: resp,
		}

		json.NewEncoder(r.Resp).Encode(jsonResp)
	}
}

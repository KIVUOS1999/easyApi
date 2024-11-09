package app

import (
	"encoding/json"
	"errors"
	"net/http"

	easyError "github.com/KIVUOS1999/easyApi/errors"
	"github.com/KIVUOS1999/easyApi/response"
)

func (a *app) handleErrorResponse(resp response.Resp, err error) {
	var customError *easyError.CustomError

	if errors.As(err, &customError) {
		resp.Resp.WriteHeader(customError.StatusCode)
		json.NewEncoder(resp.Resp).Encode(customError)

		return
	}

	customErr := easyError.CustomError{
		StatusCode: http.StatusInternalServerError,
		Response:   err.Error(),
	}

	resp.Resp.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(resp.Resp).Encode(customErr)
}

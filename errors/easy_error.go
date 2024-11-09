package easyError

import "strconv"

type CustomError struct {
	StatusCode int    `json:"status_code"`
	Response   string `json:"response"`
}

func (e *CustomError) Error() string {
	return "Error Code: " + strconv.Itoa(e.StatusCode) + " , Message: " + e.Response + ""
}

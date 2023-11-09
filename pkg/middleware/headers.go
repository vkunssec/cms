package middleware

import "net/http"

func GetHeaders() http.Header {
	headers := http.Header{}

	headers.Set("Access-Control-Allow-Headers", "*")
	headers.Set("Access-Control-Allow-Origin", "*")
	headers.Set("Content-Type", "application/json; charset=utf-8")

	return headers
}

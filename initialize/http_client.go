package initialize

import (
	"ff/g"
	"net/http"
)

func InitHttpClient() {
	g.HttpClient = &http.Client{}
}

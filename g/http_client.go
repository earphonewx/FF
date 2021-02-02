package g

import (
	"net/http"
	"sync"
)

type httpClient struct {
	*http.Client
	once sync.Once
}

var hc httpClient

func HttpClient() *http.Client {
	hc.once.Do(func() {
		hc.Client = &http.Client{}
	})
	return hc.Client
}
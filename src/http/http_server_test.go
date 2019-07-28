package http_learing

import (
	"testing"
	"net/http"
)

func TestHttpServer(t *testing.T) {
	http.ListenAndServe(":9090", nil)
}


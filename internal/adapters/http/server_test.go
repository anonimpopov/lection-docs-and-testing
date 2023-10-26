package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	var s Server

	cases := []struct{
		code int
	}{
		{
			code: http.StatusOK,
		},
		{
			code: http.StatusInternalServerError,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Healthz:%d", c.code),func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
			w := httptest.NewRecorder()
			s.healthzHandler(w, req)
			r := w.Result()
			if r.StatusCode != c.code {
				t.Fatalf("Expected %d, but was %d", c.code, r.StatusCode)
			}
		})
	}



}
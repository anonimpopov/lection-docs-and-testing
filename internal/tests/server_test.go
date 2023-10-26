//go:build testing
// +build testing

package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerSuccess(t *testing.T) {
	c := http.Client{}
	r, err := c.Get(fmt.Sprintf("http://localhost:%d/healthz", s.Port()))
	if err != nil {
		t.Fatalf("healthz request failed: %s", err)
	}

	if r.StatusCode != http.StatusOK {
		t.Fatalf("healthz request failed. Expect code 200, but was %d", r.StatusCode)
	}
}

func TestServerFailed(t *testing.T) {
	c := http.Client{}
	r, err := c.Get(fmt.Sprintf("http://localhost:%d/healthz", s.Port()))
	if err != nil {
		t.Fatalf("healthz request failed: %s", err)
	}

	if r.StatusCode != http.StatusInternalServerError {
		t.Fatalf("healthz request failed. Expect code 500, but was %d", r.StatusCode)
	}
}

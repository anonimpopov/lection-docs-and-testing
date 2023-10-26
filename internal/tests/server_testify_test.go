//go:build testify
// +build testify

package tests

import (
	"net/http"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerTestifySuccess(t *testing.T) {
	c := http.Client{}
	r, err := c.Get(fmt.Sprintf("http://localhost:%d/healthz", s.Port()))
	if err != nil {
		t.Fatalf("healthz request failed: %s", err)
	}

	require.Equal(t, http.StatusOK, r.StatusCode)
}

func TestServerTestifyFailed(t *testing.T) {
	a := assert.New(t)

	c := http.Client{}
	r, err := c.Get(fmt.Sprintf("http://localhost:%d/healthz", s.Port()))
	if err != nil {
		require.NoError(t, err)
	}

	a.Equal(http.StatusInternalServerError, r.StatusCode)
}

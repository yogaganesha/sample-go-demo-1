package main

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	http.HandleFunc("/", handler)
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, r.StatusCode, http.StatusOK)
	t.Log(response)
}

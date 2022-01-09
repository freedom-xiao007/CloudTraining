package main

import (
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/v1/hello")
	if err != nil {
		log.Fatal(err)
		return
	}

	for key, values := range resp.Header {
		log.Println(key, values)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "default", resp.Header.Get("Version"))
}

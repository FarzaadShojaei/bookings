package main

import (
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var myH myHandler

	h := Nosurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error("type is not http")
	}
}

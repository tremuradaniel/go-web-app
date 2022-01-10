package main

import (
	"net/http"
	"os"
	"testing"
)

// everything in this file will run before each test
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

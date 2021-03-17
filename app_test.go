package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNegativeIDEndpoint(t *testing.T) {
	req := httptest.NewRequest("DELETE", "http://localhost:8080/api/movies/delete/-1", nil)

	app := Setup()

	// http.Response
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
	// Do something with results:
	if resp.StatusCode == 404 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body)) // => Hello, World!
	}

}

func TestGetMovie(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/api/movies/Space", nil)

	app := Setup()

	// http.Response
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	// Do something with results:
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body)) // => Hello, World!
	}

}

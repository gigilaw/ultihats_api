package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gigilaw/ultihats/routes"
	"github.com/stretchr/testify/assert"
)

func TestRegisterEmailRoute(t *testing.T) {
	router := routes.ApiRoutes()

	payload := strings.NewReader(`{
		"email":"abc@gmail.com",
		"password": "abc12345"
		}`,
	)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/login", payload)

	fmt.Println(req)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	fmt.Println(w)
	assert.Equal(t, 200, w.Code)
}

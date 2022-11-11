package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"github.com/gigilaw/ultihats/config"
	"github.com/gigilaw/ultihats/handlers"

	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/routes"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

type RegisterUser struct {
	FirstName string
	LastName  string
	Password  string
	Height    int
	Gender    string
	Email     string
	Birthday  string
}

var testingUser = RegisterUser{
	FirstName: faker.FirstName(),
	LastName:  faker.LastName(),
	Password:  "abc12345",
	Height:    100,
	Gender:    "Female",
	Email:     faker.Email(),
	Birthday:  faker.Date(),
}

var router = routes.ApiRoutes()

func TestMain(m *testing.M) {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestUserEmailRegister(t *testing.T) {
	encoded, err := json.Marshal(testingUser)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/register/email", bytes.NewBuffer(encoded))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}
	router.ServeHTTP(w, req)

	res := handlers.DecodeResponse(w.Body)

	assert.Equal(t, config.HTTP_SUCCESS, w.Code)
	assert.Equal(t, testingUser.Email, res["Email"])
	assert.NotEmpty(t, res["DiscSkills"])
}

func TestUserLogin(t *testing.T) {
	login := map[string]string{
		"email":    testingUser.Email,
		"password": testingUser.Password,
	}

	encoded, err := json.Marshal(login)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(encoded))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	router.ServeHTTP(w, req)

	res := handlers.DecodeResponse(w.Body)

	assert.Equal(t, config.HTTP_SUCCESS, w.Code)
	assert.Equal(t, testingUser.Email, res["Email"])
	assert.NotEmpty(t, res["DiscSkills"])
}

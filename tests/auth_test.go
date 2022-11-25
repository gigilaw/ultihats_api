package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gigilaw/ultihats/config"
	"github.com/gigilaw/ultihats/handlers"
	"github.com/gigilaw/ultihats/initializers"
	"github.com/gigilaw/ultihats/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	FirstName string
	LastName  string
	Password  string
	Height    int
	Gender    string
	Email     string
	Birthday  string
}

var testingUser = TestUser{
	FirstName: faker.FirstName(),
	LastName:  faker.LastName(),
	Password:  "abc12345",
	Height:    100,
	Gender:    "Female",
	Email:     faker.Email(),
	Birthday:  faker.Date(),
}

var AuthRouter *gin.Engine

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	AuthRouter = routes.ApiRoutes()
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
	AuthRouter.ServeHTTP(w, req)

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

	AuthRouter.ServeHTTP(w, req)

	res := handlers.DecodeResponse(w.Body)

	assert.Equal(t, config.HTTP_SUCCESS, w.Code)
	assert.Equal(t, testingUser.Email, res["Email"])
	assert.NotEmpty(t, res["DiscSkills"])
}

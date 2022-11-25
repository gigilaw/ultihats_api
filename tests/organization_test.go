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

type TestOrganization struct {
	Name     string
	Email    string
	Password string
	City     string
	Est      int
}

var testingOrganization = TestOrganization{
	Name:     faker.Name(),
	Email:    faker.Email(),
	Password: "abc12345",
	City:     "Hong Kong",
	Est:      2020,
}

var OrgRouter *gin.Engine

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	OrgRouter = routes.ApiRoutes()
}

func TestOrganizationEmailRegister(t *testing.T) {
	encoded, err := json.Marshal(testingOrganization)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/organization/register", bytes.NewBuffer(encoded))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}
	OrgRouter.ServeHTTP(w, req)

	res := handlers.DecodeResponse(w.Body)
	fmt.Println(res)
	assert.Equal(t, config.HTTP_SUCCESS, w.Code)
	assert.Equal(t, testingOrganization.Email, res["Email"])
}

func TestOrganizationLogin(t *testing.T) {
	login := map[string]string{
		"email":    testingOrganization.Email,
		"password": testingOrganization.Password,
	}

	encoded, err := json.Marshal(login)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/organization/login", bytes.NewBuffer(encoded))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	OrgRouter.ServeHTTP(w, req)

	res := handlers.DecodeResponse(w.Body)
	fmt.Println(res["Email"])

	assert.Equal(t, config.HTTP_SUCCESS, w.Code)
	assert.Equal(t, testingOrganization.Email, res["Email"])
}

package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iqbalbiondy/18_Soal_Unit_Testing/Praktikum/Prioritas2/config"
	"github.com/iqbalbiondy/18_Soal_Unit_Testing/Praktikum/Prioritas2/controller"
	"github.com/iqbalbiondy/18_Soal_Unit_Testing/Praktikum/Prioritas2/models"
	
	"github.com/labstack/echo/v4"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	config.InitMigrateTest()
	e := echo.New()
	return e
}

func InsertData() error {
	user := models.User{
		ID:       1,
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

type TestCase struct {
	Method         string
	Name           string
	Path           string
	ExpectStatus   int
	ExpectResponse string
}

 
func TestLoginUser(t *testing.T) {
	defer func() {
		successTest := TestCase{

			Method:         http.MethodPost,
			Name:           "Login User",
			Path:           "/login",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success login",
		}

		e := InitEchoTestAPI()
		InsertData()
		reqStr := `{
			"id": 1,
			"name": "Alta",
			"email": "alta@gmail.com",
			"password": "123"
		}`

		req := httptest.NewRequest(successTest.Method, successTest.Path, strings.NewReader(reqStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.LoginUserController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)

		}
	}()

	failedTest := TestCase{

		Method:         http.MethodPost,
		Name:           "Login User",
		Path:           "/login",
		ExpectStatus:   http.StatusUnauthorized,
		ExpectResponse: "Fail login",
	}

	e := InitEchoTestAPI()
	InsertData()
	reqStr := `{
		"email": "altaa@gmail.com",
		"password": "123"
	}`

	req := httptest.NewRequest(failedTest.Method, failedTest.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(failedTest.Path)
	if assert.NoError(t, controller.LoginUserController(c)) {
		assert.Equal(t, failedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failedTest.ExpectResponse)

	}

}

func TestGetUsers(t *testing.T) {

	defer func() {
		successTest := TestCase{

			Method:         http.MethodGet,
			Name:           "Get All Users",
			Path:           "/Users",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success get all users",
		}

		e := InitEchoTestAPI()
		InsertData()

		req := httptest.NewRequest(successTest.Method, successTest.Path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.GetUsersController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)
		}
	}()
	failTest := TestCase{

		Method:         http.MethodGet,
		Name:           "Fail Get All Users",
		Path:           "/Users",
		ExpectStatus:   http.StatusBadRequest,
		ExpectResponse: "Database empty",
	}

	e := InitEchoTestAPI()

	req := httptest.NewRequest(failTest.Method, failTest.Path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(failTest.Path)
	if assert.NoError(t, controller.GetUsersController(c)) {
		assert.Equal(t, failTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failTest.ExpectResponse)
	}

}

func TestGetUser(t *testing.T) {

	defer func() {
		successTest := TestCase{

			Method:         http.MethodGet,
			Name:           "Get User",
			Path:           "/",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success get user",
		}

		e := InitEchoTestAPI()
		InsertData()
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

		req := httptest.NewRequest(successTest.Method, successTest.Path, nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.GetUserController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)
		}
	}()
	failedTest := TestCase{

		Method:         http.MethodGet,
		Name:           "Fail Get User",
		Path:           "/",
		ExpectStatus:   http.StatusBadRequest,
		ExpectResponse: "User not found",
	}
	e := InitEchoTestAPI()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(failedTest.Method, failedTest.Path, nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(failedTest.Path)
	if assert.NoError(t, controller.GetUserController(c)) {
		assert.Equal(t, failedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failedTest.ExpectResponse)
	}

}

func TestCreateUser(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodPost,
		Name:           "Create user",
		Path:           "/Users",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success add user",
	}

	e := InitEchoTestAPI()
	reqStr := `{
		"id": 2,
		"name":     "Alta",
		"password": "123",
		"email":    "alta@gmail.com"
		}`
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"
	req := httptest.NewRequest(testcase.Method, testcase.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.CreateUserController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}

}

func TestUpdateUser(t *testing.T) {
	defer func() {
		successTest := TestCase{

			Method:         http.MethodPut,
			Name:           "Update user",
			Path:           "/",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success Edit User",
		}

		e := InitEchoTestAPI()
		InsertData()

		reqStr := `{
			"id":       1,
			"name":     "Alterra",
			"password": "123",
			"email":    "alta@gmail.com"
			}`
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

		req := httptest.NewRequest(successTest.Method, successTest.Path, strings.NewReader(reqStr))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.UpdateUserController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)
		}
	}()

	failedTest := TestCase{

		Method:         http.MethodPut,
		Name:           "Update user",
		Path:           "/",
		ExpectStatus:   http.StatusBadRequest,
		ExpectResponse: "Bad Request",
	}

	e := InitEchoTestAPI()
	InsertData()

	reqStr := `{
		"id":       2,
		"name":     "Alterra",
		"password": "123",
		"email":    "alta@gmail.com"
		}`
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(failedTest.Method, failedTest.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(failedTest.Path)
	if assert.NoError(t, controller.UpdateUserController(c)) {
		assert.Equal(t, failedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failedTest.ExpectResponse)
	}

}

func TestDeleteUser(t *testing.T) {
	// token
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	defer func() {
		successTest := TestCase{

			Method:         http.MethodDelete,
			Name:           "Delete User",
			Path:           "/",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success Delete User",
		}

		e := InitEchoTestAPI()
		InsertData()

		req := httptest.NewRequest(successTest.Method, successTest.Path, nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.DeleteUserController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)
		}
	}()
	failedTest := TestCase{

		Method:         http.MethodDelete,
		Name:           "Failed delete User",
		Path:           "/",
		ExpectStatus:   http.StatusBadRequest,
		ExpectResponse: "Failed Delete",
	}

	e := InitEchoTestAPI()

	req := httptest.NewRequest(failedTest.Method, failedTest.Path, nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c.SetPath(failedTest.Path)
	if assert.NoError(t, controller.DeleteUserController(c)) {
		assert.Equal(t, failedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failedTest.ExpectResponse)
	}
}
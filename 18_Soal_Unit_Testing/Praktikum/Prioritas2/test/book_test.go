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

func InsertBook() error {
	books := models.Book{
		ID:        1,
		Title:     "Laskar Putih",
		Author:    "Iqbal Biondy",
		Publisher: "Mizon",
	}

	var err error
	if err = config.DB.Save(&books).Error; err != nil {
		return err
	}
	return nil
}

type TestCaseBook struct {
	Method         string
	Name           string
	Path           string
	ExpectStatus   int
	ExpectResponse string
}


func TestGetBooks(t *testing.T) {
	defer func() {
		successTest := TestCaseBook{

			Method:         http.MethodGet,
			Name:           "Get All Books",
			Path:           "/books",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success get all book",
		}

		e := InitEchoTestAPI()
		InsertBook()

		req := httptest.NewRequest(successTest.Method, successTest.Path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.GetBooksController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)

		}
	}()
	failedTest := TestCaseBook{

		Method:         http.MethodGet,
		Name:           "Failed get All Books",
		Path:           "/books",
		ExpectStatus:   http.StatusBadRequest,
		ExpectResponse: "Database Empty",
	}

	e := InitEchoTestAPI()

	req := httptest.NewRequest(failedTest.Method, failedTest.Path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(failedTest.Path)
	if assert.NoError(t, controller.GetBooksController(c)) {
		assert.Equal(t, failedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failedTest.ExpectResponse)

	}

}

func TestGetBook(t *testing.T) {
	defer func() {
		successTest := TestCase{

			Method:         http.MethodGet,
			Name:           "Get book",
			Path:           "/",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success get book",
		}

		e := InitEchoTestAPI()
		InsertBook()

		req := httptest.NewRequest(successTest.Method, successTest.Path, nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.GetBookController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)
		}
	}()
	failedTest := TestCase{

		Method:         http.MethodGet,
		Name:           "failed get book",
		Path:           "/",
		ExpectStatus:   http.StatusBadRequest,
		ExpectResponse: "Book not found",
	}

	e := InitEchoTestAPI()
	InsertBook()

	req := httptest.NewRequest(failedTest.Method, failedTest.Path, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.SetPath(failedTest.Path)
	if assert.NoError(t, controller.GetBookController(c)) {
		assert.Equal(t, failedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failedTest.ExpectResponse)
	}
}

func TestCreateBook(t *testing.T) {
	testcase := TestCase{

		Method:         http.MethodPost,
		Name:           "Create book",
		Path:           "/books",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Success add book",
	}

	e := InitEchoTestAPI()
	reqStr := `{
		"ID":        2,
		"title":     "Laskar Pelangi 2",
		"author":    "Giring",
		"publisher": "Mizan"
		}`
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"

	req := httptest.NewRequest(testcase.Method, testcase.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase.Path)
	if assert.NoError(t, controller.CreateBookController(c)) {
		assert.Equal(t, testcase.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), testcase.ExpectResponse)
	}

}

func TestUpdateBook(t *testing.T) {
	reqStr := `{
		"ID":        1,
		"title":     "Laskar Pelangi",
		"author":    "Giring",
		"publisher": "Mizan"
		}`
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"
	defer func() {
		successTest := TestCase{
			Method:         http.MethodPut,
			Name:           "Update book",
			Path:           "/",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success Edit Book",
		}

		e := InitEchoTestAPI()
		InsertBook()

		req := httptest.NewRequest(successTest.Method, successTest.Path, strings.NewReader(reqStr))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.UpdateBookController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)
		}
	}()

	failedTest := TestCase{
		Method:         http.MethodPut,
		Name:           "Failed update book",
		Path:           "/",
		ExpectStatus:   http.StatusBadRequest,
		ExpectResponse: "Failed update data",
	}

	e := InitEchoTestAPI()
	InsertBook()

	req := httptest.NewRequest(failedTest.Method, failedTest.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.SetPath(failedTest.Path)
	if assert.NoError(t, controller.UpdateBookController(c)) {
		assert.Equal(t, failedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failedTest.ExpectResponse)
	}

}

func TestDeleteBook(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWx0YSIsInVzZXJJZCI6MX0.wmnj3egcsfjvyPQ1QZFL4wZIluhMDQoADsz85Sx18cQ"
	defer func() {
		successTest := TestCase{

			Method:         http.MethodDelete,
			Name:           "Delete book",
			Path:           "/",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success Delete Book",
		}

		e := InitEchoTestAPI()
		InsertBook()

		req := httptest.NewRequest(successTest.Method, successTest.Path, nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		c.SetPath(successTest.Path)
		if assert.NoError(t, controller.DeleteBookController(c)) {
			assert.Equal(t, successTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), successTest.ExpectResponse)
		}
	}()

	failedTest := TestCase{

		Method:         http.MethodDelete,
		Name:           "Failed delete book",
		Path:           "/",
		ExpectStatus:   http.StatusBadRequest,
		ExpectResponse: "Failed delete book",
	}

	e := InitEchoTestAPI()
	InsertBook()

	req := httptest.NewRequest(failedTest.Method, failedTest.Path, nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.SetPath(failedTest.Path)
	if assert.NoError(t, controller.DeleteBookController(c)) {
		assert.Equal(t, failedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), failedTest.ExpectResponse)
	}
}
package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

//Test for FizzBuzz-2
func TestShouldReturnStatusOK(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:data")
	//set data type
	c.SetParamNames("data")
	c.SetParamValues("string")
	//set URL query
	q := req.URL.Query()
	q.Add("multiple1", "3")
	q.Add("multiple2", "5")
	q.Add("limit", "20")
	q.Add("str1", "s1")
	q.Add("str2", "s2")
	req.URL.RawQuery = q.Encode()

	// Assertions
	if assert.NoError(t, GetFizzBuzz(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "1,2,s1,4,s2,s1,7,8,s1,s2,11,s1,13,14,s1s2,16,17,s1,19,s2", rec.Body.String())
	}
}

//Test for FizzBuzz-2
func TestShouldReturnStatusBadRequest_NotInt(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:data")
	//set data type
	c.SetParamNames("data")
	c.SetParamValues("string")
	//set URL query
	q := req.URL.Query()
	q.Add("multiple1", "aa")
	q.Add("multiple2", "bb")
	q.Add("limit", "cc")
	q.Add("str1", "s1")
	q.Add("str2", "s2")
	req.URL.RawQuery = q.Encode()

	// Assertions
	if assert.NoError(t, GetFizzBuzz(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "{\"error\":\"params multiple1, multiple2 and limit must be int : you entered limit = cc, multiple1 = aa, multiple2 = bb\"}\n", rec.Body.String())
	}
}

//Test for FizzBuzz-2
func TestShouldReturnStatusBadRequest_WrongDataType(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:data")
	//set data type
	c.SetParamNames("data")
	c.SetParamValues("json")
	//set URL query
	q := req.URL.Query()
	q.Add("multiple1", "1")
	q.Add("multiple2", "2")
	q.Add("limit", "50")
	q.Add("str1", "s1")
	q.Add("str2", "s2")
	req.URL.RawQuery = q.Encode()

	// Assertions
	if assert.NoError(t, GetFizzBuzz(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "{\"error\":\"please specify the data type as string : you entered json\"}\n", rec.Body.String())
	}
}

//Test for FizzBuzz-4
func TestShouldReturnStatusBadRequest_IncorrectParameter(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:data")
	//set data type
	c.SetParamNames("data")
	c.SetParamValues("string")
	//set URL query
	q := req.URL.Query()
	q.Add("incorrectParam", "1")
	q.Add("multiple2", "2")
	q.Add("limit", "50")
	q.Add("str1", "s1")
	q.Add("str2", "s2")
	req.URL.RawQuery = q.Encode()

	// Assertions
	if assert.NoError(t, GetFizzBuzz(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "{\"error\":\"parameters expected are : [limit multiple1 multiple2 str1 str2]\"}\n", rec.Body.String())
	}
}

//Test for FizzBuzz-4
func TestShouldReturnStatusBadRequest_IncorrectNumberOfParameters(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:data")
	//set data type
	c.SetParamNames("data")
	c.SetParamValues("string")
	//set URL query
	q := req.URL.Query()
	q.Add("multiple2", "2")
	q.Add("limit", "50")
	q.Add("str1", "s1")
	q.Add("str2", "s2")
	req.URL.RawQuery = q.Encode()

	// Assertions
	if assert.NoError(t, GetFizzBuzz(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "{\"error\":\"5 parameters expected : [limit multiple1 multiple2 str1 str2]\"}\n", rec.Body.String())
	}
}

//Test for FizzBuzz-3
func TestShouldReturnStatusBadRequest_IntInferiorToOne(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:data")
	//set data type
	c.SetParamNames("data")
	c.SetParamValues("string")
	//set URL query
	q := req.URL.Query()
	q.Add("multiple1", "-1")
	q.Add("multiple2", "2")
	q.Add("limit", "50")
	q.Add("str1", "s1")
	q.Add("str2", "s2")
	req.URL.RawQuery = q.Encode()

	// Assertions
	if assert.NoError(t, GetFizzBuzz(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "{\"error\":\"limit and multiples can't be inferior to 1 : limit = 50, multiple1 = -1, multiple2 = 2\"}\n", rec.Body.String())
	}
}

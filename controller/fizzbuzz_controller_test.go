package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

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
	q.Add("multiple1", "1")
	q.Add("multiple2", "2")
	q.Add("limit", "50")
	q.Add("str1", "s1")
	q.Add("str2", "s2")
	req.URL.RawQuery = q.Encode()

	// Assertions
	if assert.NoError(t, GetFizzBuzz(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "fizzbuzz:replace multiples of 1 and 2 by s1 and s2, starting from 1 to 50", rec.Body.String())
	}
}

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
		assert.Equal(t, "{\"error\":\"params multiple1, multiple2 and limit must be int\"}\n", rec.Body.String())
	}
}

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
		assert.Equal(t, "{\"error\":\"please specify the data type as string\"}\n", rec.Body.String())
	}
}

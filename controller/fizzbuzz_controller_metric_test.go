package controller

import (
	"FizzBuzz/metric"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

//Test for FizzBuzz-15
func TestShouldReturnStatusOK_MostUsedRequest(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:data")
	//set data type
	c.SetParamNames("data")
	c.SetParamValues("fizzbuzz")
	//set URL query
	q := req.URL.Query()
	q.Add("multiple1", "6")
	q.Add("multiple2", "4")
	q.Add("limit", "50")
	q.Add("str1", "buzz")
	q.Add("str2", "light")
	req.URL.RawQuery = q.Encode()

	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)
	c2.SetPath("/fizzbuzz/:data")
	//set data type
	c2.SetParamNames("data")
	c2.SetParamValues("fizzbuzz")
	//set URL query
	q2 := req2.URL.Query()
	q2.Add("multiple1", "3")
	q2.Add("multiple2", "11")
	q2.Add("limit", "45")
	q2.Add("str1", "s1")
	q2.Add("str2", "s2")
	req2.URL.RawQuery = q2.Encode()

	// Setup
	e3 := echo.New()
	req3 := httptest.NewRequest(http.MethodGet, "/", nil)
	rec3 := httptest.NewRecorder()
	c3 := e3.NewContext(req3, rec3)
	c3.SetPath("/statistics")

	// Assertions
	GetFizzBuzz(c)
	GetFizzBuzz(c)
	GetFizzBuzz(c)

	if assert.NoError(t, metric.GetStatisticsFizzBuzz(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
		assert.Equal(t, "Most used request is : limit=50, multiple1=6, multiple2=4, str1=buzz, str2=light\nThe request was asked 3 times",
			rec3.Body.String())
	}
	rec3.Body.Reset()

	GetFizzBuzz(c2)
	GetFizzBuzz(c2)
	GetFizzBuzz(c2)

	if assert.NoError(t, metric.GetStatisticsFizzBuzz(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
		assert.Equal(t, "Most used request is : limit=50, multiple1=6, multiple2=4, str1=buzz, str2=light\nThe request was asked 3 times",
			rec3.Body.String())
	}

	rec3.Body.Reset()

	GetFizzBuzz(c2)

	if assert.NoError(t, metric.GetStatisticsFizzBuzz(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
		assert.Equal(t, "Most used request is : limit=45, multiple1=3, multiple2=11, str1=s1, str2=s2\nThe request was asked 4 times",
			rec3.Body.String())
	}
	metric.ResetFizzBuzzRequests()
}

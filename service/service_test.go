package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func SimpleTestGetFizzBuzz(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/fizzbuzz/:data")
	c.SetParamNames("multiple1", "multiple2", "limit", "str1", "str2")
	c.SetParamValues("0", "0", "0", "s1", "s2")

	// Assertions
	if assert.NoError(t, GetFizzBuzz(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "fizzbuzz:replace multiples of 0 and 0 by s1 and s2, starting from 1 to 0", rec.Body.String())
	}
}

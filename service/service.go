package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//GET API which return the name of the cats specified in QueryParam
//http://localhost:8000/fizzbuzz/json?multiple1=3&multiple2=5&limit=20&str1=fizz&str2=buzz
func GetFizzBuzz(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	multiple1, _ := strconv.Atoi(c.QueryParam("multiple1"))
	multiple2, _ := strconv.Atoi(c.QueryParam("multiple2"))
	str1 := c.QueryParam("str1")
	str2 := c.QueryParam("str2")

	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("fizzbuzz:replace multiples of %d and %d by %s and %s, starting from 1 to %d", multiple1, multiple2, str1, str2, limit))
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as String"})
	}
}

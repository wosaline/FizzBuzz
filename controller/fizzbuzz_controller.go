package controller

import (
	"FizzBuzz/service"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo"
)

func bValidParameters(params url.Values) (bool, error) {
	listOfParameters := service.GetListOfParameters()

	if len(params) != service.NbFizzBuzzParameters {
		return false, fmt.Errorf("%d parameters expected : %v", service.NbFizzBuzzParameters, listOfParameters)
	}
	isPresent := true

	for _, str := range listOfParameters {
		_, found := params[str]
		isPresent = isPresent && found
	}

	if isPresent {
		return true, nil
	} else {
		return false, fmt.Errorf("parameters expected are : %v", listOfParameters)
	}
}

//GET API which return the name of the cats specified in QueryParam
//http://localhost:8000/fizzbuzz/json?multiple1=3&multiple2=5&limit=20&str1=fizz&str2=buzz
func GetFizzBuzz(c echo.Context) error {
	isValid, errInvalidParameter := bValidParameters(c.QueryParams())
	if !isValid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errInvalidParameter.Error()})
	}

	limit, errLimit := strconv.Atoi(c.QueryParam("limit"))
	multiple1, errMultiple1 := strconv.Atoi(c.QueryParam("multiple1"))
	multiple2, errMultiple2 := strconv.Atoi(c.QueryParam("multiple2"))
	str1 := c.QueryParam("str1")
	str2 := c.QueryParam("str2")
	dataType := c.Param("data")

	if errLimit != nil || errMultiple1 != nil || errMultiple2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "params multiple1, multiple2 and limit must be int"})
	}
	if dataType != "string" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "please specify the data type as string"})
	}

	result, errFizzBuzz := service.FizzBuzz(limit, multiple1, multiple2, str1, str2)
	if errFizzBuzz != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errFizzBuzz.Error()})
	}
	return c.String(http.StatusOK, fmt.Sprint(result))
}

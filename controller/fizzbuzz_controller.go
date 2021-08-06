package controller

import (
	"FizzBuzz/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var fizzBuzzMaps map[service.FizzBuzzStructure]int = make(map[service.FizzBuzzStructure]int)

func parseValidParameters(c echo.Context) (int, int, int, string, string, error) {
	listOfParameters := service.GetListOfParameters()
	params := c.QueryParams()

	if len(params) != service.NbFizzBuzzParameters {
		return 0, 0, 0, "", "", fmt.Errorf("%d parameters expected : %v", service.NbFizzBuzzParameters, listOfParameters)
	}
	isPresent := true

	for _, str := range listOfParameters {
		_, found := params[str]
		isPresent = isPresent && found
	}

	if !isPresent {
		return 0, 0, 0, "", "", fmt.Errorf("parameters expected are : %v", listOfParameters)
	}

	limit, errLimit := strconv.Atoi(c.QueryParam("limit"))
	multiple1, errMultiple1 := strconv.Atoi(c.QueryParam("multiple1"))
	multiple2, errMultiple2 := strconv.Atoi(c.QueryParam("multiple2"))
	str1 := c.QueryParam("str1")
	str2 := c.QueryParam("str2")
	dataType := c.Param("data")

	if errLimit != nil || errMultiple1 != nil || errMultiple2 != nil {
		return limit, multiple1, multiple2, str1, str2, fmt.Errorf("params multiple1, multiple2 and limit must be int : you entered limit = %s, multiple1 = %s, multiple2 = %s", c.QueryParam("limit"), c.QueryParam("multiple1"), c.QueryParam("multiple2"))
	}
	if dataType != "string" {
		return limit, multiple1, multiple2, str1, str2, fmt.Errorf("please specify the data type as string : you entered %s", dataType)
	}

	return limit, multiple1, multiple2, str1, str2, nil
}

//GET API which returns the result of fizzbuzz for the params specified in the query
//http://localhost:8000/fizzbuzz/json?multiple1=3&multiple2=5&limit=20&str1=fizz&str2=buzz
func GetFizzBuzz(c echo.Context) error {
	limit, multiple1, multiple2, str1, str2, errInvalidParameter := parseValidParameters(c)
	if errInvalidParameter != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errInvalidParameter.Error()})
	}

	result, errFizzBuzz := service.FizzBuzz(limit, multiple1, multiple2, str1, str2)
	if errFizzBuzz != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": errFizzBuzz.Error()})
	}

	calculateStatistics(service.FizzBuzzStructure{
		Limit:     limit,
		Multiple1: multiple1,
		Multiple2: multiple2,
		Str1:      str1,
		Str2:      str2,
	})

	return c.String(http.StatusOK, fmt.Sprint(result))
}

func calculateStatistics(FBStruct service.FizzBuzzStructure) {
	fizzBuzzMaps[FBStruct]++
}

//used for tests
func getFizzBuzzMapsCount(FBStruct service.FizzBuzzStructure) int {
	return fizzBuzzMaps[FBStruct]
}

//Return the parameters corresponding to the most used request, as well as the number of hits for this request
func getMostUsedRequest() (service.FizzBuzzStructure, int) {
	if len(fizzBuzzMaps) < 1 {
		return service.FizzBuzzStructure{}, -1
	}
	max := 0
	var fbMostUsed service.FizzBuzzStructure
	for fb, fbCount := range fizzBuzzMaps {
		if fbCount > max {
			max = fbCount
			fbMostUsed = fb
		}
	}
	return fbMostUsed, max
}

//GET API which returns the parameters corresponding to the most used request
//as well as the number of hits for this request
//http://localhost:8000/statistics
func GetStatisticsFizzBuzz(c echo.Context) error {
	fbMostUsed, count := getMostUsedRequest()
	if count == -1 {
		return c.String(http.StatusOK, "No request has been made yet")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Most used request is : limit=%d, multiple1=%d, multiple2=%d, str1=%s, str2=%s\nThe request was asked %d times",
		fbMostUsed.Limit, fbMostUsed.Multiple1, fbMostUsed.Multiple2, fbMostUsed.Str1, fbMostUsed.Str2, count))
}

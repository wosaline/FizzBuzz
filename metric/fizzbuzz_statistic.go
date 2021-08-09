package metric

import (
	"FizzBuzz/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var FizzBuzzRequestsMap map[service.FizzBuzzStructure]int = make(map[service.FizzBuzzStructure]int)

func IncrementFizzBuzzRequestCount(FBStruct service.FizzBuzzStructure) {
	FizzBuzzRequestsMap[FBStruct]++
}

//used for tests
func getFizzBuzzMapsCount(FBStruct service.FizzBuzzStructure) int {
	return FizzBuzzRequestsMap[FBStruct]
}

//Return the parameters corresponding to the most used request, as well as the number of hits for this request
func getMostUsedRequest() (service.FizzBuzzStructure, int) {
	if len(FizzBuzzRequestsMap) < 1 {
		return service.FizzBuzzStructure{}, -1
	}
	max := 0
	var fbMostUsed service.FizzBuzzStructure
	for fb, fbCount := range FizzBuzzRequestsMap {
		if fbCount > max {
			max = fbCount
			fbMostUsed = fb
		}
	}
	return fbMostUsed, max
}

//reset the FizzBuzzRequests map the tests purpose
func ResetFizzBuzzRequests() {
	FizzBuzzRequestsMap = make(map[service.FizzBuzzStructure]int)
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

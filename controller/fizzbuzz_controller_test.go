package controller

import (
	"FizzBuzz/service"
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

//Test for FizzBuzz-11
func TestShouldReturnStatusOK_Statistics(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/statistics")

	// Assertions
	if assert.NoError(t, GetStatisticsFizzBuzz(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "No request has been made yet", rec.Body.String())
	}
}

//Test for FizzBuzz-12
func TestShouldBeOK_ReturnCountsForRequestStatistics(t *testing.T) {
	type args struct {
		limit     int
		multiple1 int
		multiple2 int
		str1      string
		str2      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Simple test OK", args{20, 2, 3, "fizz", "buzz"}, 1},
		{"Simple test OK2", args{20, 2, 3, "fizz", "buzz"}, 2},
		{"Simple test OK3", args{20, 2, 3, "hoy", "buzz"}, 1},
		{"Simple test OK4", args{20, 2, 3, "hoy", "buzz"}, 2},
		{"Simple test OK5", args{20, 2, 3, "fizz", "buzz"}, 3},
		{"Simple test OK6", args{25, 2, 3, "fizz", "buzz"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FBStruct := service.FizzBuzzStructure{
				Limit:     tt.args.limit,
				Multiple1: tt.args.multiple1,
				Multiple2: tt.args.multiple2,
				Str1:      tt.args.str1,
				Str2:      tt.args.str2,
			}
			calculateStatistics(FBStruct)
			count := getFizzBuzzMapsCount(FBStruct)
			if count != tt.want {
				t.Errorf("calculateStatistics() got = %v, want %v", count, tt.want)
				return
			}
		})
	}
}

//Test for FizzBuzz-15
func TestShouldBeOK_ReturnMostUsedRequestStatistics(t *testing.T) {
	type args struct {
		limit     int
		multiple1 int
		multiple2 int
		str1      string
		str2      string
	}
	tests := []struct {
		name             string
		args             args
		fbMostUsedWanted service.FizzBuzzStructure
		countWanted      int
	}{
		{"Simple test OK", args{20, 2, 3, "fizz", "buzz"}, service.FizzBuzzStructure{Limit: 20, Multiple1: 2, Multiple2: 3, Str1: "fizz", Str2: "buzz"}, 1},
		{"Simple test OK2", args{20, 2, 3, "fizz", "buzz"}, service.FizzBuzzStructure{Limit: 20, Multiple1: 2, Multiple2: 3, Str1: "fizz", Str2: "buzz"}, 2},
		{"Simple test OK3", args{20, 2, 3, "hoy", "buzz"}, service.FizzBuzzStructure{Limit: 20, Multiple1: 2, Multiple2: 3, Str1: "fizz", Str2: "buzz"}, 2},
		{"Simple test OK4", args{20, 2, 3, "hoy", "buzz"}, service.FizzBuzzStructure{Limit: 20, Multiple1: 2, Multiple2: 3, Str1: "fizz", Str2: "buzz"}, 2},
		{"Simple test OK5", args{20, 2, 3, "hoy", "buzz"}, service.FizzBuzzStructure{Limit: 20, Multiple1: 2, Multiple2: 3, Str1: "hoy", Str2: "buzz"}, 3},
		{"Simple test OK6", args{25, 2, 3, "fizz", "buzz"}, service.FizzBuzzStructure{Limit: 20, Multiple1: 2, Multiple2: 3, Str1: "hoy", Str2: "buzz"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FBStruct := service.FizzBuzzStructure{
				Limit:     tt.args.limit,
				Multiple1: tt.args.multiple1,
				Multiple2: tt.args.multiple2,
				Str1:      tt.args.str1,
				Str2:      tt.args.str2,
			}
			calculateStatistics(FBStruct)
			fbMostUsed, count := getMostUsedRequest()
			if count != tt.countWanted {
				t.Errorf("GetMostUsedRequest() got = %v, want %v for the count", count, tt.countWanted)
				return
			} else if fbMostUsed != tt.fbMostUsedWanted {
				t.Errorf("GetMostUsedRequest() got = Limit : %d, Multiple1 : %d, Multiple2 : %d, Str1 : %s, Str2 : %s \n want Limit : %d, Multiple1 : %d, Multiple2 : %d, Str1 : %s, Str2 : %s for the request",
					fbMostUsed.Limit, fbMostUsed.Multiple1, fbMostUsed.Multiple2, fbMostUsed.Str1, fbMostUsed.Str2,
					tt.fbMostUsedWanted.Limit, tt.fbMostUsedWanted.Multiple1, tt.fbMostUsedWanted.Multiple2, tt.fbMostUsedWanted.Str1, tt.fbMostUsedWanted.Str2)
				return
			}
		})
	}
}

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
	c.SetParamValues("string")
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
	c2.SetParamValues("string")
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

	if assert.NoError(t, GetStatisticsFizzBuzz(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
		assert.Equal(t, "Most used request is : limit=50, multiple1=6, multiple2=4, str1=buzz, str2=light\nThe request was asked 3 times",
			rec3.Body.String())
	}
	rec3.Body.Reset()

	GetFizzBuzz(c2)
	GetFizzBuzz(c2)
	GetFizzBuzz(c2)

	if assert.NoError(t, GetStatisticsFizzBuzz(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
		assert.Equal(t, "Most used request is : limit=50, multiple1=6, multiple2=4, str1=buzz, str2=light\nThe request was asked 3 times",
			rec3.Body.String())
	}

	rec3.Body.Reset()

	GetFizzBuzz(c2)

	if assert.NoError(t, GetStatisticsFizzBuzz(c3)) {
		assert.Equal(t, http.StatusOK, rec3.Code)
		assert.Equal(t, "Most used request is : limit=45, multiple1=3, multiple2=11, str1=s1, str2=s2\nThe request was asked 4 times",
			rec3.Body.String())
	}
}

package metric

import (
	"FizzBuzz/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

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
			IncrementFizzBuzzRequestCount(FBStruct)
			count := getFizzBuzzMapsCount(FBStruct)
			if count != tt.want {
				t.Errorf("IncrementFizzBuzzRequestCount() got = %v, want %v", count, tt.want)
				return
			}
		})
	}
	ResetFizzBuzzRequests()
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
			IncrementFizzBuzzRequestCount(FBStruct)
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
	ResetFizzBuzzRequests()
}

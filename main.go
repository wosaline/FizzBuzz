package main

import (
	"FizzBuzz/controller"

	"github.com/labstack/echo"
)

//var fizzbuzzes []entity.FizzBuzzStructure

func main() {
	/*fizzbuzzes = []entity.FizzBuzzStructure{
		{Limit: 30, Multiple1: 3, Multiple2: 5, Str1: "fizz", Str2: "buzz"},
		{Limit: 10, Multiple1: 2, Multiple2: 3, Str1: "fizz", Str2: "buzz"},
	}*/

	e := echo.New()
	//QueryParam along with PathVariable Get API
	e.GET("/fizzbuzz/:data", controller.GetFizzBuzz)
	e.GET("/statistics", controller.GetStatisticsFizzBuzz)

	//starts server en port 8000
	e.Logger.Fatal(e.Start(":8000"))
}

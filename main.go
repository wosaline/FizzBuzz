package main

import (
	"FizzBuzz/controller"
	"FizzBuzz/metric"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	//QueryParam along with PathVariable Get API
	e.GET("/fizzbuzz/:data", controller.GetFizzBuzz)
	e.GET("/statistics", metric.GetStatisticsFizzBuzz)

	//starts server en port 8000
	e.Logger.Fatal(e.Start(":8000"))
}

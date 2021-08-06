package main

import (
	"FizzBuzz/controller"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	//QueryParam along with PathVariable Get API
	e.GET("/fizzbuzz/:data", controller.GetFizzBuzz)
	e.GET("/statistics", controller.GetStatisticsFizzBuzz)

	//starts server en port 8000
	e.Logger.Fatal(e.Start(":8000"))
}

package main

import(
	"fmt"
	"FizzBuzz/entity"
)


func main() {
	res, _ := entity.FizzBuzz(6,7,10,"fizz","buzz")
	fmt.Println(res)
	//fmt.Println(err)
}
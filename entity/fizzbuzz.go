 //three integers int1, int2 and limit, and two strings str1and str2

 package entity

 import(
 	"fmt"
 )

//FizzBuzzStructure contains all the parameters for FizzBuzz function
type FizzBuzzStructure struct {
	limit 			int    	`json:"limit"`
	multiple1		int  	`json:"multiple1"`
	multiple2       int  	`json:"multiple2"`
	str1       		string  `json:"str1"`
	str2        	string 	`json:"str2"`
}

 //FizzBuzz replaces :
 //- all multiples of multiple1 with str1
 //- all multiples of multiple2 with str2
 //- all multiples of multiple1 and multiple2 with str1str2
 //Range of numbers goes from 1 to limit
func FizzBuzz(limit,multiple1,multiple2 int,str1,str2 string) (string,error) {
	var arrFizzBuzz=""
	var i=0
	var temp=""

	if limit<1 || multiple1<1 || multiple2<1{
		return arrFizzBuzz,fmt.Errorf("Limit and multiples can't be inferior to 1 : limit = %d, multiple1 = %d, multiple2 = %d", limit, multiple1, multiple2)
	}
	
	for from:=1;from<=limit;from++ {
		temp=""
		if from%multiple1==0 {
			temp+=str1
		}
		if from%multiple2==0 {
			temp+=str2
		}
		if len(temp)==0{
			temp=fmt.Sprintf("%d", from)
		}
		arrFizzBuzz+=temp
		if from<limit{
			arrFizzBuzz+=","
		}
		i++
	}

	return arrFizzBuzz,nil
}
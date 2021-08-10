# FizzBuzz
FizzBuzz/GoLang

## Prerequisite
To compile this project, you need to have :
- `go` installed

## Commands to run

Before you run the code, make sure to create go.mod and go.sum files to track the code's dependencies :
```sh
go mod init FizzBuzz
go get github.com/labstack/echo
```

Run the code
```sh
go run main.go
```

## API
`FizzBuzz` has two possible routes
<br>

### **/fizzbuzz/string?fizzbuzz**
![](https://img.shields.io/badge/Request-GET-green.svg?style=flat)
```http
http://localhost:8000/fizzbuzz/fizzbuzz?multiple1=<multiple1>&multiple2=<multiple2>&limit=<limit>&str1=<str1>&str2=<str2>
```

| Field | Description
| ----- | -----------
| multiple1 | first multiple which will be replace by <str1>
| multiple2 | second multiple which will be replace by <str2>
| limit | limit of the fizzbuzz
| str1 | word to replace number multiple of <multiple1>
| str2 | word to replace number multiple of <multiple2>

#### Results

- **200 OK**

```http
http://localhost:8000/fizzbuzz/fizzbuzz?multiple1=3&multiple2=5&limit=50&str1=fizz&str2=buzz
```
```string
1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz
```
<br>

- **400 BAD REQUEST**
  
*There are several types of bad requests*
<br>
  
1. Multiples or limit inferior to 1
```http
http//localhost:8000/fizzbuzz/fizzbuzz?multiple1=-1&multiple2=5&limit=50&str1=fizz&str2=buzz
```
```json
{"error":"limit and multiples can't be inferior to 1 : limit = 50, multiple1 = -1, multiple2 = 5"}
```
<br>

2. At least one wrong parameter name
```http
http://localhost:8000/fizzbuzz/fizzbuzz?notaparameter=3&multiple2=5&limit=50&str1=fizz&str2=buzz
```
```json
{"error":"parameters expected are : [limit multiple1 multiple2 str1 str2]"}
```
<br>

3. Wrong data type
```http
http://localhost:8000/fizzbuzz/json?multiple1=3&multiple2=5&limit=50&str1=fizz&str2=buzz
```
```json
{"error":"please specify the data type as fizzbuzz : you entered json"}
```
<br>

4. At least one parameter missing
```http
http://localhost:8000/fizzbuzz/fizzbuzz?multiple2=5&limit=50&str1=fizz&str2=buzz
```
```json
{"error":"5 parameters expected : [limit multiple1 multiple2 str1 str2]"}
```
<br>

5. Multiple or limit is not an int
```http
http://localhost:8000/fizzbuzz/fizzbuzz?multiple1=aa&multiple2=5&limit=50&str1=fizz&str2=buzz
```
```json
{"error":"params multiple1, multiple2 and limit must be int : you entered limit = 50, multiple1 = aa, multiple2 = 5"}
```
<br>

### **/statistics**
![](https://img.shields.io/badge/Request-GET-green.svg?style=flat)
```http
http://localhost:8000/statistics
```

#### Results

- **200 OK**

*There are two possible results*
<br>
1. No fizzbuzz request before  

```string
No request has been made yet
```
<br>

2. At least one request was made
```string
Most used request is : limit=50, multiple1=3, multiple2=5, str1=fizz, str2=buzz
The request was asked 2 times
```

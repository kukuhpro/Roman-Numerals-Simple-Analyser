# Roman Numerals Simple Analyser

### Requirements
* Golang 1.10

### Get Started
1. Copy this project directory to `$GOPATH/src/roman`
2. Running script 
```golang
go run main.go
```
3. Input your text
```
2018/11/28 21:57:33 Running Auto Migration on sqlite...
2018/11/28 21:57:33 Running on grpc server....
2018/11/28 21:57:33 Connect GRPC to port :3214
// input your text here
```

### Running Test 
* Client Part 
```golang
go test -v ./client_test
```

### System Design 
System Design split into 2 parts 
* Client will handle for parsing text into token analyser, this token analyzer will used for check, if text is accepted then the process will go on, if not will return error.
* Server will get Token Analyzer from client, and used that token to define what kind of process that will do. In in this part also handle for database connection for write, read data. 

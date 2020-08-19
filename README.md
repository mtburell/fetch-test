# Fetch Rewards Homework Assignment
[![codecov](https://codecov.io/gh/mtburell/fetch-test/branch/master/graph/badge.svg?token=Z8VXG09KHW)](https://codecov.io/gh/mtburell/fetch-test)

#### Problem statement: 
Accept a list of email addresses and return an integer indicating the number of unique email addresses. Where "unique" email addresses means they will be delivered to the same account using Gmail account matching. Specifically: Gmail will ignore the placement of "." in the username. And it will ignore any portion of the username after a "+".

Examples:
test.email@gmail.com, test.email+spam@gmail.com and testemail@gmail.com will all go to the same address, and thus the result should be 1.
For any requirements not specified via an example, use your best judgement to determine expected result.

## Technology Summary

This project includes a webserver written in Go. For more information about installing Go, click [here](https://golang.org/doc/install).


## Installing and Running Locally

#### To Run: 
Pull this project, navigate to the root directory, and run 
```
go run cmd/main.go
```

#### To Build: 
Pull this project, navigate to the root directory, and run
```
go build cmd/main.go
```

# README

One day build of the Unix `seq` command. `seq` prints a sequence
of numbers. This program does not support all the features in 
`seq`. The only flag supported is the `-w`, however, it does not 
support right zero-padding after decimal place.

Program was created for demonstration purposes.

Usage:

```go
// Example 1:
go run main.go 5

// Example 2:
go run main.go 1 5

// Example 3:
go run main.go 1 .5 2

// Example 4:
go run main.go -w 1 2 100
```

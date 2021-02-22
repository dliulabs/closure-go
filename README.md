# closure-demo

Basic closure: a func() that returns a func(), it encloses a set of data and a func() to iterates through the set of data.

```
cd closure-demo
go mod init
go mod tidy
go run closure.go
```

# Isolating Data using Closure

Calculating Fibonacci for example, requires maintaining states so that number sequence can be returned in order: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

```
cd isolating-data
go mod init
go mod tidy
go fmt
go run fibonacci.go
```

# Wrap Middleware

using closure to wrap a middleware

# Passing private dataset

Using closure to pass private dataset to a function
# Refactoring a CPF validator in Go

This content is inspired in part by Branas.io's Clean Code and Clean Architecture course.

For more information access:

[clean code e clean architecture](https://app.branas.io/clean-code-e-clean-architecture)

The CPF validator originally implemented in TypeScript is available [here](https://github.com/rodrigobranas/cccat6_refactoring/tree/master/src/example2)

I re-implemented this validator in Go, as well as the unit tests. To run the tests, with test coverage measurement:

```$ go test -v -coverprofile=coverage.out ./...```

View code coverage in a web browser:

```$ go tool cover -html=coverage.out```

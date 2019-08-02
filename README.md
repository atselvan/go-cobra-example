# Go CLI using cobra module

An example application build using the [cobra](https://github.com/spf13/cobra) module

## Installing

```go
go get github.com/spf13/cobra/cobra
```

## Initializing a project

```go
$GOPATH/bin/cobra init -l none -a atselvan --pkg-name src/com/privatesquare/go/go-cobra-example

```

```-l none``` : By default cobra will add the Apache license to you go files, if you do not want a default license to 
be added you can this flag can be used

```-a author``` : Adding an author name to the copy right

_Both these flags are optional_

```--pkg-name com/privatesquare/go/go-cobra-example``` : The full package name of your application

```Output:```

```go
Your Cobra applicaton is ready at
/Users/allanselvan/workspace/go/go-cobra-example/src/com/privatesquare/go/go-cobra-example
```

This would create a main.go file and a root.cmd file

Adding a subcommand

```go
$GOPATH/bin/cobra add add -l none -a atselvan

```

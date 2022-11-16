# What to have learned

- [ ] How to operate the command line args in golang
- [ ] How to test the command line args in golang


## How to operate the command line args in golang

We use the `flag` package to operate the command line args in golang.


## How to test the command line args in golang

We use the `testing` package to test the command line args in golang.

`flag.Commandline.Set()` must be executed before `flag.Parse()`.

<br>

- Reference:
  - [【Go】 ユニットテストでflagへ引数を渡す際のハマりどころ](https://qiita.com/vengavengavnega/items/874212b929ba53ce2810)
  - [Testing flag parsing in Go programs](https://eli.thegreenplace.net/2020/testing-flag-parsing-in-go-programs/)


## How to use the dynamic type value for CLI args

- Reference: https://pkg.go.dev/flag#Value

We can use the `flag.Value` interface to define the dynamic type value for CLI args.

> Value is the interface to the dynamic value stored in a flag. (The default value is represented as a string.)

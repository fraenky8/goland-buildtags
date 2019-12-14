# Build tags

Used as bug report here: https://youtrack.jetbrains.com/issue/GO-8609

Running only the unit tests:

```
go test -v ./...
=== RUN   TestFoo_Do
foo->Do()
--- PASS: TestFoo_Do (0.00s)
=== RUN   TestFoo_Greet
hello Goland
--- PASS: TestFoo_Greet (0.00s)
PASS
ok  	github.com/fraenky8/goland-buildtags	0.007s
```

Running only the integration tests:

```
go test -v -tags=integration ./...
=== RUN   TestFoo_Integration
--- PASS: TestFoo_Integration (0.00s)
    foo_integration_test.go:10: startingFoo integration tests
PASS
ok  	github.com/fraenky8/goland-buildtags
```



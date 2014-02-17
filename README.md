## url-router

Sinatra-like URL router for Golang.

## Install

```bash
$ go get github.com/azer/url-router
```

## Usage

```go
routes = urlrouter.New("/users/:name", "/page/:page")

match := routes.Match("/users/john")

match
// => urlrouter.Match {
//      params: { name: "john" }
//      pattern: "/users/:name"
//    }

match = routes.Match("nonexisting")

match
// => nil
```

See `test/test.go` for more info.

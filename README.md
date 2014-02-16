## url-router

Sinatra-like URL router for Golang.

## Install

```bash
$ go get github.com/azer/url-router
```

## Usage

```golang
routes = urlrouter.New("/users/:name", "/page/:page")

match := routes.Match("/users/john?foo=bar&qux=corge")

match
// => urlrouter.Match {
//      params: { name: "john" }
//      pattern: "/users/:name"
//    }

match = routes.Match("nonexisting")

match
// => nil
```

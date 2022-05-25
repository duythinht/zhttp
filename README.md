## ZHTTP

The small lib to convert generic Handler to net/http.HandleFunc

zhttp is designed to be the simplest way to declare a http.HandlerFunc, which have a lot of boilerplate code, you can use it with net/http, chi-router or gorilla mux

## Install

`go get -u github.com/duythinht/zhttp`

### Features

* Convert generic `func(context.Context, *Input) (*Output, error)` to `net/http.HandlerFunc`
* Handle Error with statusCode by generic interface `zhttp.Error`
* Some helper function to declare common http error

### Example

```go
package main

import (
	"net/http"
)


// HelloRequest take input from body by those methods:
//      * url query, eg: /?name=abc-xyz
//      * POST body, default is application/json
//      * Bind(r *http.Request) error method (Boundable interface), which you can override
type HelloRequest struct {
    Name string `url:"name"`
}

// HelloResponse can be marshal to json
// you can define
// func(out *HelloResponse) Marshal(context.Context) ([]byte, error)
// to customize the response
type HelloResponse struct {
    Message string
}

func Hello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
    return &HelloResponse{
        Message: fmt.Sprintf("Hello, %s!").
    }, nil
}

func main() {
    http.Handle("/", zhttp.Handler(Hello))
    http.ListenAndServe(":3000", r)
}
// curl http://localhost:3000/?name=World
// {"message": "Hello, World!"}
```

### Why I made it?

Because I'm tired of typing a lot of boilerplate if else, w.WriteHeader and w.Write...  you will cost a bit about performance (for the raw hello world without handle error cases, you cost for 40-45% time slower because underly it check the error for you), but you don't need write a lot of boilerplate code everywhere

### Is it any good?

May be!

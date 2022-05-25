package zhttp

import (
	"context"
	"net/http"
)

// Marshalable is interface, that override marshal output type
type Marshalable interface {
	Marshal(ctx context.Context) ([]byte, error)
}

// Boundable interface hold Bind method, which can binding input type from http.Request
type Boundable interface {
	//Bind value from request
	Bind(r *http.Request) error
}

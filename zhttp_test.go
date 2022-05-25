package zhttp_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/duythinht/zhttp"
)

func TestWrapHandlerFunction(t *testing.T) {

	type input struct {
		Name string `url:"name"`
	}

	type output struct {
	}

	name := "hello"

	h := zhttp.Handler(func(_ context.Context, in *input) (out *output, err error) {
		if in.Name != name {
			t.Logf("expected: %s, got: %s", name, in.Name)
			t.Fail()
		}
		return nil, nil
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/?name=hello", nil)

	h.ServeHTTP(w, r)
}

func BenchmarkTypeAssertion(b *testing.B) {

	errMock := errors.New("mock")

	e := zhttp.BadRequest(errMock)

	for i := 0; i < b.N; i++ {
		if httpErr, ok := any(e).(zhttp.Error); ok {
			httpErr.HTTPError()
		}
	}
}

func BenchmarkSwitchTypeAssertion(b *testing.B) {

	errMock := errors.New("mock")

	e := error(zhttp.BadRequest(errMock))

	for i := 0; i < b.N; i++ {
		switch ex := e.(type) {
		case zhttp.Error:
			ex.HTTPError()
		default:

		}
	}
}

func BenchmarkInterfaceCallWithOutAssertion(b *testing.B) {

	errMock := errors.New("mock")

	e := zhttp.Error(zhttp.BadRequest(errMock))

	for i := 0; i < b.N; i++ {
		e.HTTPError()
	}
}

func BenchmarkNoneTypeAssertion(b *testing.B) {

	errMock := errors.New("mock")

	e := zhttp.BadRequest(errMock)

	for i := 0; i < b.N; i++ {
		e.HTTPError()
	}
}

func BenchmarkHTTPHandlerFunc(b *testing.B) {

	type input struct {
	}

	type output struct {
		Message string
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&output{
			Message: "Hello world!",
		})
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "", nil)

	for i := 0; i < b.N; i++ {
		h(w, r)
	}
}

func BenchmarkGenericHandler(b *testing.B) {

	type input struct {
	}

	type output struct {
		Message string
	}

	h := zhttp.Handler(func(ctx context.Context, in *input) (*output, error) {
		return &output{
			Message: "Hello world!",
		}, nil
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "", nil)

	for i := 0; i < b.N; i++ {
		h(w, r)
	}
}

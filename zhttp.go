package zhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

type ContextKeyType int

const (
	RequestContextKey ContextKeyType = iota
)

// HandlerFunc is type signtature for handler function
type HandlerFunc[T1, T2 any] func(context.Context, *T1) (*T2, error)

// Handler convert zhttp.HandlerFunc to http.HandlerFunc
func Handler[T1, T2 any](f func(context.Context, *T1) (*T2, error)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var in T1

		if input, ok := any(&in).(Boundable); ok {
			err := input.Bind(r)
			if err != nil {
				handlerErr(w, err)
				return
			}
		} else {
			err := bind(r, &in)
			if err != nil {
				handlerErr(w, err)
				return
			}
		}

		ctx := context.WithValue(r.Context(), RequestContextKey, r)

		out, err := f(ctx, &in)

		if err != nil {
			handlerErr(w, err)
			return
		}

		if m, ok := any(out).(Marshalable); ok {
			data, err := m.Marshal(ctx)
			if err != nil {
				handlerErr(w, err)
			}
			fmt.Fprintf(w, "%s", data)
			return
		}

		if err = json.NewEncoder(w).Encode(out); err != nil {
			handlerErr(w, err)
		}
	})
}

func handlerErr(w http.ResponseWriter, err error) {
	if e, ok := err.(Error); ok {
		statusCode, body := e.HTTPError()
		w.WriteHeader(statusCode)
		fmt.Fprintf(w, "%s", body)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "%s", err.Error())
}

func bind(r *http.Request, in any) error {

	v := reflect.ValueOf(in).Elem()
	t := reflect.TypeOf(in).Elem()

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)

		var param string

		if urlTagValue, ok := ft.Tag.Lookup("url"); !ok {
			param = r.URL.Query().Get(ft.Name)
		} else {
			param = r.URL.Query().Get(urlTagValue)
		}

		if len(param) == 0 {
			continue
		}

		switch fv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			paramValue, _ := strconv.ParseInt(param, 10, 64)
			fv.SetInt(paramValue)
		case reflect.Bool:
			paramValue, _ := strconv.ParseBool(param)
			fv.SetBool(paramValue)
		case reflect.Float32, reflect.Float64:
			paramValue, _ := strconv.ParseFloat(param, 64)
			fv.SetFloat(paramValue)
		case reflect.String:
			fv.SetString(param)
		}
	}

	if r.Method == http.MethodGet {
		return nil
	}

	return json.NewDecoder(r.Body).Decode(in)
}

// RequestFromContext get the *http.Request from wrapped context
// this func use for some general propose, eg: upload file, handle form...
func RequestFromContext(ctx context.Context) *http.Request {
	r, ok := ctx.Value(RequestContextKey).(*http.Request)
	if ok {
		return r
	}
	return nil
}

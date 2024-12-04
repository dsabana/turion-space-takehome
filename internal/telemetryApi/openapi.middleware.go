package telemetryApi

import (
	"bytes"
	"context"
	"errors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ErrEmptySchema is returned when an []byte is passed to
// NewOpenAPIValidationMiddleware instead of a valid schema.
var ErrEmptySchema = errors.New("an empty OpenAPI schema has no effect")

// ErrFailedChain is returned when the middleware fails to rewrite the
// the response capture from the wrapped handler back to the wrapping
// handler's http.ResponseWriter
var ErrFailedChain = errors.New("failed to chain the http.ResponseWriters")

func newOpenAPIValidationMiddleware(schema []byte) (mux.MiddlewareFunc, error) {
	if len(schema) == 0 {
		return nil, ErrEmptySchema
	}

	doc, err := openapi3.NewLoader().LoadFromData(schema)
	if err != nil {
		return nil, err
	}

	for _, s := range doc.Servers {
		server := openapi3.Server{
			URL: s.URL,
		}

		serverURL, err := url.Parse(server.URL)
		if err != nil {
			return nil, err
		}

		scheme := "http"
		if serverURL.Scheme == "http" {
			scheme = "https"
		}

		serverURL.Scheme = scheme
		server.URL = serverURL.String()
		doc.Servers = append(doc.Servers, &server)
	}

	apiRouter, err := gorillamux.NewRouter(doc)
	if err != nil {
		return nil, err
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route, pathParams, err := apiRouter.FindRoute(r)
			if err != nil {
				writeAPIErrorResponse(w, http.StatusBadRequest,
					APIError{
						Status:  http.StatusBadRequest,
						Message: err.Error(),
					},
				)
				return
			}

			reqInp := &openapi3filter.RequestValidationInput{
				Request:    r,
				PathParams: pathParams,
				Route:      route,
				Options: &openapi3filter.Options{
					MultiError: true,
					AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
						return nil
					},
				},
			}

			if err := openapi3filter.ValidateRequest(r.Context(), reqInp); err != nil {
				writeAPIErrorResponse(w, http.StatusBadRequest,
					APIError{
						Status:  http.StatusBadRequest,
						Message: err.Error(),
					},
				)

				return
			}

			i := newInterceptingResponseWriter(w.Header().Clone())
			next.ServeHTTP(i, r)

			respInp := &openapi3filter.ResponseValidationInput{
				RequestValidationInput: reqInp,
				Status:                 i.statusCode,
				Header:                 i.header,
				Body:                   ioutil.NopCloser(i.Reader()),
			}

			if err := openapi3filter.ValidateResponse(r.Context(), respInp); err != nil {
				writeAPIErrorResponse(w, http.StatusInternalServerError,
					APIError{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					},
				)

				return
			}

			if err := i.Chain(w); err != nil {
				writeAPIErrorResponse(w, http.StatusInternalServerError,
					APIError{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					},
				)
			}
		})
	}, nil
}

type interceptingResponseWriter struct {
	header     http.Header
	statusCode int
	body       *bytes.Buffer
	written    bool
}

func newInterceptingResponseWriter(header http.Header) *interceptingResponseWriter {
	return &interceptingResponseWriter{
		header:     header,
		statusCode: http.StatusOK,
		body:       &bytes.Buffer{},
		written:    false,
	}
}

func (i *interceptingResponseWriter) Chain(w http.ResponseWriter) error {
	w.WriteHeader(i.statusCode)

	_, err := w.Write(i.body.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (i *interceptingResponseWriter) Header() http.Header {
	return i.header
}

func (i *interceptingResponseWriter) Reader() io.Reader {
	b := i.body.Bytes()
	i.body = bytes.NewBuffer(b)

	return bytes.NewReader(b)
}

func (i *interceptingResponseWriter) Write(p []byte) (int, error) {
	i.written = true

	return i.body.Write(p)
}

func (i *interceptingResponseWriter) WriteHeader(statusCode int) {
	i.statusCode = statusCode
}

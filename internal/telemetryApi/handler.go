package telemetryApi

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func SetupRouter(s Service) (*mux.Router, error) {
	router := mux.NewRouter()

	mw, err := newOpenAPIValidationMiddleware(Schema)
	if err != nil {
		return nil, err
	}

	registerDocumentationHandler(router)
	registerHandlers(router, mw, s)

	return router, nil
}

func registerDocumentationHandler(r *mux.Router) {
	r.Path("/docs/openapi.json").HandlerFunc(getDocumentation)

	r.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/openapi.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.PathPrefix("/doc").Handler(http.RedirectHandler("/docs/", http.StatusPermanentRedirect))
}

func registerHandlers(r *mux.Router, validationMiddleware mux.MiddlewareFunc, s Service) {
	r.StrictSlash(true)
	r.Use(commonMiddleware)

	v1GetTelemetry := r.PathPrefix("/api/v1/telemetry").Subrouter()
	v1GetTelemetry.Use(validationMiddleware)
	v1GetTelemetry.HandleFunc("", GetTelemetry(s)).Methods(http.MethodGet)
	v1GetTelemetry.HandleFunc("/current", GetCurrentTelemetry(s)).Methods(http.MethodGet)
	v1GetTelemetry.HandleFunc("/anomalies", GetAnomalyTelemetry(s)).Methods(http.MethodGet)
}

// RegisterCorsHandler enables cors support for this API.
func RegisterCorsHandler(r *mux.Router) http.Handler {
	if !APIConfig.CorsEnabled {
		return r
	}

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		AllowedOrigins: []string{APIConfig.CorsOrigins},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowCredentials: true,
	})

	return c.Handler(r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

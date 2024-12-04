//go:generate go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.16.2
//go:generate oapi-codegen --old-config-style -generate types -o ../../pkg/openapi/types_gen.go -package openapi openapi.json
//go:generate oapi-codegen --old-config-style -generate client -o ../../pkg/openapi/client_gen.go -package openapi openapi.json

package telemetryApi

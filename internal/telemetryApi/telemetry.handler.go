package telemetryApi

import (
	"encoding/json"
	"fmt"
	"github.com/dsabana/turion-space-takehome/pkg/openapi"
	"log"
	"net/http"
)

func GetTelemetry(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := r.URL.Query().Get("start_time")
		endTime := r.URL.Query().Get("end_time")

		// Convert empty string to nil
		var startTimePtr, endTimePtr *string
		if startTime != "" {
			startTimePtr = &startTime
		}
		if endTime != "" {
			endTimePtr = &endTime
		}

		data, err := s.GetTelemetryData(r.Context(), startTimePtr, endTimePtr)
		if err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, ErrRetrievingObject)
		}

		if err = json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("unable to encode GET response: %s", err)
			generateErrorResponse(w, http.StatusInternalServerError, ErrEncodingResponse)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func GetCurrentTelemetry(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.GetTelemetryCurrentData(r.Context())
		if err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, ErrRetrievingObject)
		}

		if err = json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("unable to encode GET response: %s", err)
			generateErrorResponse(w, http.StatusInternalServerError, ErrEncodingResponse)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func GetAnomalyTelemetry(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := r.URL.Query().Get("start_time")
		endTime := r.URL.Query().Get("end_time")

		// Convert empty string to nil
		var startTimePtr, endTimePtr *string
		if startTime != "" {
			startTimePtr = &startTime
		}
		if endTime != "" {
			endTimePtr = &endTime
		}

		data, err := s.GetTelemetryAnomaliesData(r.Context(), startTimePtr, endTimePtr)
		if err != nil {
			generateErrorResponse(w, http.StatusInternalServerError, ErrRetrievingObject)
		}

		if err = json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("unable to encode GET response: %s", err)
			generateErrorResponse(w, http.StatusInternalServerError, ErrEncodingResponse)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

type APIError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

// Error returns a string representation of an APIError and implements the error interface.
func (e APIError) Error() string {
	return fmt.Sprintf("HTTP Status Code: %d, Message: %s", e.Status, e.Message)
}

func writeAPIErrorResponse(w http.ResponseWriter, statusCode int, err APIError) error {
	w.WriteHeader(statusCode)

	response := openapi.ErrorResponse{
		Message: &err.Message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}

	return nil
}

func generateErrorResponse(w http.ResponseWriter, statusCode int, errorDetails error) {
	writeAPIErrorResponse(
		w,
		statusCode,
		APIError{
			Message: errorDetails.Error(),
			Status:  statusCode,
		},
	)
}

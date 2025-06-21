package helper

import (
	"context"
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func failResponseWriter(ctx context.Context, w http.ResponseWriter, err error, errStatusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	w.WriteHeader(errStatusCode)
	resp.StatusCode = errStatusCode
	resp.Message = err.Error()
	resp.Data = nil

	responseBytes, _ := json.Marshal(resp)
	w.Write(responseBytes)
}

func successResponseWriter(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	w.WriteHeader(statusCode)
	resp.StatusCode = statusCode
	resp.Message = "success"
	resp.Data = data

	responseBytes, _ := json.Marshal(resp)
	w.Write(responseBytes)
}

func WriteResponseSuccess(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	w.WriteHeader(statusCode)
	resp.StatusCode = statusCode
	resp.Message = message
	resp.Data = data

	responseBytes, _ := json.Marshal(resp)
	w.Write(responseBytes)
}

func WriteResponse(ctx context.Context, w http.ResponseWriter, err error, data any) {
	switch err.(type) {
	case *ErrForbidden, ErrForbidden:
		failResponseWriter(ctx, w, err, http.StatusForbidden)
	case *ErrUnauthorized, ErrUnauthorized:
		failResponseWriter(ctx, w, err, http.StatusUnauthorized)
	case *ErrNotFound, ErrNotFound:
		failResponseWriter(ctx, w, err, http.StatusNotFound)
	case *ErrBadRequest, ErrBadRequest:
		failResponseWriter(ctx, w, err, http.StatusBadRequest)
	case *ErrInternalServer, ErrInternalServer:
		failResponseWriter(ctx, w, err, http.StatusInternalServerError)
	case nil:
		successResponseWriter(ctx, w, data, http.StatusOK)
	default:
		failResponseWriter(ctx, w, err, http.StatusInternalServerError)
	}
}

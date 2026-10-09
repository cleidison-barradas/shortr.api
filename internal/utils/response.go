package utils

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/domain"
)

type ErrorCode string

const (
	INVALID_USER              ErrorCode = "invalid_user"
	USER_NOT_FOUND            ErrorCode = "user_not_found"
	USER_EMAIL_ALREADY_EXISTS ErrorCode = "user_email_already_exists"
	UNEXPECTED_ERROR          ErrorCode = "unexpected_error"
	MISSING_AUTH_HEADER       ErrorCode = "missing_auth_header"
	INVALID_AUTH_TOKEN        ErrorCode = "invalid_auth_token"
	BODY_REQUEST_ERROR        ErrorCode = "body_request_error"
	INVALID_PASSWORD          ErrorCode = "invalid_password"
	SHORT_LINK_NOT_FOUND      ErrorCode = "short_link_not_found"
	SHORT_LINK_ALREADY_EXISTS ErrorCode = "short_link_already_exists"
	LINK_EXPIRED              ErrorCode = "link_expired"
)

type ApiResponse struct {
	Result     any  `json:"result,omitempty"`
	Success    bool `json:"success,omitempty"`
	StatusCode int  `json:"status_code,omitempty"`
}

type AppError struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

type apiErrorMapping struct {
	target  error
	status  int
	code    ErrorCode
	message string
}

var erroMappings = []apiErrorMapping{
	{domain.ErrInvalidUser, http.StatusBadRequest, INVALID_USER, "invalid user"},
	{domain.ErrUserNotFound, http.StatusNotFound, USER_NOT_FOUND, "user not found"},
	{domain.ErrUserEmailAlreadyExists, http.StatusBadRequest, USER_EMAIL_ALREADY_EXISTS, "user email already exists"},
	{app.ErrBodyRequest, http.StatusBadRequest, BODY_REQUEST_ERROR, "invalid body request"},
	{app.ErrInvalidPassword, http.StatusUnauthorized, INVALID_PASSWORD, "invalid password"},
	{domain.ErrShorlinkAlreadyExists, http.StatusBadRequest, SHORT_LINK_ALREADY_EXISTS, "link already exists"},
	{app.ErrLinkExpired, http.StatusGone, LINK_EXPIRED, "link expired"},
}

func Success(w http.ResponseWriter, r *http.Request, data ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.StatusCode)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := http.StatusInternalServerError
	code := UNEXPECTED_ERROR
	message := err.Error()

	for _, m := range erroMappings {
		if errors.Is(err, m.target) {
			statusCode, code, message = m.status, m.code, m.message
			break
		}
	}

	if statusCode >= 500 {
		slog.ErrorContext(
			r.Context(),
			"unhandled error",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(AppError{
		Code:    code,
		Message: message,
	})
}

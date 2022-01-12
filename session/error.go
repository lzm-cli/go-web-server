package session

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

type Error struct {
	Status      int    `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"description"`
	trace       string
}

func (sessionError Error) Error() string {
	str, err := json.Marshal(sessionError)
	if err != nil {
		log.Panicln(err)
	}
	return string(str)
}

func ParseError(err string) (Error, bool) {
	var sessionErr Error
	json.Unmarshal([]byte(err), &sessionErr)
	return sessionErr, sessionErr.Code > 0 && sessionErr.Description != ""
}

func BadRequestError(ctx context.Context) Error {
	description := "The request body can’t be pasred as valid data."
	return createError(ctx, http.StatusAccepted, http.StatusBadRequest, description, nil)
}

func NotFoundError(ctx context.Context, description string) Error {
	if description == "" {
		description = "The endpoint is not found."
	}
	return createError(ctx, http.StatusAccepted, http.StatusNotFound, description, nil)
}

func AuthorizationError(ctx context.Context) Error {
	description := "Unauthorized, maybe invalid token."
	return createError(ctx, http.StatusAccepted, 401, description, nil)
}

func ForbiddenError(ctx context.Context) Error {
	description := http.StatusText(http.StatusForbidden)
	return createError(ctx, http.StatusAccepted, http.StatusForbidden, description, nil)
}

func ValidationError(ctx context.Context, description string) Error {
	return createError(ctx, http.StatusAccepted, http.StatusBadRequest, description, nil)
}

func TooManyRequestsError(ctx context.Context) Error {
	description := http.StatusText(http.StatusTooManyRequests)
	return createError(ctx, http.StatusAccepted, http.StatusTooManyRequests, description, nil)
}

func ServerError(ctx context.Context, err error) Error {
	description := http.StatusText(http.StatusInternalServerError)
	return createError(ctx, http.StatusInternalServerError, http.StatusInternalServerError, description, err)
}

func BlazeServerError(ctx context.Context, err error) Error {
	description := "Blaze server error."
	return createError(ctx, http.StatusInternalServerError, 7000, description, err)
}

func BlazeTimeoutError(ctx context.Context, err error) Error {
	description := "The blaze operation timeout."
	return createError(ctx, http.StatusInternalServerError, 7001, description, err)
}

func TransactionError(ctx context.Context, err error) Error {
	description := http.StatusText(http.StatusInternalServerError)
	return createError(ctx, http.StatusInternalServerError, 10001, description, err)
}

func BadDataError(ctx context.Context) Error {
	description := "The request data has invalid field."
	return createError(ctx, http.StatusAccepted, 10002, description, nil)
}

func InsufficientAccountBalanceError(ctx context.Context) Error {
	description := "Insufficient balance."
	return createError(ctx, http.StatusAccepted, 20117, description, nil)
}

func createError(ctx context.Context, status, code int, description string, err error) Error {
	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	trace := fmt.Sprintf("[ERROR %d] %s\n%s:%d %s", code, description, file, line, funcName)
	if err != nil {
		if sessionError, ok := err.(Error); ok {
			trace = trace + "\n" + sessionError.trace
		} else {
			trace = trace + "\n" + err.Error()
		}
	}

	return Error{
		Status:      status,
		Code:        code,
		Description: description,
		trace:       trace,
	}
}

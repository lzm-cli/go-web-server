package session

import (
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

func BadRequestError() Error {
	description := "The request body can’t be pasred as valid data."
	return createError(http.StatusAccepted, http.StatusBadRequest, description, nil)
}

func NotFoundError() Error {
	description := "The endpoint is not found."
	return createError(http.StatusAccepted, http.StatusNotFound, description, nil)
}

func AuthorizationError() Error {
	description := "Unauthorized, maybe invalid token."
	return createError(http.StatusAccepted, 401, description, nil)
}

func ForbiddenError() Error {
	description := http.StatusText(http.StatusForbidden)
	return createError(http.StatusAccepted, http.StatusForbidden, description, nil)
}

func ValidationError(description string) Error {
	return createError(http.StatusAccepted, http.StatusBadRequest, description, nil)
}

func TooManyRequestsError() Error {
	description := http.StatusText(http.StatusTooManyRequests)
	return createError(http.StatusAccepted, http.StatusTooManyRequests, description, nil)
}

func ServerError(err error) Error {
	description := http.StatusText(http.StatusInternalServerError)
	return createError(http.StatusInternalServerError, http.StatusInternalServerError, description, nil)
}

func BlazeServerError(err error) Error {
	description := "Blaze server error."
	return createError(http.StatusInternalServerError, 7000, description, nil)
}

func BlazeTimeoutError(err error) Error {
	description := "The blaze operation timeout."
	return createError(http.StatusInternalServerError, 7001, description, nil)
}

func TransactionError(err error) Error {
	description := http.StatusText(http.StatusInternalServerError)
	return createError(http.StatusInternalServerError, 10001, description, nil)
}

func BadDataError() Error {
	description := "The request data has invalid field."
	return createError(http.StatusAccepted, 10002, description, nil)
}

func InsufficientAccountBalanceError() Error {
	description := "Insufficient balance."
	return createError(http.StatusAccepted, 20117, description, nil)
}

func createError(status, code int, description string, err error) Error {
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
	if description == "" {
		description = err.Error()
	}
	return Error{
		Status:      status,
		Code:        code,
		Description: description,
		trace:       trace,
	}
}

package web

import "net/http"

type WebError struct {
	error
	statusCode int
}

func (we *WebError) PrintError(w http.ResponseWriter) {
	http.Error(w, we.Error(), we.statusCode)
}

func NewWebError(err error, statusCode int) WebError {
	return WebError{
		error:      err,
		statusCode: statusCode,
	}
}

func NewNotFoundError(err error) WebError {
	return NewWebError(err, http.StatusNotFound)
}

func NewBadRequestError(err error) WebError {
	return NewWebError(err, http.StatusBadRequest)
}

func NewInternalServerError(err error) WebError {
	return NewWebError(err, http.StatusInternalServerError)
}

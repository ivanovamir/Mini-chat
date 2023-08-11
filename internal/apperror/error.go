package apperror

import (
	"encoding/json"
	"errors"
)

type AppError struct {
	Err              error  `json:"err,omitempty"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func (a *AppError) Error() string { return a.Message }
func (a *AppError) Unwrap() error { return a.Err }

func (a *AppError) Marshal() []byte {
	marshal, err := json.Marshal(a)

	if err != nil {
		return nil
	}

	return marshal
}

func NewAppError(options ...Option) *AppError {
	ae := &AppError{}

	for _, opt := range options {
		opt(ae)
	}
	ae.Err = errors.New(ae.Message)
	return ae
}

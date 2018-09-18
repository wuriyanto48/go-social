package api

import (
	"context"
)

// Result generic result
type Result interface{}

// Service interface
type Service interface {
	GetAuthURI() (string, error)
	GetAccessToken(context.Context, string) error
	GetUser(context.Context) (Result, error)
}

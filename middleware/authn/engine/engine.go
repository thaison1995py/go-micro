package engine

import "context"

type ContextType int

const (
	ContextTypeGrpc = iota
	ContextTypeMicro
)

// Authenticator interface
type Authenticator interface {
	// Authenticate returns a claims info and nil error (if available).
	Authenticate(ctx context.Context, ctxType ContextType) (Claims, error)

	// CreateIdentity injects user claims into context.
	CreateIdentity(ctx context.Context, ctxType ContextType, claims Claims) (context.Context, error)
}

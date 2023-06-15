package middleware

import (
	"context"
	"errors"
)

// GetVerifiedClaims - is a function that handles the retrieval of verified claims.
//
//	@param ctx - context.Context
//	@param role - string
//	@return *DataI
//	@return error
func GetVerifiedClaims(ctx context.Context, role string) (*DataI, error) {
	// get the claims from the context
	claims, ok := ctx.Value(ContextKey).(*DataI)
	if !ok {
		return &DataI{}, errors.New("authentication failed: invalid token")
	}

	// check if the user has the required role
	if role != "" && !claims.HasRole(role) {
		return &DataI{}, errors.New("user does not have the required role")
	}

	return claims, nil
}

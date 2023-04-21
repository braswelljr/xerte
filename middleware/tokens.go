package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	Data       *DataI
	UserData   *User
	PrivateKey = os.Getenv("PRIVATE_KEY")
)

// ValidateToken - ValidateToken is a function that handles the verification of tokens.
//
//	@param ctx - context.Context
//	@param token - string
//	@param next - func(context.Context) (interface{}, error)
//	@return response
//	@return error
func ValidateToken(token string) (*SignedParams, error) {
	// get the private key from the secrets
	privateKey := PrivateKey

	// parse the token
	parsedToken, err := jwt.ParseWithClaims(token, &SignedParams{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		return nil, errors.New("authentication failed: invalid token")
	}

	// get the claims
	claims, ok := parsedToken.Claims.(*SignedParams)

	// check if the claims are valid
	if !ok || !parsedToken.Valid {
		return nil, errors.New("authentication failed: invalid token")
	}

	// Ensure token is valid not expired
	if !claims.VerifyExpiresAt(time.Now().Local(), true) {
		return nil, errors.New("authentication failed: token has expired")
	}

	// set the role in the context
	if err := Store.SetCtxValue("roles", []string{claims.User.Role}); err != nil {
		return nil, errors.New("authentication failed: unable to set role in context")
	}

	// get the role from the context
	roles := Store.GetCtxValue("roles").([]string)
	// check if the role is valid
	if len(roles) < 1 {
		return nil, errors.New("authentication failed: unable to get role from context")
	}

	// loop through the roles and check if the role is valid
	for _, role := range roles {
		if role == claims.User.Role {
			// set the user id in the context
			if err := Store.SetCtxValue("userID", claims.User.ID); err != nil {
				return nil, errors.New("authentication failed: unable to set user id in context")
			}
		}
	}

	// return the claims
	return claims, nil
}

// GetTokens - is a function that handles the retrieval of tokens.
//
//	@param user - *store.User
//	@return string
//	@return string
//	@return error
func GetToken(user *User) (string, error) {
	// get the private key from the secrets
	privateKey := PrivateKey

	// create a new token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &SignedParams{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    ContextKey.(string),
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}).SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}

	// return the token and refresh token
	return token, nil
}

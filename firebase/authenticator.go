package firebase

import (
	"context"
	"errors"

	"firebase.google.com/go/auth"

	"github.com/braswelljr/xerte/models"
)

// Login - is a function that logs in a user.
//
//	@param ctx - context.Context
//	@param email - string
//	@param password - string
//	@return *auth.UserRecord
//	@return error
func (a *Auth) Login(ctx context.Context, email, password string) (*auth.UserRecord, error) {
	// login the user with email and password
	user, err := a.Client.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("failed to login user")
	}

	// check for user password
	params := (&auth.UserToUpdate{}).Email(user.Email).Password(password)

	// update the user
	userRecord, err := a.Client.UpdateUser(ctx, user.UID, params)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return userRecord, nil
}

// Signup - is a function that signs up a user.
//
//	@param ctx - context.Context
//	@param email - string
//	@param password - string
//	@return *auth.UserRecord
//	@return error
func (a *Auth) Signup(ctx context.Context, params *models.SignupRequest) (*auth.UserRecord, error) {
	// create the user
	p := (&auth.UserToCreate{}).Email(params.Email).Password(params.Password).DisplayName(params.Fullname).PhoneNumber(params.Phone)

	// create the user
	userRecord, err := a.Client.CreateUser(ctx, p)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return userRecord, nil
}

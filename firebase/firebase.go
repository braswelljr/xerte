package firebase

import (
	"context"
	_ "embed"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	//go:embed keys/xerte-storage.json
	FirebaseConfig []byte
)

type Firestore struct {
	Client *firestore.Client
}

type Auth struct {
	Client *auth.Client
}

// New - is a function that initializes firebase app.
//
//	@param ctx - context.Context
//	@return *firebase.App
//	@return error
func New(ctx context.Context) (*firebase.App, error) {
	// get the config
	opts := option.WithCredentialsJSON(FirebaseConfig)

	// initialize the app
	app, err := firebase.NewApp(ctx, nil, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app %w", err)
	}

	return app, nil
}

// InitializeFirestore - is a function that initializes firestore client.
//
//	@param ctx - context.Context
//	@return *firestore.Client
//	@return error
func InitFirestore(ctx context.Context) (*Firestore, error) {
	// initialize the app
	app, err := New(ctx)
	if err != nil {
		return nil, err
	}

	// initialize the firestore client
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firestore client %w", err)
	}

	return &Firestore{Client: client}, nil
}

// InitAuthentication - is a function that initializes firebase authentication.
//
//	@param ctx - context.Context
//	@return *auth.Client
//	@return error
func InitAuth(ctx context.Context) (*Auth, error) {
	// initialize the app
	app, err := New(ctx)
	if err != nil {
		return nil, err
	}

	// initialize the auth client
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize auth client %w", err)
	}

	return &Auth{Client: client}, nil
}

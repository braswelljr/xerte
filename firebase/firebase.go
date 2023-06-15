package firebase

import (
	"context"
	_ "embed"
	"fmt"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	//go:embed keys/xerte-storage.json
	FirebaseConfig []byte
)

// InitializeFirebase - is a function that initializes firebase app.
//
//	@param ctx - context.Context
//	@return *firebase.App
//	@return error
func InitializeFirebase(ctx context.Context) (*firebase.App, error) {
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
func InitializeFirestore(ctx context.Context) (*firestore.Client, error) {
	// initialize the app
	app, err := InitializeFirebase(ctx)
	if err != nil {
		return nil, err
	}

	// initialize the firestore client
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firestore client %w", err)
	}

	return client, nil
}

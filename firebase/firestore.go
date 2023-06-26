package firebase

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
)

// Get - is a function that gets a single document from the database.
//
//	@param ctx - context.Context
//	@param collection - string
//	@param field - string
//	@param operation - string
//	@param value - interface{}
//	@return T
//	@return error
func (fs *Firestore) Get(ctx context.Context, collection, field, operation string, value interface{}) (*firestore.DocumentSnapshot, error) {
	// get the document
	documentSnapshot, err := fs.Client.Collection(collection).Where(field, operation, value).Limit(1).Documents(ctx).Next()
	if err != nil {
		return nil, errors.New("failed to get document")
	}

	return documentSnapshot, nil
}

// Add - is a function that adds a single document to the database.
//
//	@param ctx - context.Context
//	@param collection - string
//	@param data - interface{}
//	@return error
func (fs *Firestore) Add(ctx context.Context, collection string, data interface{}) (*firestore.DocumentSnapshot, error) {
	// add the document
	documentRef, _, err := fs.Client.Collection(collection).Add(ctx, data)
	if err != nil {
		return nil, errors.New("failed to add document")
	}

	// check if the document was added
	documentSnapshot, err := documentRef.Get(ctx)
	if err != nil {
		return nil, errors.New("failed to get document")
	}

	return documentSnapshot, nil
}

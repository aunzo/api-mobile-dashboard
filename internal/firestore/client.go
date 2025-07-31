package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

type ClientWrapper struct {
	Client *firestore.Client
	Ctx    context.Context
}

func NewClientWrapper(client *firestore.Client, ctx context.Context) *ClientWrapper {
	return &ClientWrapper{
		Client: client,
		Ctx:    ctx,
	}
}

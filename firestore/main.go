// Package main provides ...
package main

import (
	"context"

	"cloud.google.com/go/datastore"
)

func main() {
	client, err := datastore.NewClient(context.Background(), "sanxbox")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	commit, err := client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		return nil
	})
	if err != nil {
		panic(err)
	}
}

type ClientTxWrapper struct {
	ctx    context.Context
	client *datastore.Client
	tx     *datastore.Transaction
}

func clientTxWrapper(ctx context.Context) *ClientTxWrapper {
	return &ClientTxWrapper{}
}

func (c *ClientTxWrapper) Put(key *datastore.Key, v any) error {
	var (
		k  datastore.Key
		pk datastore.PendingKey
	)
	return nil
}

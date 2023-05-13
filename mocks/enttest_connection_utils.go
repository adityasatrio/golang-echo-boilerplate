package mocks

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"myapp/ent"
	"myapp/ent/enttest"
	"testing"
)

func TestConnection(t *testing.T) (*ent.Client, context.Context) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	ctx := context.Background()

	return client, ctx
}

func TestConnectionTx(t *testing.T) (*ent.Client, *ent.Tx, context.Context) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	ctx := context.Background()

	txClient, err := client.Tx(ctx)
	require.NoError(t, err)
	require.NotNil(t, txClient.Client())

	return client, txClient, ctx
}

func TestConnectionClose(client *ent.Client) {
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			fmt.Printf("error on connection db %s\n", err)
		}
	}(client)

}

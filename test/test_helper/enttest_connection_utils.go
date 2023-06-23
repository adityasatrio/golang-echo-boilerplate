package test_helper

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"myapp/ent"
	"myapp/ent/enttest"
	"testing"
)

func TestDbConnection(t *testing.T) (newClient *ent.Client, newContext context.Context) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	ctx := context.Background()

	require.NotNil(t, ctx)
	require.NotNil(t, client) //this lazy caller, mandatory for calling txClient.Client() so singleton struct will have same address

	return client, ctx
}

func TestDbConnectionTx(t *testing.T) (newClient *ent.Client, transaction *ent.Tx, newContext context.Context) {

	client, ctx := TestDbConnection(t)
	txClient, err := client.Tx(ctx)

	require.NoError(t, err)
	require.NotNil(t, txClient.Client()) //this lazy caller, mandatory for calling txClient.Client() so singleton struct will have same address

	return client, txClient, ctx
}

func TestDbConnectionClose(client *ent.Client) {
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			fmt.Printf("error on connection db %s\n", err)
		}
	}(client)
}

func TestDbConnectionCloseTx(transaction *ent.Tx) {
	defer func(transaction *ent.Tx) {
		err := transaction.Client().Close()
		if err != nil {
			fmt.Printf("error on connection db %s\n", err)
		}
	}(transaction)
}

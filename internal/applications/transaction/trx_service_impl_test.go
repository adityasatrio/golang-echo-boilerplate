package transaction

import (
	"context"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"myapp/ent"
	"myapp/ent/enttest"
	"testing"
)

func TestTrxServiceImpl_WithSuccessfulTx(t *testing.T) {

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	// CreateTx a TrxServiceImpl instance with the mock client.
	trxService := NewTrxService(client)
	ctx := context.Background()

	// Test a successful transaction.
	err := trxService.WithTx(context.Background(), func(tx *ent.Tx) error {
		// Insert some data into the database.
		_, err := tx.SystemParameter.Create().
			SetKey("John").
			SetValue("john@doe.com").
			SetCreatedBy("john doe").
			SetUpdatedBy("john doe").
			Save(ctx)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Errorf("WithTx returned an unexpected error: %v", err)
	}

}

func TestTrxServiceImpl_WithFailedTx(t *testing.T) {

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	// CreateTx a TrxServiceImpl instance with the mock client.
	trxService := NewTrxService(client)
	ctx := context.Background()

	// Test a failed transaction.
	expectedErr := errors.New("an error occurred during the transaction")
	err := trxService.WithTx(context.Background(), func(tx *ent.Tx) error {
		_, err := tx.SystemParameter.Create().
			SetKey("John cena").
			SetValue("john@cena.com").
			SetCreatedBy("john cena").
			SetUpdatedBy("john cena").
			Save(ctx)

		if err != nil {
			return err
		}
		return expectedErr
	})

	if err == nil {
		t.Error("WithTx did not return an expected error")
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("WithTx returned an unexpected error: %v", err)
	}
}

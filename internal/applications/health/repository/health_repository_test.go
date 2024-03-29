package repository

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/ent/enttest"
	"testing"
)

func TestHealthRepositoryImpl_Success(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	defer client.Close()

	// create a new repository
	repo := NewHealthRepository(client)
	ctx := context.Background()

	// test case 1
	t.Run("Health method success", func(t *testing.T) {
		msgParameter := "hello from service layer "
		expected := msgParameter + "hello from repository layer hello from query parameter"
		actual, err := repo.Health(ctx, msgParameter)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual["final_msg"])
		assert.Equal(t, "UP", actual["ctx_status"])
		assert.Equal(t, "echo", actual["ctx_name"])
		assert.Equal(t, "UP", actual["db_status"])
		assert.Equal(t, "mysql", actual["db_name"])
	})
}

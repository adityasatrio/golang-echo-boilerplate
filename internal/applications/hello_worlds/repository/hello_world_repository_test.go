package repository

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/ent/enttest"
	mock_repository "myapp/mocks/hello_worlds/repository"
	"testing"
)

func TestHelloWorldsRepositoryImpl_Hello(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	defer client.Close()

	// create a new repository
	repo := NewHelloWorldsRepository(client)
	ctx := context.Background()

	// test case 1
	t.Run("Hello method success", func(t *testing.T) {
		expected := "hello from repository-impl layer"
		actual, err := repo.Hello(ctx, "", "")
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	// test case 2
	t.Run("Hello method error from repository-impl layer", func(t *testing.T) {
		expectedErr := fmt.Errorf("error from repository-impl layer")
		mockRepo := &mock_repository.HelloWorldsRepository{}
		mockRepo.On("Hello", ctx, "", "repository").Return("", expectedErr)
		actual, err := mockRepo.Hello(context.Background(), "", "repository")
		assert.Equal(t, "", actual)
		assert.Equal(t, expectedErr, err)
	})
}

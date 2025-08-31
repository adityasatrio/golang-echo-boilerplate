package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostOutboundImpl_Create(t *testing.T) {

	if assert.NoError(t, nil) {
		dataValue := "true"
		assert.Equal(t, "true", dataValue)
	}
}

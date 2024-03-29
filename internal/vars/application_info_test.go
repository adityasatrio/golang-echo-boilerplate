package vars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplicationInfo(t *testing.T) {
	name := ApplicationName()
	assert.NotNil(t, name)
	assert.Equal(t, name, "myApp") //default value
}

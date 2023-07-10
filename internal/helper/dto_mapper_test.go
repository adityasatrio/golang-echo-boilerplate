package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Source struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Target struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMapper(t *testing.T) {
	source := Source{Name: "John", Age: 25}
	target := Target{}

	err := Mapper(&target, source)
	assert.NoError(t, err)

	assert.Equal(t, source.Name, target.Name)
	assert.Equal(t, source.Age, target.Age)
}

func TestMapAndAssign(t *testing.T) {
	source := Source{Name: "John", Age: 25}
	target := Target{}

	result, err := MapAndAssign(source, &target)
	assert.NoError(t, err)
	assert.NotNil(t, result)

}

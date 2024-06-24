package test18

import (
	"testing"

	"github.com/stretchr/testify/assert"

	test18 "test18.qstand.art/m/v2"
)

func TestExercise001(t *testing.T) {
	input := []string{"One", "Two", "Three"}
	assert.Equal(t, test18.Exercise001(input), "One, Two, Three")
}

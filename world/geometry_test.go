package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeoMid(t *testing.T) {
	assert.Equal(t, P{0, 0}, BB{P{-1, -1}, P{1, 1}}.Mid())
	assert.Equal(t, P{0.5, 0.5}, BB{P{0, 0}, P{1, 1}}.Mid())
}

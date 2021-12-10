package world

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToP(t *testing.T) {
	res := [][]P{
		{}, // Outside
		{P{0, -1}, P{-1, 0}},
		{P{0, -1}, P{1, 0}},
		{P{1, 0}, P{-1, 0}},
		{P{1, 0}, P{0, 1}},
		{P{0, -1}, P{1, 0}, P{0, 1}, P{-1, 0}}, // Saddle
		{P{0, -1}, P{0, 1}},
		{P{0, 1}, P{-1, 0}},
		{P{0, 1}, P{-1, 0}},
		{P{0, -1}, P{0, 1}},
		{P{1, 0}, P{0, 1}, P{0, -1}, P{-1, 0}}, // Saddle
		{P{1, 0}, P{0, 1}},
		{P{1, 0}, P{-1, 0}},
		{P{0, -1}, P{1, 0}},
		{P{0, -1}, P{-1, 0}},
		{}, // Inside
	}

	b := BB{P{-1, -1}, P{1, 1}}

	t.Run("Testing outside", func(t *testing.T) {
		assert.Equal(t, []P{}, Triangle{-100}.ToP(b))
	})

	for k, tr := range Triangulations {
		t.Run(fmt.Sprintf("Testing case %d", k), func(t *testing.T) {
			r := tr.ToP(b)

			assert.Equal(t, res[k], r)
		})
	}
}

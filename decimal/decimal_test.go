package decimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasics(t *testing.T) {
	d1 := New(1, 0)
	d2 := New(2, 0)
	d3 := d2.Add(d1)
	assert.Equal(t, d3.StringFixed(1), "3.0", "they should be equal")
	d1 = d1.Mul(d3).Mul(d3)
	assert.Equal(t, d1.StringFixed(1), "9.0", "they should be equal")
}

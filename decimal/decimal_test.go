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
	// json
	d9 := New(434, -1)
	bbb, _ := d9.MarshalJSON()
	d4 := New(1, 0)
	d4r := &d4
	err := d4r.UnmarshalJSON(bbb)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, d4.StringFixed(2), "4.34", "they should be equal")
}

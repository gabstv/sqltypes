package sqltypes

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	j := []byte(`{"a": null, "b": "with quotes", "c": 100}`)
	m := make(map[string]NullString)
	assert.NoError(t, json.Unmarshal(j, &m))
	assert.Equal(t, NullString(""), m["a"])
	assert.Equal(t, NullString("with quotes"), m["b"])
	assert.Equal(t, NullString("100"), m["c"])
}

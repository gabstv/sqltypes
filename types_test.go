package sqltypes

import (
	"encoding/json"
	"testing"
	"time"

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

func TestDecimal(t *testing.T) {
	assert.Equal(t, DecimalFromString("10.01"), DecimalFromString("10,01"))
	assert.Equal(t, DecimalFromString("123,010.01"), DecimalFromString("123.010,01"))
}

func TestNullDate(t *testing.T) {
	dd := NullDate("2018-03-09")
	assert.Equal(t, 2018, dd.Year())
	//
	ymdstr, err := dd.Value()
	assert.NoError(t, err)
	assert.Equal(t, ymdstr.(string), "2018-03-09")
	//
	str2 := struct {
		Date NullDate `json:"date"`
	}{}
	assert.NoError(t, json.Unmarshal([]byte("{\"date\":\"2019-07-22\"}"), &str2))
	assert.Equal(t, 22, str2.Date.Day())
	assert.Equal(t, 2019, str2.Date.Year())
	//
	// testing scan
	dd2 := &dd
	assert.NoError(t, dd2.Scan(time.Date(2010, time.Month(1), 10, 0, 0, 0, 0, time.Local)))
	assert.Equal(t, dd2.Year(), 2010)
	assert.Equal(t, dd2.Month(), time.Month(1))
	assert.Equal(t, dd2.Day(), 10)
	//
	assert.NoError(t, dd2.Scan("2013-09-25"))
	assert.Equal(t, dd2.Year(), 2013)
	assert.Equal(t, dd2.Month(), time.Month(9))
	assert.Equal(t, dd2.Day(), 25)
}

func TestNullDateMarshal(t *testing.T) {
	var dd NullDate
	dd = NullDate("1987-03-09")
	jj := make(map[string]interface{})
	jj["date"] = dd
	jj["a"] = "b"
	jj["c"] = 1000
	bb, err := json.Marshal(jj)
	assert.NoError(t, err)
	assert.NotNil(t, bb)
	t.Log(string(bb))
	assert.Equal(t, `{"a":"b","c":1000,"date":"1987-03-09"}`, string(bb))
}

type ntmap struct {
	Ts map[string]ntt
}

type ntt struct {
	T NullTime
	V string
	N time.Time
}

func TestNullTimeMap(t *testing.T) {
	z := ntmap{
		Ts: map[string]ntt{
			"a": ntt{
				V: "a",
				T: NullTime(time.Now()),
				N: time.Now(),
			},
			"b": ntt{
				V: "b",
			},
		},
	}
	zbytes, err := json.Marshal(&z)
	assert.NoError(t, err)
	t.Log(string(zbytes))
	z2 := &ntmap{}
	assert.NoError(t, json.Unmarshal(zbytes, z2))
}

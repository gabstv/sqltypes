package sqltypes

import (
	"database/sql"
	"database/sql/driver"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type NullInt0 int

// Scan implements the Scanner interface.
func (n *NullInt0) Scan(value interface{}) error {
	if value == nil {
		*n = 0
		return nil
	}

	ni64 := sql.NullInt64{}
	err := ni64.Scan(value)
	if err != nil {
		return err
	}

	*n = NullInt0(int(ni64.Int64))
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt0) Value() (driver.Value, error) {
	if n == 0 {
		return nil, nil
	}
	return int64(n), nil
}

type NullString string

// Scan implements the Scanner interface.
func (n *NullString) Scan(value interface{}) error {
	if value == nil {
		*n = ""
		return nil
	}
	vv, ok := value.(string)
	if ok {
		*n = NullString(vv)
	} else {
		vv2, _ := value.([]byte)
		*n = NullString(string(vv2))
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullString) Value() (driver.Value, error) {
	if n == "" {
		return nil, nil
	}
	return string(n), nil
}

// implements json.Unmarshaler
func (n *NullString) UnmarshalJSON(v []byte) error {
	if v == nil {
		return nil
	}
	if len(v) == 0 {
		return nil
	}
	if string(v) == "null" {
		return nil
	}
	if len(v) == 1 {
		*n = NullString(string(v))
		return nil
	}
	vv := string(v)
	if vv[0] == '"' && vv[len(vv)-1] == '"' {
		unq, err := strconv.Unquote(vv)
		if err != nil {
			return err
		}
		*n = NullString(unq)
		return nil
	}
	*n = NullString(vv)
	return nil
}

func (n NullString) String() string {
	return string(n)
}

type NullTime time.Time

func (t NullTime) T() time.Time {
	return time.Time(t)
}

func (t *NullTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if v, ok := value.(time.Time); ok {
		*t = NullTime(v)
		return nil
	}
	var e9 error
	var t9 time.Time
	switch v := value.(type) {
	case []byte:
		t9, e9 = time.Parse("2006-01-02 15:04:05", string(v))
	case string:
		t9, e9 = time.Parse("2006-01-02 15:04:05", v)
	}
	if e9 == nil {
		*t = NullTime(t9)
	}
	return e9
}

func (t NullTime) Value() (driver.Value, error) {
	v := t.T()
	if v.IsZero() {
		return nil, nil
	}
	return v, nil
}

// implements json.Unmarshaler
func (n *NullTime) UnmarshalJSON(v []byte) error {
	if v == nil {
		return nil
	}
	if len(v) == 0 {
		return nil
	}
	if string(v) == "null" {
		return nil
	}
	t2 := &time.Time{}
	err := t2.UnmarshalJSON(v)
	if err != nil {
		return err
	}
	*n = NullTime(*t2)
	return nil
}

// implements json.Marshaler
func (n *NullTime) MarshalJSON() ([]byte, error) {
	t := n.T()
	return t.MarshalJSON()
}

type NullDecimal decimal.Decimal

func (d NullDecimal) D() decimal.Decimal {
	return decimal.Decimal(d)
}

func (d *NullDecimal) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if v, ok := value.(decimal.Decimal); ok {
		*d = NullDecimal(v)
		return nil
	}
	ddd := decimal.New(0, 0)
	dd := &ddd
	err := dd.Scan(value)
	if err != nil {
		return err
	}
	*d = NullDecimal(*dd)
	return nil
}

func (d NullDecimal) Value() (driver.Value, error) {
	v := d.D()
	return v.Value()
}

// implements json.Unmarshaler
func (n *NullDecimal) UnmarshalJSON(v []byte) error {
	if v == nil {
		return nil
	}
	if len(v) == 0 {
		return nil
	}
	if string(v) == "null" {
		return nil
	}
	ddd := decimal.New(0, 0)
	t2 := &ddd
	err := t2.UnmarshalJSON(v)
	if err != nil {
		return err
	}
	*n = NullDecimal(*t2)
	return nil
}

// implements json.Marshaler
func (n *NullDecimal) MarshalJSON() ([]byte, error) {
	t := n.D()
	return t.MarshalJSON()
}

func DecimalFromString(s string) decimal.Decimal {
	dotI := strings.LastIndex(s, ".")
	comI := strings.LastIndex(s, ",")
	if dotI >= 0 && dotI > comI {
		s = strings.Replace(s, ",", "", -1)
	} else if comI >= 0 && comI > dotI {
		s = strings.Replace(s, ".", "", -1)
		s = strings.Replace(s, ",", ".", -1)
	}
	d, _ := decimal.NewFromString(s)
	return d
}

package sqltypes

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// NullBool is a bool that can be NULL (from DB)
type NullBool bool

// Scan implements the Scanner interface.
func (n *NullBool) Scan(value interface{}) error {
	if value == nil {
		*n = false
		return nil
	}

	ni64 := sql.NullInt64{}
	err := ni64.Scan(value)
	if err != nil {
		return err
	}

	*n = NullBool(ni64.Int64 != 0)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullBool) Value() (driver.Value, error) {
	if n == false {
		return int64(0), nil
	}
	return int64(1), nil
}

// NullInt0 is a normal int (0 = nil)
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

// NullIntM1 is a normal int (-1 = nil)
type NullIntM1 int

// Scan implements the Scanner interface.
func (n *NullIntM1) Scan(value interface{}) error {
	if value == nil {
		*n = -1
		return nil
	}

	ni64 := sql.NullInt64{}
	err := ni64.Scan(value)
	if err != nil {
		return err
	}

	*n = NullIntM1(int(ni64.Int64))
	return nil
}

// Value implements the driver Valuer interface.
func (n NullIntM1) Value() (driver.Value, error) {
	if n == -1 {
		return nil, nil
	}
	return int64(n), nil
}

// IsNull = (v == -1)
func (n NullIntM1) IsNull() bool {
	return n == -1
}

type NullUint64 uint64

// Scan implements the Scanner interface.
func (n *NullUint64) Scan(value interface{}) error {
	if value == nil {
		*n = 0
		return nil
	}

	vv := uint64(0)
	err := convertAssign(&vv, value)
	if err != nil {
		return err
	}
	*n = NullUint64(vv)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUint64) Value() (driver.Value, error) {
	if n == 0 {
		return nil, nil
	}
	return uint64(n), nil
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

// UnmarshalJSON implements json.Unmarshaler
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

// MarshalJSON implements json.Marshaler
func (n NullTime) MarshalJSON() ([]byte, error) {
	t := n.T()
	return t.MarshalJSON()
}

//
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

// UnmarshalJSON implements json.Unmarshaler
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

// MarshalJSON implements json.Marshaler
func (n NullDecimal) MarshalJSON() ([]byte, error) {
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

// NullFloat64 is a float64 with the 0 value being nil (on sending to sql)
type NullFloat64 float64

// Scan implements the Scanner interface.
func (n *NullFloat64) Scan(value interface{}) error {
	if value == nil {
		*n = 0
		return nil
	}

	nf64 := sql.NullFloat64{}
	err := nf64.Scan(value)
	if err != nil {
		return err
	}

	*n = NullFloat64(nf64.Float64)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullFloat64) Value() (driver.Value, error) {
	if n == 0 {
		return nil, nil
	}
	return float64(n), nil
}

//
//

type NullDate string

func (d NullDate) YMD() (year, month, day int) {
	ds := strings.Split(string(d), "-")
	if len(ds) != 3 {
		return 0, 0, 0
	}
	year, _ = strconv.Atoi(ds[0])
	month, _ = strconv.Atoi(ds[1])
	day, _ = strconv.Atoi(ds[2])
	return
}

func (d NullDate) T() time.Time {
	yy, mm, dd := d.YMD()
	return time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
}

func (d NullDate) IsZero() bool {
	yy, mm, dd := d.YMD()
	return yy == 0 || mm == 0 || dd == 0
}

func (d *NullDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if v, ok := value.(time.Time); ok {
		// v = v.UTC()
		*d = NullDate(v.Format("2006-01-02"))
		return nil
	}
	switch v := value.(type) {
	case []byte:
		return d.strscan(string(v))
	case string:
		return d.strscan(v)
	}
	*d = NullDate("0000-00-00")
	return nil
}

func (d *NullDate) strscan(v string) error {
	ymd := strings.Split(string(v), "-")
	if len(ymd) != 3 {
		return fmt.Errorf("invalid date '%s'", string(v))
	}
	*d = NullDate(v)
	return nil
}

// Value database/sql/
func (d NullDate) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	yy, mm, dd := d.YMD()
	return fmt.Sprintf("%04d-%02d-%02d", yy, mm, dd), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (d *NullDate) UnmarshalJSON(v []byte) error {
	if v == nil {
		return nil
	}
	if len(v) == 0 {
		return nil
	}
	if string(v) == "null" {
		return nil
	}
	if len(v) <= 2 {
		return fmt.Errorf("invalid date '%s'", string(v))
	}
	str := string(v)
	// remove quotes "0000-00-00"
	return d.strscan(str[1 : len(str)-1])
}

// MarshalJSON implements json.Marshaler
func (d *NullDate) MarshalJSON() ([]byte, error) {
	var yy, mm, dd int
	if d != nil {
		yy, mm, dd = d.YMD()
	}
	return []byte("\"" + fmt.Sprintf("%04d-%02d-%02d", yy, mm, dd) + "\""), nil
}

// Year returns the year
func (d NullDate) Year() int {
	yy, _, _ := d.YMD()
	return yy
}

// Month returns the month
func (d NullDate) Month() time.Month {
	_, mm, _ := d.YMD()
	return time.Month(mm)
}

// Day returns the day
func (d NullDate) Day() int {
	_, _, dd := d.YMD()
	return dd
}

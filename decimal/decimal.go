package decimal

import (
	"database/sql/driver"
	"encoding/binary"

	"github.com/shopspring/decimal"
)

//TODO: get more funcs
// https://github.com/shopspring/decimal/blob/master/decimal.go

type Decimal []byte

func New(value int64, exp int32) Decimal {
	intp := make([]byte, 10)
	expp := make([]byte, 5)
	binary.PutVarint(intp, value)
	binary.PutVarint(expp, int64(exp))
	intp = append(intp, expp...)
	return Decimal(intp)
}

func NewFromString(value string) (Decimal, error) {
	ld, err := decimal.NewFromString(value)
	if err != nil {
		return Decimal{}, err
	}
	return put(ld), nil
}

func NewFromFloat(value float64) Decimal {
	return put(decimal.NewFromFloat(value))
}

func NewFromFloatWithExponent(value float64, exp int32) Decimal {
	return put(decimal.NewFromFloatWithExponent(value, exp))
}

// Abs returns the absolute value of the decimal.
func (d Decimal) Abs() Decimal {
	dd := get(d)
	dd = dd.Abs()
	return put(dd)
}

// Add returns d + d2.
func (d Decimal) Add(d2 Decimal) Decimal {
	dd := get(d)
	dd2 := get(d2)
	dd = dd.Add(dd2)
	return put(dd)
}

func (d Decimal) Sub(d2 Decimal) Decimal {
	dd := get(d)
	dd2 := get(d2)
	dd = dd.Sub(dd2)
	return put(dd)
}

func (d Decimal) Neg() Decimal {
	dd := get(d)
	dd = dd.Neg()
	return put(dd)
}

func (d Decimal) Mul(d2 Decimal) Decimal {
	dd := get(d)
	dd2 := get(d2)
	dd = dd.Mul(dd2)
	return put(dd)
}

func (d Decimal) Div(d2 Decimal) Decimal {
	dd := get(d)
	dd2 := get(d2)
	dd = dd.Div(dd2)
	return put(dd)
}

// QuoRem
// DivRound
// Mod
// Pow
// Cmp
// Equal
// Equals
// GreaterThan
// GreaterThanOrEqual
// LessThan
// LessThanOrEqual
// Sign
// Exponent
// Coefficient
// IntPart
// Rat
// Float64

// String
// String returns the string representation of the decimal
// with the fixed point.
//
// Example:
//
//     d := New(-12345, -3)
//     println(d.String())
//
// Output:
//
//     -12.345
//
func (d Decimal) String() string {
	return get(d).String()
}

// StringFixed returns a rounded fixed-point string with places digits after
// the decimal point.
//
// Example:
//
// 	   NewFromFloat(0).StringFixed(2) // output: "0.00"
// 	   NewFromFloat(0).StringFixed(0) // output: "0"
// 	   NewFromFloat(5.45).StringFixed(0) // output: "5"
// 	   NewFromFloat(5.45).StringFixed(1) // output: "5.5"
// 	   NewFromFloat(5.45).StringFixed(2) // output: "5.45"
// 	   NewFromFloat(5.45).StringFixed(3) // output: "5.450"
// 	   NewFromFloat(545).StringFixed(-1) // output: "550"
//
func (d Decimal) StringFixed(places int32) string {
	return get(d).StringFixed(places)
}

// Round
// Floor
// Ceil
// Truncate

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Decimal) UnmarshalJSON(decimalBytes []byte) error {
	d9 := decimal.New(0, 0)
	ddv := &d9
	err := ddv.UnmarshalJSON(decimalBytes)
	if err != nil {
		return err
	}
	*d = put(*ddv)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d Decimal) MarshalJSON() ([]byte, error) {
	dd := get(d)
	return dd.MarshalJSON()
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (d *Decimal) UnmarshalBinary(data []byte) error {
	d9 := decimal.New(0, 0)
	ddv := &d9
	err := ddv.UnmarshalBinary(data)
	if err != nil {
		return err
	}
	*d = put(*ddv)
	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (d Decimal) MarshalBinary() (data []byte, err error) {
	dd := get(d)
	return dd.MarshalBinary()
}

// Scan implements the sql.Scanner interface for database deserialization.
func (d *Decimal) Scan(value interface{}) error {
	d9 := decimal.New(0, 0)
	ddv := &d9
	err := ddv.Scan(value)
	if err != nil {
		return err
	}
	*d = put(*ddv)
	return nil
}

// Value implements the driver.Valuer interface for database serialization.
func (d Decimal) Value() (driver.Value, error) {
	return d.String(), nil
}

// UnmarshalText
// MarshalText
// GobEncode
// GobDecode
// StringScaled
// Min
// Max

func put(d decimal.Decimal) Decimal {
	intpart := d.IntPart()
	exp := d.Exponent()
	intpb := make([]byte, 10)
	exppb := make([]byte, 5)
	binary.PutVarint(intpb, intpart)
	binary.PutVarint(exppb, int64(exp))
	intpb = append(intpb, exppb...)
	return Decimal(intpb)
}

func get(d Decimal) decimal.Decimal {
	bb := []byte(d)
	intpart, _ := binary.Varint(bb[0:10])
	exp64, _ := binary.Varint(bb[10:])
	return decimal.New(intpart, int32(exp64))
}

package sqltypes

import (
	"github.com/shopspring/decimal"
)

// String returns NullString
func String(v string) NullString {
	return NullString(v)
}

// Bool returns NullBool
func Bool(v bool) NullBool {
	return NullBool(v)
}

// Int0 returns NullInt0
func Int0(v int) NullInt0 {
	return NullInt0(v)
}

// IntM1 returns NullIntM1
func IntM1(v int) NullIntM1 {
	return NullIntM1(v)
}

// Decimal returns a NullDecimal
func Decimal(v decimal.Decimal) NullDecimal {
	return NullDecimal(v)
}

package sqltypes

import (
	"database/sql"
	"database/sql/driver"
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

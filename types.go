package sqltypes

import "database/sql/driver"

type NullInt0 int

// Scan implements the Scanner interface.
func (n *NullInt0) Scan(value interface{}) error {
	if value == nil {
		*n = 0
		return nil
	}
	vv, _ := value.(int64)
	*n = NullInt0(int(vv))
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt0) Value() (driver.Value, error) {
	if n == 0 {
		return nil, nil
	}
	return int(n), nil
}

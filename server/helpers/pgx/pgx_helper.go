package pgx

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

// NewTextFromString creates a valid pgtype.Text from a string
func NewTextFromString(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: true}
}

// NewNumericFromFloat creates a valid pgtype.Numeric from a float64
func NewNumericFromFloat(f float64) pgtype.Numeric {
	var num pgtype.Numeric
	num.Valid = true
	numStr := strconv.FormatFloat(f, 'f', -1, 64)
	_ = num.Scan(numStr)
	return num
}

// IsNumericNegative checks if a pgtype.Numeric value is negative
func IsNumericNegative(n pgtype.Numeric) bool {
	if !n.Valid {
		return false
	}

	var val float64
	if n.Scan(&val) == nil && val < 0 {
		return true
	}
	return false
}

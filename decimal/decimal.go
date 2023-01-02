package decimal

import (
	"math"
)

var ZERO Decimal = NewDecimal(0, 0)
var ONE Decimal = NewDecimal(1, 0)
var TEN Decimal = NewDecimal(10, 0)
var NEGATIVE Decimal = NewDecimal(-1, 0)

type Decimal struct {
	value     int64
	precision int //nachkommastellen
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func New(value int64, precision int) Decimal {
	return NewDecimal(value, precision)
}

func NewDecimal(value int64, precision int) Decimal {
	return Decimal{value: value, precision: precision}
}

func (ca Decimal) GetValue() int64 {
	return ca.value
}

func (ca Decimal) GetPrecission() int {
	return ca.precision
}

func (ca Decimal) AtPrecision(precission int) Decimal {
	if precission == ca.precision {
		return ca
	}
	increaseBy := precission - ca.precision
	factor := math.Pow10(increaseBy)
	value := float64(ca.value) * factor
	return Decimal{precision: precission, value: int64(value)}
}

func (ca Decimal) Add(add Decimal) Decimal {
	maxPrecison := max(ca.precision, add.precision)
	value := ca.AtPrecision(maxPrecison).value + add.AtPrecision(maxPrecison).value
	return Decimal{precision: maxPrecison, value: value}
}

func (ca Decimal) Subtract(add Decimal) Decimal {
	maxPrecison := max(ca.precision, add.precision)
	value := ca.AtPrecision(maxPrecison).value - add.AtPrecision(maxPrecison).value
	return Decimal{precision: maxPrecison, value: value}
}

func (ca Decimal) Multiply(add Decimal) Decimal {
	return Decimal{
		precision: (ca.precision + add.precision),
		value:     (ca.value * add.value),
	}
}

func (ca Decimal) Divide(add Decimal) Decimal {
	maxPrecison := max(ca.precision, add.precision)
	value := ca.AtPrecision(maxPrecison).value / add.AtPrecision(maxPrecison).value
	return Decimal{precision: maxPrecison, value: value}
}

/*
 * returns
 * a positive number if ca is greater than other
 * 0 if values are qeual
 * a negative number ist ca is lower than other
 */
func (ca Decimal) CompareTo(other Decimal) int {
	maxPrecison := max(ca.precision, other.precision)
	return int(ca.AtPrecision(maxPrecison).value - other.AtPrecision(maxPrecison).value)
}

func (ca Decimal) Invert() Decimal {
	return Invert(ca)
}

func Invert(ca Decimal) Decimal {
	return NewDecimal(ca.value*-1, ca.precision)
}

func absInt64(v int64) int64 {
	if v < 0 {
		return v * -1
	} else {
		return v
	}
}

func length(v int) int {
	if v == 0 {
		return 0
	}
	abs := Abs(v)
	l := 0
	for int(math.Pow10(l)) < abs {
		l++
	}
	return l + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

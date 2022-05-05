package decimal

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

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

func padRightZeros(number int, len int) string {
	var buffer bytes.Buffer

	if number > 0 {
		buffer.WriteString(fmt.Sprint(number))
	}

	idx := 0

	for idx <= (len - length(number)) {
		idx++
		buffer.WriteString("0")
	}
	return buffer.String()
}

func padLeftZeros(number int, len int) string {
	var buffer bytes.Buffer

	for int(math.Pow10(len-1)) > number {
		len--
		buffer.WriteString("0")
	}
	if number > 0 {
		buffer.WriteString(fmt.Sprint(number))
	}
	return buffer.String()
}

func NewDecimal(value int64, precision int) Decimal {
	return Decimal{value: value, precision: precision}
}

func DecimalFromString(value string) (Decimal, error) {

	if re, err := regexp.Compile("([0-9]+)(,([0-9]+))?"); err != nil {
		return NewDecimal(0, 0), err
	} else {
		parts := re.FindStringSubmatch(value)
		if len(parts) == 0 {
			return NewDecimal(0, 0), errors.New("FEHLER Beim Parsen, nichts gefunden")
		} else {
			preComma := parts[1]
			postComma := parts[3]
			preNum, _ := strconv.ParseInt(preComma, 10, 0)
			postNum, _ := strconv.ParseInt(postComma, 10, 0)
			precission := len(postComma)
			value := (preNum * int64(math.Pow10(precission))) + postNum
			return NewDecimal(value, precission), nil
		}
	}

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

func (ca Decimal) Format() string {
	vorkomma := ca.value / int64(math.Pow10(ca.precision))
	nachkomma := absInt64(ca.value - (vorkomma * int64(math.Pow10(ca.precision))))
	return fmt.Sprint(vorkomma) + "," + padLeftZeros(int(nachkomma), ca.precision) + " Eur"
}

func (ca Decimal) FormatPrecission(precision int) string {
	vorkomma := ca.value / int64(math.Pow10(ca.precision))
	nachkomma := ca.value - (vorkomma * int64(math.Pow10(ca.precision)))
	nachkomma = absInt64(nachkomma)
	precisionFactor := math.Pow10(precision - ca.precision)
	nachkommaPrecised := int(float64(nachkomma) * precisionFactor)
	return fmt.Sprint(vorkomma) + "," + padLeftZeros(nachkommaPrecised, precision) + " Eur"
}

func absInt64(v int64) int64 {
	if v < 0 {
		return v * -1
	} else {
		return v
	}
}

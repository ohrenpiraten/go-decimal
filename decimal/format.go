package decimal

import (
	"bytes"
	"fmt"
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func prefix(v int64) string {
	if v < 0 {
		return "-"
	} else {
		return ""
	}
}

func abs(v int64) int64 {
	if v < 0 {
		return v * -1
	} else {
		return v
	}
}

func (ca Decimal) Format() string {
	vorkomma := ca.value / int64(math.Pow10(ca.precision))
	nachkomma := absInt64(ca.value - (vorkomma * int64(math.Pow10(ca.precision))))

	p := message.NewPrinter(language.German)
	s := p.Sprintf("%d", vorkomma)

	return fmt.Sprint(s) + "," + padLeftZeros(int(nachkomma), ca.precision)
}

func (ca Decimal) FormatPrecission(precision int) string {
	p := message.NewPrinter(language.German)

	value := ca.value
	absValue := abs(value)
	vorkomma := absValue / int64(math.Pow10(ca.precision))
	vorkommaFormatted := p.Sprintf("%d", vorkomma)
	nachkomma := absValue - (vorkomma * int64(math.Pow10(ca.precision)))
	nachkomma = absInt64(nachkomma)
	precisionFactor := math.Pow10(precision - ca.precision)
	nachkommaPrecised := int(float64(nachkomma) * precisionFactor)
	return prefix(value) + vorkommaFormatted + "," + padLeftZeros(nachkommaPrecised, precision)
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

package decimal

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

func DecimalFromString(value string) (Decimal, error) {

	if re, err := regexp.Compile(`(\-|\+)?(([0-9]\.)*([0-9]+))(,([0-9]+))?`); err != nil {
		return ZERO, err
	} else {
		parts := re.FindStringSubmatch(value)
		if len(parts) == 0 {
			return NewDecimal(0, 0), errors.New("FEHLER Beim Parsen, nichts gefunden")
		} else {

			reDeleteDots, err := regexp.Compile(`\.`)

			if err != nil {
				return ZERO, err
			}

			prefix := parts[1]
			preComma := reDeleteDots.ReplaceAllLiteralString(parts[2], "")
			postComma := parts[6]
			preNum, _ := strconv.ParseInt(preComma, 10, 0)
			postNum, _ := strconv.ParseInt(postComma, 10, 0)
			precission := len(postComma)
			prefixFactor := 1
			if prefix == "-" {
				prefixFactor = -1
			}
			value := (preNum * int64(math.Pow10(precission))) + postNum

			return NewDecimal(int64(prefixFactor)*value, precission), nil
		}
	}

}

func Parse(value string) (Decimal, error) {
	return DecimalFromString(value)
}

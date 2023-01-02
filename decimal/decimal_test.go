package decimal

import (
	"testing"
)

func TestParse(t *testing.T) {

	assert := func(strV string, value int64, precission int) {
		dec, err := DecimalFromString(strV)

		if nil != err {
			t.Errorf(err.Error())
		}
		if dec.value != value {
			t.Errorf("[parse: %s] expected value %d but was %d ", strV, value, dec.value)
		}
		if dec.precision != precission {
			t.Errorf("[parse: %s] expected value %d but was %d ", strV, precission, dec.precision)
		}
	}
	assert("-1.077,53", -107753, 2)
	assert("1", 1, 0)
	assert("1,0", 10, 1)
	assert("123,180", 123180, 3)
	assert("0,", 0, 0)
	assert("0,0", 0, 1)
	assert("1,1", 11, 1)
	assert("0,999", 999, 3)

}

func TestLength(t *testing.T) {

	assert := func(value int, len int) {
		if length(1) != 1 {
			t.Errorf("length(1) should not be %d", length(1))
		}
	}

	assert(0, 0)
	assert(1, 1)
	assert(9, 1)
	assert(10, 2)
	assert(-10, 2)
	assert(99, 2)
	assert(-99, 2)
}

func TestRighPadd(t *testing.T) {
	assertRightPad(t, 99, 2, "99")
	assertRightPad(t, 99, 0, "99")
	assertRightPad(t, 99, 1, "99")
	assertRightPad(t, 99, 2, "99")
	assertRightPad(t, 99, 3, "990")
	assertRightPad(t, 99, 4, "9900")
}

func TestLeftPadd(t *testing.T) {
	assertLeftPad(t, 0, 0, "") // länge null?
	assertLeftPad(t, 0, 1, "0")
	assertLeftPad(t, 0, 2, "00")

	assertLeftPad(t, 9, 3, "009")
	assertLeftPad(t, 9, 1, "9")
	assertLeftPad(t, 9, 0, "9")
	assertLeftPad(t, 99, 2, "99")
	assertLeftPad(t, 1000, 0, "1000")
	assertLeftPad(t, 1000, 3, "1000")
	assertLeftPad(t, 1000, 4, "1000")
	assertLeftPad(t, 1000, 5, "01000")

	assertLeftPad(t, 999, 0, "999")
	assertLeftPad(t, 999, 3, "999")
	assertLeftPad(t, 999, 4, "0999")
	assertLeftPad(t, 999, 5, "00999")

	assertLeftPad(t, 1001, 3, "1001")
	assertLeftPad(t, 1001, 4, "1001")
	assertLeftPad(t, 1001, 5, "01001")
}

func TestIncreasePrecision(t *testing.T) {
	amount := Decimal{value: 1100, precision: 2}.AtPrecision(3)
	assertCurrencyAmountValue(t, amount, 11000)
	assertCurrencyAmountPrecision(t, amount, 3)
}

func TestDecreasePrecision(t *testing.T) {
	amount := Decimal{value: 1234, precision: 2}.AtPrecision(1)
	assertCurrencyAmountValue(t, amount, 123)
	assertCurrencyAmountPrecision(t, amount, 1)
}

func TestAddSamePrecission(t *testing.T) {
	sum := Decimal{value: 1100, precision: 2}.Add(Decimal{value: 20022, precision: 2})
	assertCurrencyAmountValue(t, sum, 21122)
	assertCurrencyAmountPrecision(t, sum, 2)
}

func TestAddIncreasingPrecission(t *testing.T) {
	sum := Decimal{value: 1100, precision: 2}.Add(Decimal{value: 200229, precision: 3})
	assertCurrencyAmountValue(t, sum, 211229)
	assertCurrencyAmountPrecision(t, sum, 3)
}

func assertCurrencyAmountValue(t *testing.T, amount Decimal, value int) {
	if int(amount.value) != value {
		t.Errorf("wrong value. expreced '%v, but  got '%v'", value, amount.value)
	}
}

func assertCurrencyAmountPrecision(t *testing.T, amount Decimal, precision int) {
	if amount.precision != precision {
		t.Errorf("wrong value. expreced '%v, but  got '%v'", precision, amount.precision)
	}
}

func TestPrint(t *testing.T) {
	a_1234_2 := Decimal{value: 1234, precision: 2}
	if a_1234_2.Format() != "12,34" {
		t.Errorf("Falsches Format. want '12,34' got '%v'", a_1234_2.Format())
	}
}

func TestFormatScalePadded(t *testing.T) {
	zero := Decimal{value: 0, precision: 2}
	formatted := zero.Format()
	expected := "0,00"
	if formatted != expected {
		t.Errorf("Falsches Format. want '%v' got '%v'", expected, formatted)
	}
}

func TestPrintWithDots(t *testing.T) {
	number := Decimal{value: 12345678912, precision: 4}
	formatted := number.Format()
	expected := "1.234.567,8912"
	if formatted != expected {
		t.Errorf("Falsches Format. expect '%v' got '%v'", expected, formatted)
	}
}

func TestFormatPrecission(t *testing.T) {
	assert := func(decimal Decimal, precission int, expected string) {
		formatted := decimal.FormatPrecission(precission)
		if formatted != expected {
			t.Errorf("Falsches Format. expect to Format %d:%d as '%v' but got '%v'", decimal.value, decimal.precision, expected, formatted)
		}
	}
	assert(Decimal{value: -412345678912, precision: 4}, 2, "-41.234.567,89")
	assert(Decimal{value: 12345678912, precision: 4}, 2, "1.234.567,89")
	assert(Decimal{value: 10199, precision: 2}, 2, "101,99")
	assert(Decimal{value: -10199, precision: 2}, 2, "-101,99")
	assert(Decimal{value: -55, precision: 2}, 2, "-0,55")
}

func assertFormatPrecision(t *testing.T, decimal Decimal, precission int, expected string) {
	formatted := decimal.FormatPrecission(precission)
	if formatted != expected {
		t.Errorf("Falsches Format. expect to Format %d:%d as '%v' but got '%v'", decimal.value, decimal.precision, expected, formatted)
	}
}

func assertLeftPad(t *testing.T, value int, len int, expected string) {
	actual := padLeftZeros(value, len)
	if actual != expected {
		t.Errorf("expected %d:%d to be left-padded as '%v', but was '%v'", value, len, expected, actual)
	}
}

func assertRightPad(t *testing.T, value int, len int, expected string) {
	actual := padRightZeros(value, len)
	if actual != expected {
		t.Errorf("expected %d:%d to be right-padded as '%v', but was '%v'", value, len, expected, actual)
	}
}

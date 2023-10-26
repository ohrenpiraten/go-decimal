package decimal

func (dec Decimal) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dec.Format() + `"`), nil
}

func (dec *Decimal) UnmarshalJSON(data []byte) error {
	v := data[1 : len(data)-1]
	d, err := Parse(string(v))
	if err != nil {
		return err
	}
	*dec = d
	return err
}

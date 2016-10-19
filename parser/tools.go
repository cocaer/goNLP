package begojson

import (
	"strconv"
	"unicode/utf8"
)

func assertValueNotNull(v *begoValue) {
	if v == nil {
		panic("*begoValue cannot be nil")
	}
}

/*get the _type of begoValue*/
func getJSONType(v *begoValue) jsonType {
	assertValueNotNull(v)
	return v._type
}

func getNumber(v *begoValue) float64 {
	assertValueNotNull(v)
	return v.value
}

func setNumber(v *begoValue, n float64) {

	freeValue(v)
	v.value = n
	v._type = jsonNUMBER
}

func getBoolen(v *begoValue) bool {
	if v == nil || (v._type != jsonFALSE && v._type != jsonTRUE) {
		panic("*begoValue is wrong in getBoolen")
	}

	return v._type == jsonTRUE
}

func setBoolen(v *begoValue, b bool) {

	freeValue(v)
	v._type = jsonFALSE
	if b {
		v._type = jsonTRUE
	}
}

func setString(v *begoValue, s string) {
	freeValue(v)
	v._type = jsonSTRING
	v.str = s
}

func freeValue(v *begoValue) {
	assertValueNotNull(v)
	v.str = ""
}

func isDigit(ch byte) bool {

	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func isDigit1To9(ch byte) bool {

	if ch >= '1' && ch <= '9' {
		return true
	}
	return false
}

type stack []byte

func (c *context) pushByte(b byte) {
	c.s = append(c.s, b)
}

func (c *context) pushBytes(b []byte) {
	c.s = append(c.s, b...)
}

func (c *context) popBytes(length int) []byte {
	if length > len(c.s) {
		panic("pop number is bigger than size of stack")
	}
	ret := c.s[len(c.s)-length:]
	c.s = c.s[0 : len(c.s)-length]
	return ret
}

func getu4(s []byte) (rune, uint64) {
	if len(s) < 5 || s[0] != 'u' {
		return -1, 1
	}
	r, err := strconv.ParseUint(string(s[1:5]), 16, 64)
	if err != nil {
		return -1, 1
	}
	return rune(r), r
}

func push4u(r rune, c *context) {
	buff := make([]byte, 6)
	n := utf8.EncodeRune(buff, r)
	for i := 0; i < n; i++ {
		c.pushByte(buff[i])
	}
}

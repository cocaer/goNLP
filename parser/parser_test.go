package begojson

import "testing"

/*
func TestParserExceptValue(t *testing.T){
	v := begoValue{  }

	v._type=jsonFALSE
	status :=parser(&v,   " ")
	testBase(ParserExpectValue, status, t, "parser fail")
	testBase(jsonNULL, getJSONType(&v), t, "test parser null")
}*/

func testBase(except interface{}, actual interface{}, t *testing.T, msg string) {

	if except != actual {
		t.Error(msg)
	}
}

/*
func TestParserNull(t *testing.T) {
	var v begoValue
	var status begoParserStatus
	setBoolen(&v, false)
	v, status = initAndParser("null")
	testBase(ParserOk, status, t, "parser fail")
	testBase(jsonNULL, getJSONType(&v), t, "test parser null")
}

func TestParserInvalid(t *testing.T) {
	var v begoValue
	var status begoParserStatus
	setBoolen(&v, false)
	v, status = initAndParser("?")
	testBase(ParserInvalidValue, status, t, "parser fail")
	testBase(jsonNULL, getJSONType(&v), t, "test parser Invalid")

	v, status = initAndParser("nul")

	testBase(ParserInvalidValue, status, t, "parser fail")
	testBase(jsonNULL, getJSONType(&v), t, "test parser Invalid")
}

func TestParserRootNotSingular(t *testing.T) {

	v, status := initAndParser("nulllllll")
	testBase(ParserRootNotSingular, status, t, "parser fail")
	testBase(jsonNULL, getJSONType(&v), t, "test parser null")

}

func TestParserFalse(t *testing.T) {

	v, status := initAndParser("false")
	testBase(ParserOk, status, t, "parser fail")

	testBase(jsonFALSE, getJSONType(&v), t, "test parser false")
}

func TestParserTrue(t *testing.T) {
	v, status := initAndParser("true")
	testBase(ParserOk, status, t, "parser fail")

	testBase(jsonTRUE, getJSONType(&v), t, "test parser true")
	testBase(true, getBoolen(&v), t, "get true wrong")
	setBoolen(&v, false)
	testBase(jsonFALSE, v._type, t, "set fasle wrong")

}
*/
func TestParserNumber(t *testing.T) {

	testNumberEqual(t, "  223  ", 223)
	testNumberEqual(t, "0.123", 0.123)
	testNumberEqual(t, "1.123", 1.123)
	testNumberEqual(t, "-0.123", -0.123)
	testNumberEqual(t, "8e10", 8e10)
	testNumberEqual(t, "1.234E+10   ", 1.234E+10)
	testNumberEqual(t, "1e-10000", 0.0)
	testNumberEqual(t, "1E+10", 1E+10)
	testNumberEqual(t, "1E-10", 1E-10)
	testNumberEqual(t, "1E-10", 1E-10)
	testNumberInvaild(t, "inf", ParserInvalidValue)
	testNumberInvaild(t, "+0", ParserInvalidValue)

	testNumberInvaild(t, "1.", ParserInvalidValue)
	testNumberInvaild(t, ".123", ParserInvalidValue)
	testNumberInvaild(t, "NAN", ParserInvalidValue)
}

func TestParserStack(t *testing.T) {
	c := context{}

	c.pushByte('a')
	c.pushByte('b')
	c.pushByte('c')

	v := c.popBytes(2)

	if v[0] != 'b' && v[1] != 'c' {
		t.Error()
	}
	c.pushBytes([]byte{'e', 'f'})

	v = c.popBytes(2)

	if v[0] != 'e' && v[1] != 'f' {
		t.Error()
	}

}

func testNumberInvaild(t *testing.T, strNum string, s begoParserStatus) {
	v, status := initAndParser(strNum)
	testBase(s, status, t, "parser fail")
	v._type = jsonNULL

}

func testNumberEqual(t *testing.T, strNum string, num float64) {
	v, status := initAndParser(strNum)
	testBase(ParserOk, status, t, "parser fail")
	testBase(num, getNumber(&v), t, "test parser number")
}

func initAndParser(json string) (v begoValue, s begoParserStatus) {
	v._type = jsonFALSE
	s = parser(&v, json)
	return v, s
}

func TestParserString(t *testing.T) {
	/*testStringEqual(t, "", `""`)
	testStringEqual(t, "\x30", "\"\x30\"")
	testStringEqual(t, "abc", `"abc"`)
	testStringEqual(t, "\\", `"\\"`)
	testStringEqual(t, "劲", `"\u52B2"`)
	testStringEqual(t, "\xF0\x9D\x84\x9E", "\"\\uD834\\uDD1E\"")
	testStringEqual(t, "\xF0\x9D\x84\x9E", "\"\\ud834\\udd1e\"")

	testStringEqual(t, "json\n", `"json\n"`) // { "name" : "json\n"}

	testParserMissQuotationMark(t)
	testParserInvalidStringEscape(t)*/
}

func testParserMissQuotationMark(t *testing.T) {
	_, status := initAndParser(`"json`)
	testBase(ParserMissQuotationMark, status, t, "parser Miss Quotation")
}

func testParserInvalidStringEscape(t *testing.T) {
	_, status := initAndParser(`"\v"`)
	testBase(ParserInvalidStringEscape, status, t, "parser  invalid string escape ")
	_, status = initAndParser("\"\x10\"")
	testBase(ParserInvalidStringChar, status, t, "parser invalid char")

}

func testStringEqual(t *testing.T, s1 string, s2 string) {
	v, status := initAndParser(s2)
	testBase(ParserOk, status, t, "parser string fail")
	//t.Error(".....", v.str)
	//t.Error(".....", s1)
	testBase(s1, v.str, t, "parser string equal fail")
}

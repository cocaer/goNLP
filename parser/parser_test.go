package begojson

import (
	"testing"
)

/*
func TestParserExceptValue(t *testing.T){
	v := begoValue{  }

	v._type=jsonFALSE
	status :=parser(&v,   " ")
	testBase(ParserExpectValue, status, t, "parser fail")
	testBase(jsonNULL, getJSONType(&v), t, "test parser null")
}*/

func testBase(except interface{}, actual interface{}, t *testing.T, msg string) {

	//fmt.Println("!!!!!!!!!!",except,actual)
	if except != actual {
		t.Error(msg)
	}
}

func TestParserNull(t *testing.T) {

	v, status := initAndParser("null")

	testBase(ParserOk, status, t, "parser fail")
	testBase(jsonNULL, getJSONType(&v), t, "test parser null")
}

func TestParserInvalid(t *testing.T) {
	v, status := initAndParser("?")

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
}

func TestParserNumber(t *testing.T) {

	testNumberEqual(t, "  223", 223)
	testNumberEqual(t, "0.123", 0.123)
	testNumberEqual(t, "1.123", 1.123)
	testNumberEqual(t, "-0.123", -0.123)
	testNumberEqual(t, "8e10", 8e10)
	testNumberEqual(t, "1.234E+10", 1.234E+10)
	testNumberEqual(t, "1e-10000", 0.0)
	testNumberEqual(t, "1E+10", 1E+10)
	testNumberEqual(t, "1E-10", 1E-10)

	testNumberInvaild(t, "inf", ParserInvalidValue)
	testNumberInvaild(t, "+0", ParserInvalidValue)

	testNumberInvaild(t, "1.", ParserInvalidValue)
	testNumberInvaild(t, ".123", ParserInvalidValue)
	testNumberInvaild(t, "NAN", ParserInvalidValue)
}

func testNumberInvaild(t *testing.T, strNum string, status begoParserStatus) {
	v, status := initAndParser(strNum)
	testBase(status, status, t, "parser fail")
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

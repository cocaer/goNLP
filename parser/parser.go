package begojson

import (
	"fmt"
	"strconv"
)

//import "fmt"

/* type for BegoNull ...*/
type jsonType int8

/* type for ParserOk...*/
type begoParserStatus int8

/* valueType in json*/
const (
	jsonNULL jsonType = iota
	jsonFALSE
	jsonTRUE
	jsonNUMBER
	jsonSTRING
	jsonARRAY
	jsonOBJECT
)

/* result from parser*/
const (
	ParserOk begoParserStatus = iota
	ParserExpectValue
	ParserInvalidValue
	ParserRootNotSingular //3
	ParserNumberTooBig
	ParserInvalidStringEscape //5
	ParserMissQuotationMark
	ParserInvalidStringChar
)

/* store info from parser*/
type begoValue struct {
	_type jsonType
	value float64 /*it's sad  that golang has no union type like C! */
	str   string  /*why dont use interface{} */
	len   int     /*because interface{} spends more space and time*/
}

/*our context to store json file string and other things*/
type context struct {
	json  string
	index int //index for json string
	s     stack
}

/*skip the white space*/
func (c *context) parserWhiteSpace() {

	i := c.index
	str := c.json
	length := len(str)

	if i >= length {
		return
	}

	for str[i] == ' ' || str[i] == '\t' || str[i] == '\n' || str[i] == '\r' {
		i++
		if i >= length {
			c.index = i
			return
		}
	}
	c.index = i
}

/*parser true false  null*/
func (c *context) parserCommon(aimStr string, v *begoValue) begoParserStatus {
	i := c.index
	str := c.json
	length := len(aimStr)
	leng := len(str)
	//if leng <= i+length-1 || strings.EqualFold(str[i:i+length-1], aimStr) {
	if leng <= i+length-1 || str[i:i+length-1] == aimStr {
		return ParserInvalidValue
	}
	c.index += length

	if aimStr == "null" {

		v._type = jsonNULL
	} else if aimStr == "false" {

		v._type = jsonFALSE
	} else {

		v._type = jsonTRUE
	}
	return ParserOk
}

/*parser number*/
func (c *context) parserNumber(v *begoValue) begoParserStatus {

	i := c.index
	json := c.json
	leng := len(c.json)
	if i < leng && json[i] == '-' {
		i++
	}

	if i < leng && json[i] == '0' {
		i++
	} else {
		if i >= leng || !isDigit1To9(json[i]) {

			return ParserInvalidValue
		}

		for i = i + 1; i < leng && isDigit(json[i]); i++ {
		}
	}

	if i < leng && json[i] == '.' {
		i++
		if i >= leng || !isDigit(json[i]) {
			return ParserInvalidValue
		}
		for i = i + 1; i < leng && isDigit(json[i]); i++ {
		}

	}

	if i < leng && (json[i] == 'e' || json[i] == 'E') {
		i++
		if i < leng && json[i] == '+' || json[i] == '-' {
			i++
		}

		if i >= leng || !isDigit(json[i]) {
			return ParserInvalidValue
		}
		for i = i + 1; i < leng && isDigit(json[i]); i++ {
		}

	}

	value, erron := strconv.ParseFloat(c.json[c.index:i], 64)

	//number is  to big
	if erron != nil {
		return ParserNumberTooBig
	}

	v.value = value
	v._type = jsonNUMBER
	c.index = i

	return ParserOk

}

func parser(v *begoValue, json string) begoParserStatus {
	c := context{json: json, index: 0}
	leng := len(json)
	v._type = jsonNULL //initize the type

	c.parserWhiteSpace()
	ret := c.parserValue(v)
	c.parserWhiteSpace()

	if ret == ParserOk {
		c.parserWhiteSpace()
		if c.index < leng {
			fmt.Println("c.index, leng", c.index, leng)
			ret = ParserRootNotSingular
		}
	}

	return ret
}

/*return the status of parser*/
func (c *context) parserValue(v *begoValue) begoParserStatus {
	switch ch := c.json[c.index]; ch {
	case 'n':
		return c.parserCommon("null", v)
	case 't':
		return c.parserCommon("true", v)
	case 'f':
		return c.parserCommon("false", v)
	case '"':
		return c.parserString(v)
	default:
		return c.parserNumber(v)
	}
}

/*parser string*/
func (c *context) parserString(v *begoValue) begoParserStatus {

	head := len(c.s)
	i := c.index + 1

	for ; i < len(c.json); i++ {

		ch := c.json[i]
		//	fmt.Printf("#####%d %c\n", i, ch)
		switch ch {

		case '"':
			size := len(c.s) - head
			str := string(c.popBytes(size))
			setString(v, str)
			c.index = i + 1
			return ParserOk

		case '\\':
			i++
			ch = c.json[i]
			switch ch {
			case '"':
				c.pushByte('"')
				break
			case '\\':
				c.pushByte('\\')
				break
			case '/':
				c.pushByte('/')
				break
			case 'b':
				c.pushByte('\b')
				break
			case 'f':
				c.pushByte('\f')
				break
			case 'n':
				c.pushByte('\n')
				break
			case 'r':
				c.pushByte('\r')
				break
			case 't':
				c.pushByte('\t')
				break
			default:
				//TODO

				return ParserInvalidStringEscape
			}
			break

		default:
			c.pushByte(ch)
			//	fmt.Println("!!!!!!!!!", i, len(c.json))
			if i >= len(c.json)-1 {
				//TODO
				return ParserMissQuotationMark
			} else if ch < 0x20 {
				return ParserInvalidStringChar
			}
		}

	}
	//fmt.Print("ffffffffffffff")
	return ParserInvalidStringEscape
}

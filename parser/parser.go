package begojson

import "strconv"

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
	ParserInvalidUnicodeHex //6
	ParserInvalidSurrogate  //7
	ParserMissCommaOrSquareBracket
)

/* store info from parser*/
/*it's sad  that golang has no union type like C! */
/*why dont use interface{}  because interface{} spends more space and time*/

type begoValue struct {
	_type jsonType     //store type
	value float64      //store number in json
	str   string       //store string in json
	a     *[]begoValue //store array in json
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

/*parser string*/
func (c *context) parserString(v *begoValue) begoParserStatus {

	head := len(c.s)
	i := c.index + 1

	for ; i < len(c.json); i++ {
		ch := c.json[i]
		switch ch {

		case '"':
			size := len(c.s) - head
			str := string(c.popBytes(size))
			setString(v, str)
			c.index = i + 1
			//c.s = nil    >_< this is a bug !dont free the memory
			//when parsering array, also use this stack
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
			case 'u':
				p := []byte(c.json[i:])
				r, v := getu4(p)
				if v < 0xD800 || v > 0xDBFF {
					if r < 0 {
						return ParserInvalidStringChar
					}
					push4u(r, c)
					i += 4
				} else {
					if len(c.json) <= 12 {
						return ParserInvalidSurrogate
					}
					i += 6
					p1 := []byte(c.json[i:])

					_, v1 := getu4(p1)

					if v1 < 0xDc00 || v1 > 0xDFFF {
						return ParserInvalidSurrogate
					}
					v = (((v - 0xD800) << 10) | (v1 - 0xDC00)) + 0x10000 //it may be confusing...just convert \uhhhh\uhhhh to U+hhhhh
					push4u(rune(v), c)
					i += 4
				}

			default:
				return ParserInvalidStringEscape
			}
			break

		default:
			if i > len(c.json)-1 {
				//TODO
				c.index = i + 1
				return ParserMissQuotationMark
			} else if ch < 0x20 {
				c.index = i + 1
				return ParserInvalidStringChar
			}
			c.pushByte(ch)
		}
	}
	return ParserMissQuotationMark
}

func parser(v *begoValue, json string) begoParserStatus {
	c := context{json: json, index: 0}
	length := len(json)
	v._type = jsonNULL //initize the type

	c.parserWhiteSpace()
	ret := c.parserValue(v)
	c.parserWhiteSpace()

	if ret == ParserOk {
		c.parserWhiteSpace()
		if c.index < length {
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
	case '[':
		return c.parserArray(v)
	default:
		return c.parserNumber(v)
	}
}

/*parser array*/
func (c *context) parserArray(v *begoValue) begoParserStatus {

	c.parserWhiteSpace() //strip the space behind [
	c.index++

	if c.json[c.index] == ']' {
		c.index++
		v._type = jsonARRAY
		v.value = 0 //use value to count the elem in array
		return ParserOk
	}

	for {
		e := begoValue{}

		if c.parserValue(&e) != ParserOk {
			break
		}

		c.pushValue(e)
		c.parserWhiteSpace()

		v.value++ //elem +1

		if c.json[c.index] == ',' {
			c.index++
			c.parserWhiteSpace()
		} else if c.json[c.index] == ']' {

			c.index++
			v._type = jsonARRAY
			//copyStackToValue(c, v) //将暂存区的东西放到v.a中
			tmp := make([]begoValue, int(v.value), int(v.value))
			copy(tmp, c.popValues(int(v.value))[:])
			v.a = &tmp
			return ParserOk
		} else {
			return ParserMissCommaOrSquareBracket
		}

	}

	return ParserOk
}

func copyStackToValue(c *context, v *begoValue) {
}

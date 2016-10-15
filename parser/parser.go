package begojson



import (
	"strings"
	"strconv"
)

//import "fmt"



/* type for BegoNull ...*/
type jsonType int8

/* type for ParserOk...*/
type begoParserStatus int8

/* valueType in json*/
const (
	jsonNULL jsonType= iota
	jsonFALSE
	jsonTRUE
	jsonNUMBER
	jsonSTRING
	jsonARRAY
	jsonOBJECT
)

/* result from parser*/
const (
	ParserOk begoParserStatus= iota
	ParserExpectValue
	ParserInvalidValue
	ParserRootNotSingular
	ParserNumberTooBig
)



/* store info from parser*/
type begoValue struct {
	_type jsonType
	value float64

}

/*our context to store json file string and other things*/
type context struct {
	json  string
	index int //index for json string
	length int
}

/*skip the white space*/
func (c *context) parserWhiteSpace() {

	i := c.index
	str := c.json
	length := c.length

	if( i >=length ) {
		return
	}

	for str[i] == ' ' || str[i] == '\t' || str[i] == '\n' || str[i] == '\r' {
		i++
		if( i >=length ) {
			c.index = i
			return
		}
	}
	c.index = i
}



/*parser true false  null*/
func (c *context )parserCommon(aimStr string, v *begoValue) begoParserStatus{
	i   :=c.index
	str :=c.json
	length :=len( aimStr)

	if c.length <= i+ length -1 || strings.EqualFold(str[ i : i+length-1] , aimStr){
		return ParserInvalidValue
	}
	c.index +=length

	if  aimStr=="null"{

		v._type=jsonNULL
	}else if aimStr=="false"{

		v._type=jsonFALSE
	}else {

		v._type=jsonTRUE
	}
	return  ParserOk
}



func isDigit(ch byte) bool{

	if ch >= '0'&& ch <='9'{
		return true
	}
	return false
}

func isDigit1To9(ch byte) bool{

	if ch >= '1'&& ch <='9'{
		return true
	}
	return false
}


/*parser number*/

func(c *context) parserNumber(v *begoValue) begoParserStatus{

	i := c.index
	json :=c.json

	if i<c.length && json[i] == '-'{
		i++
	}

	if i<c.length && json[i] == '0'{
		i++
	}else{
		if i >= c.length || !isDigit1To9( json[i]){

			return ParserInvalidValue
		}

		for i=i+1;i< c.length && isDigit( json[i]) ;i++ {
		}
	}


	if i< c.length && json[i] == '.'{
		i++
		if i >= c.length || !isDigit( json[i]){
			return ParserInvalidValue
		}
		for i=i+1;i< c.length && isDigit( json[i]) ;i++ {
		}

	}


	if i<c.length && (json[i]=='e'|| json[i] =='E'){
		i++
		if i<c.length && json[i] == '+' ||json[i] == '-'{
			i++
		}

		if i >= c.length || !isDigit( json[i]){
			return ParserInvalidValue
		}
		for i=i+1;i< c.length && isDigit( json[i]) ;i++ {
		}

	}

	value, erron := strconv.ParseFloat(c.json[c.index:i],64)

	//number is  to big
	if erron!=nil {
		return ParserNumberTooBig
	}

	v.value = value
	v._type = jsonNUMBER
	c.index = i

	return ParserOk

}


/*return the status of parser*/
func (c *context) parserValue(v *begoValue) begoParserStatus {
	switch ch := c.json[c.index]; ch {
	case 'n':
		return c.parserCommon("null",v)
	case 't':
		return c.parserCommon("true",v)
	case 'f':
		return c.parserCommon("false",v)
	case '"':
		return ParserOk
	default:
		return c.parserNumber(v)
	}
}

func parser(v *begoValue, json string) begoParserStatus{
	c := context{json: json, index: 0, length:len(json)}
	v._type = jsonNULL	//initize the type

	c.parserWhiteSpace()
	ret :=c.parserValue(v)
	c.parserWhiteSpace()

	if  ret ==ParserOk{

		c.parserWhiteSpace()
		if c.index  < c.length{
			ret = ParserRootNotSingular
		}
	}

	return  ret
}

/*get the _type of begoValue*/
func getJSONType(v *begoValue) jsonType {

	if v == nil {
		panic("*begoValue cannot be nil")
	}
	return v._type
}

func getNumber(v *begoValue) float64{
	if v == nil {
		panic("*begoValue cannot be nil")
	}
	return v.value
}
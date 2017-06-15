package hmm

import "github.com/cocaer/goNLP/data"

import "unicode"

const (
	B = iota
	M
	E
	S
	SUM_STATUS
)

type Model struct {
	StartPro [SUM_STATUS]float64
	TransPro *[SUM_STATUS][SUM_STATUS]float64
	EmitPro  *[SUM_STATUS]map[rune]float64
}

func (self *Model) getEmitPro(status int, r rune) float64 {
	re := IMPOSSIBLEPRO
	if v, ok := self.EmitPro[status][r]; ok {
		re = v
	}
	return re
}

func (self *Model) Viterbi(str string) []int {
	ssrune := []rune(str)
	strLen := len(ssrune)
	var weight [][]float64
	var path [][]int
	for i := 0; i < SUM_STATUS; i++ {
		weight = append(weight, make([]float64, len(ssrune)))
		path = append(path, make([]int, len(ssrune)))
	}

	for i := 0; i < SUM_STATUS; i++ {
		weight[i][0] = self.StartPro[i] + self.getEmitPro(i, ssrune[0])
	}
	for i := 1; i < strLen; i++ {
		for j := 0; j < SUM_STATUS; j++ {
			weight[j][i] = IMPOSSIBLEPRO
			path[j][i] = j
			for k := 0; k < SUM_STATUS; k++ {
				tmp := weight[k][i-1] + self.TransPro[k][j] + self.getEmitPro(j, ssrune[i])

				if tmp > weight[j][i] {
					weight[j][i] = tmp
					path[j][i] = k
				}
			}
		}
	}
	//	result := ""
	status := SUM_STATUS - 2
	if weight[status][strLen-1] < weight[SUM_STATUS-1][strLen-1] {
		status = SUM_STATUS - 1
	}

	result := make([]int, strLen)
	result[strLen-1] = status
	for i := strLen - 1; i > 0; i-- {
		result[i-1] = path[status][i]
		status = path[status][i]
	}
	return result
}

func (self Model) Cut(s string) []string {
	status := self.Viterbi(s)
	result := make([]string, 0)
	ssrune := []rune(s)
	begin := 0
	end := 0
	processDigital(status, s)
	processLetter(status, s)
	for end < len(ssrune) {
		if status[begin] == S {
			result = append(result, string(ssrune[begin]))
			begin++
			end++
		} else {
			for end < len(ssrune) && status[end] != E {
				end++
			}
			end++
			result = append(result, string(ssrune[begin:end]))
			begin = end

		}
	}
	return result
}

func isLetter(r rune) bool {
	flag := false
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		flag = true
	}
	return flag
}

func processLetter(status []int, s string) {
	ssrune := []rune(s)
	var i = 0
	for i < len(ssrune) {
		if isLetter(ssrune[i]) {
			status[i] = B
			i++
			for i < len(ssrune) && isLetter(ssrune[i]) {
				status[i] = M
				i++
			}
			if i < len(ssrune) {
				status[i-1] = E
			} else if isLetter(ssrune[i-1]) {
				status[i-1] = E
			}
		}
		i++
	}
}
func processDigital(status []int, s string) {
	ssrune := []rune(s)
	var i = 0
	for i < len(ssrune) {
		if unicode.IsDigit(ssrune[i]) {
			status[i] = B
			i++
			for i < len(ssrune) && unicode.IsDigit(ssrune[i]) {
				status[i] = M
				i++
			}
			if i < len(ssrune) {
				status[i-1] = E
			} else if unicode.IsDigit(ssrune[i-1]) {
				status[i-1] = E
			}
		}
		i++
	}
}

func NewModel() *Model {
	m := new(Model)
	m.StartPro = data.StartProMaterix
	m.EmitPro = data.EmitProMaterix
	m.TransPro = &data.TransferMatrix
	return m
}

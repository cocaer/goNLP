package hmm

import (
	"github.com/cocaer/goNLP/data"
)

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

func (self *Model) Cut(s string) []string {
	ss := splitSentence(s)
	result := make([]string, 0)
	for _, senctence := range ss {
		result = append(result, self.CutSentence(senctence)...)
	}
	return result
}

func (self *Model) CutSentence(s string) []string {
	if len(s) == 0 {
		return nil
	}
	status := self.Viterbi(s)
	result := make([]string, 0)
	ssrune := []rune(s)
	begin := 0
	end := 0
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

func splitSentence(s string) []string {
	var ma = map[rune]byte{'。': 1, '！': 1, '，': 1, '？': 1}
	ssrune := []rune(s)
	tmp := make([]rune, 0)
	ss := make([]string, 0)
	for begin := 0; begin < len(ssrune); begin++ {
		tmp = append(tmp, ssrune[begin])
		if _, ok := ma[ssrune[begin]]; ok {
			ss = append(ss, string(tmp))
			tmp = make([]rune, 0)
		} else if begin == len(ssrune)-1 {
			ss = append(ss, string(tmp))
		}
	}
	return ss
}

func NewModel() *Model {
	m := new(Model)
	m.StartPro = data.StartProMaterix
	m.EmitPro = data.EmitProMaterix
	m.TransPro = &data.TransferMatrix
	return m
}

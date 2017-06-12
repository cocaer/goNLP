package main

import (
	"github.com/cocaer/goNLP/seg/hmm"
)

func main() {
	hmm.HmmSaveTrainingFile()
	hmm.Viterbi("小明硕士毕业于中国科学院计算所")
}

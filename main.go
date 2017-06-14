package main

import "github.com/cocaer/goNLP/seg/hmm"

func main() {
	//hmm.HmmSaveTrainingFile()
	m := hmm.NewModel()
	m.Viterbi("我恨你")

}

package main

import (
	"fmt"

	"github.com/cocaer/goNLP/seg/hmm"
)

//import "github.com/cocaer/goNLP/seg/hmm"

func main() {
	hmm.HmmSaveTrainingFile()
	m := hmm.NewModel()
	fmt.Println(m.Cut("王者荣耀是一款十分优秀的手游，最高同时两百万人在线。PS：我瞎说的"))

}

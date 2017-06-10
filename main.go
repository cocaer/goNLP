package main

import (
	"github.com/cocaer/goNLP/seg/hmm"
	"github.com/cocaer/goNLP/data"
	"fmt"
)

func main() {
	hmm.HmmSaveTrainingFile()
	fmt.Print(data.EmitProMaterix[0]['汗'])
	fmt.Println(data.StartProMaterix[1])
}

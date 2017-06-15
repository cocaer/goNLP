package main

import "github.com/cocaer/goNLP/seg/hmm"
import "fmt"

func main() {
	m := hmm.NewModel()
	fmt.Println(m.Cut("我来到USA10年了"))
}

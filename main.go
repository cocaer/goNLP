package main

import "github.com/cocaer/goNLP/seg/hmm"

func main() {
	m := hmm.NewModel()
	m.Cut("他来到了网易杭研大厦")
}

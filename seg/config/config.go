package seg

import (
	"os"
)

var goPath = os.Getenv("GOPATH")

var SegConfig = map[string]string{
	"mmDictPath":       goPath + "/src/github.com/cocaer/goNLP/data/mmdict.utf8",
	"hmmBEMSFile":      goPath + "/src/github.com/cocaer/goNLP/data/bmes.utf8",
	"hmmModelFile":     goPath + "/src/github.com/cocaer/goNLP/data/hmmmodel.go",
	"hmmTrainingFile":  goPath + "/src/github.com/cocaer/goNLP/data/msr_training.utf8",
	"bhmmBEMSFile":     goPath + "/src/github.com/cocaer/goNLP/data/bmes_merge.utf8",
	"bhmmModelFile":    goPath + "/src/github.com/cocaer/goNLP/data/bhmmmodel.go",
	"bhmmTrainingFile": goPath + "/src/github.com/cocaer/goNLP/data/msr_training.utf8",
}

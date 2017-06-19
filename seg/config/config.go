package seg

import (
	"os"
)

var goPath = os.Getenv("GOPATH")

var SegConfig = map[string]string{
	"mmDictPath":      goPath + "/src/github.com/cocaer/goNLP/data/mmdict.utf8",
	"hmmBEMSFile":     goPath + "/src/github.com/cocaer/goNLP/data/bmes.utf8",
	"hmmModelFile":    goPath + "/src/github.com/cocaer/goNLP/data/hmmmodel.go",
	"hmmTrainingFile": goPath + "/src/github.com/cocaer/goNLP/data/pku_training.utf8",
}

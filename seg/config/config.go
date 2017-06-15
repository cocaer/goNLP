package seg

import (
	"os"
)

var goPath = os.Getenv("GOPATH")

var SegConfig = map[string]string{
	"mmDictPath":      goPath + "/src/github.com/cocaer/goNLP/data/mmdict.utf8",
	"hmmTrainingFile": goPath + "/src/github.com/cocaer/goNLP/data/bmes.utf8",
	"hmmModelFile":    goPath + "/src/github.com/cocaer/goNLP/data/hmmmodel.go",
}

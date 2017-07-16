package seg

import (
	"os"
)

var goPath = os.Getenv("GOPATH")

var SegConfig = map[string]string{
	"mmDictPath":   goPath + "/src/github.com/cocaer/goNLP/data/mmdict.utf8",
	"hmmBEMSFile":  goPath + "/src/github.com/cocaer/goNLP/data/merge",
	"hmmModelFile": goPath + "/src/github.com/cocaer/goNLP/data/hmmmodel.go",
}

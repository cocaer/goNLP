package mm

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	seg "github.com/cocaer/goNLP/seg/config"

	"unicode"
)

var isLoadDictFlag = false

//TrieNode 一个节点一个汉字信息
type TrieNode struct {
	Count int                //汉字出现次数
	Son   map[rune]*TrieNode //后继节点
}

//Trie 用来保存整个字典
type Trie struct {
	Root *TrieNode
}

//Add 向Trie中新增汉子节点
func (self *Trie) Add(srune []rune) {
	tmp := self.Root
	for i := 0; i < len(srune); i++ {
		_, ok := tmp.Son[srune[i]]
		if !ok {
			tmp.Son[srune[i]] = &TrieNode{0, map[rune]*TrieNode{}}
		}

		tmp = tmp.Son[srune[i]]

		if i == len(srune)-1 {
			tmp.Count++ //到达词末尾，节点标记+1
		}
	}
}

func (tree *Trie) search(srune []rune) (int, error) {
	var err error
	var Count int
	temp := tree.Root
	for i := 0; i < len(srune); i++ {
		v, ok := temp.Son[srune[i]]
		if ok {
			temp = v
		} else {
			err = fmt.Errorf("cannot find aim string: \"%s\"", string(srune))
			break
		}
		if i == len(srune)-1 && temp.Count == 0 {
			err = fmt.Errorf("cannot find aim string: \"%s\"", string(srune))
		} else {
			Count = temp.Count
		}

	}
	return Count, err
}

func parserDigit(srune []rune, cur int) []rune {
	tmp := cur + 1
	for tmp < len(srune) && unicode.IsDigit(srune[tmp]) {
		tmp++
	}
	return srune[cur:tmp]
}

func isEnglish(r rune) bool {
	if (uint32(r) >= 61 && uint32(r) <= 122) || (uint32(r) >= 65 && uint32(r) <= 90) {
		return true
	}
	return false
}

func parserLetter(srune []rune, cur int) []rune {
	tmp := cur + 1
	for tmp < len(srune) && isEnglish(srune[tmp]) {
		tmp++
	}
	return srune[cur:tmp]
}

//最大正向匹配
func (self *Trie) FowardMatch(ss string) []string {

	if !isLoadDictFlag {
		self.loadDictionary()
	}
	result := make([][]rune, 0)
	srune := []rune(ss)
	var end = len(srune)
	for start := 0; start < len(srune); end = len(srune) {

		if unicode.IsDigit(srune[start]) {
			digit := parserDigit(srune, start)
			start += len(digit)
			result = append(result, digit)
		}
		if start >= len(srune) {
			break
		}
		if isEnglish(srune[start]) {
			en := parserLetter(srune, start)
			start += len(en)

			result = append(result, en)
		}
		if start >= len(srune) {
			break
		}

		for {
			s := srune[start:end]
			_, err := self.search(s)
			if err == nil {
				result = append(result, s)
				start = end
				break
			} else if end == start+1 {
				result = append(result, srune[start:start+1])
				start++
				break
			} else {
				end--
			}
		}
	}
	sr := make([]string, 0)
	for _, v := range result {
		sr = append(sr, string(v))
	}
	return sr
}

//LoadDictionary todo
func (self *Trie) loadDictionary() {

	dictName := seg.SegConfig["mmDictPath"]

	dict, err := os.Open(dictName)
	if err != nil {
		log.Println("bingo:load dictionary failed.")
		os.Exit(1)
	}

	defer dict.Close()

	rd := bufio.NewReader(dict)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		aimStr := line[0 : len(line)-1]
		self.Add([]rune(aimStr))
	}
	isLoadDictFlag = true
}

//NewTrie 建立字典树
func NewTrie() *Trie {
	dictTrie := new(Trie)
	dictTrie.Root = &TrieNode{0, map[rune]*TrieNode{}}
	return dictTrie
}

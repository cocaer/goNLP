# NLP工具库

## seg分词模块(MM,HMM，CRF)

- HMMmodel使用

```go
package main

import "github.com/cocaer/goNLP/seg/hmm"
import "fmt"
func main() {
    m := hmm.NewModel()
    fmt.Println(m.Cut("王者荣耀是一款十分优秀的手游，最高同时两百万人在线。PS：我瞎说的"))
}

```

``` txt
output:[他来 到 了 网易 杭研 大厦]
```


## TODO

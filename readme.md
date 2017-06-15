# NLP工具库

## seg分词模块(MM,HMM，CRF)

- HMMmodel使用

```go
package main

import "github.com/cocaer/goNLP/seg/hmm"

func main() {
    m := hmm.NewModel()
    m.Cut("他来到了网易杭研大厦")
}
```

``` txt
output:[他来 到 了 网易 杭研 大厦]
```


## TODO

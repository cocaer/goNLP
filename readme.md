# NLP工具库

## seg分词模块(MM,HMM，CRF)

- HMMmodel使用

```go
package main

import "github.com/cocaer/goNLP/seg/hmm"
import "github.com/cocaer/goNLP/seg/bhmm"
import "github.com/cocaer/goNLP/seg/mm"
import "fmt"
func main() {
    //m := mm.NewModel() //提供最大正向和逆向匹配
    m :=bhmm.NewModel() //基于hmm的优化实现
    //m := hmm.NewModel() //纯hmm实现，没有优化
    fmt.Println(m.Cut("结婚的和尚未结婚的"))
}

```

``` txt
output:[结婚 的 和 尚未 结婚 的]
```


## TODO

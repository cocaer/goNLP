package main

import (
	"fmt"

	"./seg/mm"
)

func main() {
	s := seg.NewTrie()
	r1 := s.FowardMatch("google登录后即可在您的任何设备上获取自己的书签、历史记录、密码和其他设置。此外，您还会自动登录到12345google")
	fmt.Println(r1)

}

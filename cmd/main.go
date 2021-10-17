package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	tree "word-filter"
)

func main() {
	r, err := os.Open("D:\\download\\sensitive-words-master\\sensitive-words-master\\色情类.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("ReadAll error", err)
		return
	}
	temp := strings.TrimSpace(string(data))
	temp = strings.Replace(temp, "\n", "", -1)
	temp = strings.Replace(temp, " ", "", -1)
	words := strings.Split(temp, ",")
	buildTree := tree.BuildTree(words)
	text := "有这么大嘛,aa啊，可以，奶子，阿阿阿"
	ok, w := buildTree.FindFirstMatchedWord(text)
	if ok {
		fmt.Println("发现违禁词:" + w)
	} else {
		fmt.Println("无违禁词")
	}
}

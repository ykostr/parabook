package main

import (
	"encoding/xml"
	"fmt"
	"github.com/ykostr/parabook/src/pbo"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("book")
	p := new(pbo.ParallelBook)
	xml.Unmarshal(b, p)
	for _, v := range p.Pairs {
		fmt.Println(v)
	}
	fmt.Println(p.Lang1.Value)
}

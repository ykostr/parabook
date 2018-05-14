package main

import (
	"encoding/xml"
	"fmt"
	"github.com/ykostr/parabook/src/parser"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("book")
	p := new(parser.ParallelBook)
	xml.Unmarshal(b, p)
	for _, v := range p.Tags {
		fmt.Println(v)
	}
}

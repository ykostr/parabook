package parser

import (
	"encoding/xml"
)

type ParallelBook struct {
	XMLName xml.Name `xml:"ParallelBook"`
	Tags    []Tag    `xml:"p"`
}

type Tag struct {
	L string `xml:"l,attr"`
	S string `xml:"s,attr"`
	T string `xml:"t,attr"`
}

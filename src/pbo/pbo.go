package pbo

import (
	"encoding/xml"
)

type ParallelBook struct {
	XMLName xml.Name `xml:"ParallelBook"`
	Lang1   xml.Attr `xml:"lang1,attr"`
	Author1 xml.Attr `xml:"author1,attr"`
	Title1  xml.Attr `xml:"title1,attr"`
	Info1   xml.Attr `xml:"info1,attr"`
	Lang2   xml.Attr `xml:"lang2,attr"`
	Author2 xml.Attr `xml:"author2,attr"`
	Title2  xml.Attr `xml:"title2,attr"`
	Info2   xml.Attr `xml:"info2,attr"`
	Info    xml.Attr `xml:"info,attr"`
	Pairs   []Pair   `xml:"p"`
}

type Pair struct {
	Level     string `xml:"l,attr"`
	Source    string `xml:"s,attr"`
	Translate string `xml:"t,attr"`
}

package main

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type root struct {
	XMLName xml.Name `xml:"root"`
	ChA []cha `xml:"cha"`
}
type cha struct  {
	B string `xml:"b"`
	C string `xml:"c"`
}
type chd struct  {
	XMLName xml.Name `xml:"chd"`
	F string `xml:"F"`
}

func main()  {
	//normalCheck()

	loopingCheck()

}

//func normalCheck()  {
//	r := root{}
//	xmlContent, _ := ioutil.ReadFile("data/data.xml")
//
//	err := xml.Unmarshal(xmlContent, &r)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(r.ChA.B)
//	fmt.Println(r.ChA.C)
//}

func loopingCheck() {
	r := root{}
	xmlContent, _ := ioutil.ReadFile("data/data2.xml")

	err := xml.Unmarshal(xmlContent, &r)
	if err != nil {
		panic(err)
	}

	for _,items := range r.ChA {
		fmt.Println(items.B)
		fmt.Println(items.C)
	}


}
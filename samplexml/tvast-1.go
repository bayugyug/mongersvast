package main

import (
	"fmt"

	mvast "github.com/bayugyug/mongersvast"
)

func main() {

	var xml, body string

	//LOAD from a sample xml string
	body = mvast.XMLInlineLinear
	vt := mvast.VAST{}
	vt.FromString(body)
	xml, _ = vt.ToString()
	fmt.Println(xml)

}

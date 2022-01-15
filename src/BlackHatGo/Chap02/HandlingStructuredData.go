package main

import (
	"fmt"
	_ "encoding/json"
	"encoding/xml"
)
/*
// json
type Foo struct {
	Bar string
	Baz string
}
*/
// xml
type Foo struct {
	Bar string	`xml: "id, attr"`
	Baz string	`xml: "parent>child"`
}


func main() {

	f := Foo {"Joe Junior", "Hello Shabado"}
//	b, _ := json.Marshal(f)
	b, _ := xml.Marshal(f)
	fmt.Println(string(b))
//	json.Unmarshal(b, &f)
	xml.Unmarshal(b, &f)
}

package main

import (
	"encoding/json"
	"fmt"
)

/*
The '-' says to not include this field in json
Run this to see what happens with 'json:"something else"'
 */
type Person struct {
	First string
	Last string `json:"-"`
	Age int `json:"something else"`
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}

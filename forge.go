package main

import "fmt"
import "encoding/json"

func main() {
	res, err := ParseFile("test.fg")
	if err != nil {
		panic(err)
	}

	d, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(d[:]))
}

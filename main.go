package main

import (
	"encoding/json"
	"fmt"

	"github.com/tyler-johnson/forge/forge/fcl"
)

func main() {
	file, err := fcl.ParseFile("./forge/fcl/example/user.fg")
	if err != nil {
		panic(err)
	}

	o, _ := json.MarshalIndent(file, "", "  ")
	fmt.Println(string(o[:]))
}

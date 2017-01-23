package main

import (
	"os"

	"github.com/tyler-johnson/forge/forge/fcl"
	"github.com/tyler-johnson/forge/forge/fcl/printer"
)

const test = `
jfjksdfj fjsiefji sgeisef # blah
sgheiefsifesife {} # my awesome comment
foo bar # my comment
# test
sfejisi seifjsiefsef {}
sfs sefsiejfs
sfeogesge
sefsekffs eeigigig

fsdfsdffsef sefsef
asgege seseg {
	sfesff
}
foo bar
`

func main() {
	// tree, err := fcl.ParseFile("./forge/fcl/example/user.fg")
	tree, err := fcl.Parse([]byte(test))
	if err != nil {
		panic(err)
	}

	result, err := printer.Print(tree)
	if err != nil {
		panic(err)
	}

	outfile, err := os.Create("./user_result.fg")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	_, err = outfile.Write(result)
	if err != nil {
		panic(err)
	}

	// o, _ := json.MarshalIndent(file, "", "  ")
	// fmt.Println(string(o[:]))
}

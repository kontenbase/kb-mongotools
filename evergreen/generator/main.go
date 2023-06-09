package main

import (
	"fmt"

	"github.com/kontenbase/kb-mongotools/evergreen"
)

func main() {
	c, err := evergreen.Load()
	if err != nil {
		panic(err)
	}

	y, err := c.GitHubPRAliasesYAML()
	if err != nil {
		panic(err)
	}

	fmt.Println(y)
}

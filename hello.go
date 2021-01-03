package main

import (
	"fmt"
	"github.com/ezchi/go-hello-world/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fnt.Println("That us start.")
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

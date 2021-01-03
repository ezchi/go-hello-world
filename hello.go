package main

import "fmt"
import "github.com/ezchi/go-hello-world/morestrings"
import "github.com/google/go-cmp/cmp"

func main() {
	fmt.Println("That us start.")
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

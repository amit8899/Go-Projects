// How to use struct that is imported from another package?

package main

import (
	parent "family/father"
	child "family/father/son"
	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("Mr. Amit Rai"))

	c := new(child.Son)
	fmt.Println(c.Data("Riley"))

}

// Download the VS code go extension, make sure to install all the tools after it yells at you

package main

import (
	"fmt"
	"os"
	"github.com/lucjohnson/info344-class-code/gohello/reverse"
)

func main() {
	fmt.Println(reverse.Reverse("Hello World!"))
	fmt.Println(os.Args[0])
}
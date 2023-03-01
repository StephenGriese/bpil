package main

import (
	"fmt"

	"github.com/StephenGriese/kitty/kitty"
	"github.com/StephenGriese/pufs/pufs"
)

func main() {
	fmt.Printf("kitty value: %v\n", kitty.KittyValue)
	fmt.Printf("pufs returns: %v\n", pufs.Pufs())

}

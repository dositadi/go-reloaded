package main

import (
	c "edit-tool/internal/config"
	h "edit-tool/pkg/utils"
	"fmt"
)

func main() {
	app := c.App{}

	app.Run()

	name := "divine"
	h.CapIndices("understanding", 6)

	fmt.Println(name)
}

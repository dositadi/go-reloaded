package main

import (
	c "edit-tool/internal/config"
	h "edit-tool/pkg/utils"
	"fmt"
)

func main() {
	app := c.App{}

	app.Run()

	val := ",and"

	val2 := "pro"

	fmt.Println(h.ContainsPunctuation("boring"), val)
	h.AppendChar(&val2, ',')
	fmt.Println(val2)
}

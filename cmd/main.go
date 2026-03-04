package main

import (
	c "edit-tool/internal/config"
)

//  This is where the edit tool engine begins work
func main() {
	app := c.App{}

	app.Run()
}

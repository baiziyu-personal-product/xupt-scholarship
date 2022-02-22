package main

import (
	"github.com/fatih/color"
	"xupt-scholarship/controllers"
	"xupt-scholarship/initialize"
)

func main() {
	initialize.Init()

	color.Green("Serve config has been emitted!!!")
	controllers.Router()
}

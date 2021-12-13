package main

import (
	"xupt-scholarship/controllers"
	"xupt-scholarship/initialize"
)

func main() {
	initialize.InitServeConfig()
	controllers.Router()
}

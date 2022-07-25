package main

import (
	app_generic "go-learning/internal/app/generic_implementation"
	app_interface "go-learning/internal/app/interface_implementation"
)

func main() {
	listInterface := app_interface.InitList()
	listInterface.Add("5")
	listGeneric := app_generic.InitList[string]()
	listGeneric.Add("5")
}

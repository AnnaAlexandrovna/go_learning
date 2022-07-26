package main

import (
	appGeneric "go-learning/internal/app/genericImplementation"
	appInterface "go-learning/internal/app/interfaceImplementation"
)

func main() {
	listInterface := appInterface.InitList()
	listInterface.Add("5")
	listGeneric := appGeneric.InitList[string]()
	listGeneric.Add("5")
}

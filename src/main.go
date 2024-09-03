/*
Copyright © 2024 Antoine Harle antoine.harle@proton.me

*/
package main

import (
	"github.com/AntoineHX/multi-motors-controller/src/cmd"
	_ "github.com/AntoineHX/multi-motors-controller/src/cmd/controller"
	_ "github.com/AntoineHX/multi-motors-controller/src/cmd/motor"
)

func main() {
	cmd.Execute()
}

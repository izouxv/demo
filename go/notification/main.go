package main

import "notification/cmd"

var version string // set by the compiler

func main() {
	cmd.Execute(version)
}


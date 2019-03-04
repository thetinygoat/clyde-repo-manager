package main

import "os"

// get command line args
func getCLA() []string {
	cla := os.Args[1:]
	return cla
}

// initialize the cli handler
func init() {
	cla := getCLA()
	switch cla[0] {
	case "install":
		install()
	case "update":
		update()
	case "remove":
		remove()
	case "list":
		list()
	default:
		exception(cla[0])
	}
}

package main

import (
	"fmt"
	"os"
)

func install() error {
	fmt.Println("install function called")
	return nil
}

func remove() error {
	fmt.Println("remove function called")
	return nil
}

func update() error {
	fmt.Println("update function called")
	return nil
}
func list() error {
	fmt.Println("list function called")
	return nil
}

func exception(cmd string) {
	fmt.Printf("'%s' command not found\n", cmd)
	os.Exit(1)
}

func main() {
}

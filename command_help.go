package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, command := range getCommands() {
		fmt.Printf("%v: ", command.name)
		fmt.Printf("%v\n", command.description)
	}
	return nil
}

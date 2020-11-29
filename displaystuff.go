package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func displayMenu() {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Println("Reddit Army Menu")
	fmt.Println("---------------------")
	fmt.Printf("(%s)how Soldiers (%s)dd Soldier (%s)ink (%s)uit\n", yellow("S"), yellow("A"), yellow("L"), red("Q"))
	fmt.Printf(">> ")
	takeInput()

}

func displaySoldierDatabase(soldiers []string) {

	fmt.Println("# DEV USER ID    | DEV SECRET KEY              | REDDIT USERNAME  | REDDIT PASSWORD")
	fmt.Println("-----------------------------------------------------------------------------------")

	for i, s := range soldiers {
		fmt.Printf("%v ", i+1)
		c := strings.Split(s, ",")
		for x, title := range c {
			fmt.Printf(title)
			if x != 3 {
				fmt.Print(" | ")
			} else {
				fmt.Println()
			}

		}
	}

	fmt.Println()

	displayMenu()

}

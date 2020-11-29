package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var soldiers = openDatabase("botarmy_database.txt")

func main() {

	displaySnoo()

	// extracts API info and reddit login data from a plain text database document into a string splice

	color.Yellow("opening army database and recruiting soldiers...")

	if len(soldiers) == 1 {

		color.Cyan("...one soldier was added to your Reddit bot army.  time to upgrade for more recruits!")
		fmt.Println()

	} else {
		color.Cyan("...a total of %v soldiers were added to your Reddit bot army.\n", len(soldiers))
	}

	displayMenu()

}

func takeInput() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	switch char {
	case 'A', 'a':
		color.Cyan("Adding soldier to database...")
		break
	case 'S', 's':
		color.Cyan("Showing soldier database...")
		displaySoldierDatabase(soldiers)

		break
	case 'Q', 'q':
		os.Exit(0)

	default:
		fmt.Println("I'm sorry, but this is not a valid entry.  Please try again.")
		displayMenu()
	}

}

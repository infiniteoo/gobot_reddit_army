package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aggrolite/geddit"
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

func authRedditSession(userID string, userSecret string, username string, password string) {

	o, err := geddit.NewOAuthSession(
		userID,
		userSecret,
		"Testing Geddit Bot with OAuth v0.1 by u/aggrolite - see source @ github.com/aggrolite/geddit/master",
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Login using our personal reddit account.
	err = o.LoginAuth(username, password)
	if err != nil {
		log.Fatal(err)
	}
	// We can pass options to the query if desired (blank for now).
	opts := geddit.ListingOptions{}

	// Fetch posts from r/videos, sorted by Hot.
	posts, err := o.SubredditSubmissions("videos", geddit.HotSubmissions, opts)
	if err != nil {
		log.Fatal(err)
	}
	// Save each post linking to youtube.com.
	for _, p := range posts {
		if p.Domain == "youtube.com" {
			// Save under category name "videos".
			err = o.Save(p, "videos")
			if err != nil {
				// Log any error, but keep going.
				log.Printf("Error! Problem saving submission: %v", err)
			}

			// this block here will upvote.
			/* err = o.Vote(p, "1")
			if err != nil {
				// Log any error, but keep going.
				log.Printf("Error! Problem upvoting submission: %v", err)
			} */
		}
	}

}

func openDatabase(txt string) []string {

	f, err := os.Open(txt)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var soldiers []string

	for scanner.Scan() {

		soldiers = append(soldiers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

		for i, s := range soldiers {
			fmt.Println(i, s)

		}

	}

	return soldiers

}

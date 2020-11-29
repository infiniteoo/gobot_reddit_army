package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

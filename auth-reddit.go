package main

import (
	"log"

	"github.com/aggrolite/geddit"
)

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

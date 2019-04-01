package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, ""+
			`calendar downloads and extrasts the list of fixtures.
Then creates, in google calendar, one event by fixture.

`)
		flag.PrintDefaults()
	}
	address := flag.String("address", "", "address of official site")
	number := flag.Int("number", 22, "fixtures count")
	team := flag.String("team", "", "specific team")
	league := flag.String("league", "N2", "league team")
	calendarId := flag.String("calendar", "primary", "calendar identifier")
	flag.Parse()

	dom, err := extractDom(*address, *number)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	fixtures, err := extractFixtures(*team, dom)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	srv, err := createCalendarService()
	if err != nil {
		log.Fatalf("Unable to create calendar service: %v", err)
	}

	err = srv.createEvents(*league, *calendarId, fixtures)
	if err != nil {
		log.Fatalf("Unable to create events: %v", err)
	}
	fmt.Println("Succeed !")
}

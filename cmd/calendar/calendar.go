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
	url := flag.String("url", "", "official website of ffrs")
	count := flag.Int("fixtures", 22, "fixtures count")
	team := flag.String("team", "", "team")
	league := flag.String("league", "N2", "league")
	calendarId := flag.String("calendar", "primary", "iCal identifier")
	flag.Parse()

	dom, err := extractDom(*url, *count)
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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"rinkgo/parsers"
	"strconv"
	"strings"
	"time"
)

//// << following code https://developers.google.com/calendar/quickstart/go

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

//// >>

type Servirce struct {
	calendar *calendar.Service
}

func createCalendarService() (*Servirce, error) {

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to read client secret file: %v", err))
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to parse client secret file to config: %v", err))
	}

	client := getClient(config)
	srv, err := calendar.New(client)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve Calendar client: %v", err))
	}

	return &Servirce{
		calendar: srv,
	}, nil
}

func convertDates(date string) (time.Time, time.Time, error) {
	result := strings.Split(date, "/")
	if len(result) != 3 {
		return time.Time{}, time.Time{},
			errors.New(fmt.Sprintf("Unable to parse date: %s", date))
	}
	y, err := strconv.ParseInt(result[2], 10, 64)
	if err != nil {
		return time.Time{}, time.Time{},
			errors.New(fmt.Sprintf("Unable to parse date: %s", date))
	}
	m, err := strconv.ParseInt(result[1], 10, 64)
	if err != nil {
		return time.Time{}, time.Time{},
			errors.New(fmt.Sprintf("Unable to parse date: %s", date))
	}
	d, err := strconv.ParseInt(result[0], 10, 64)
	if err != nil {
		return time.Time{}, time.Time{},
			errors.New(fmt.Sprintf("Unable to parse date: %s", date))
	}
	start := time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.UTC)
	return start, start.Add(time.Hour * 24), nil
}

func (s *Servirce) createEvents(league, calendarId string, fixtures []parsers.Fixture) error {
	for _, v := range fixtures {
		start, end, err := convertDates(v.Date)
		if err != nil {
			return err
		}
		event := &calendar.Event{
			Summary:     v.HomeTeam + " - " + v.OutTeam,
			Description: "Match " + league,
			Start: &calendar.EventDateTime{
				DateTime: start.Format("2006-01-02T15:04:05"),
				TimeZone: "Europe/Paris",
			},
			End: &calendar.EventDateTime{
				DateTime: end.Format("2006-01-02T15:04:05"),
				TimeZone: "Europe/Paris",
			},
		}
		_, err = s.calendar.Events.Insert(calendarId, event).Do()
		if err != nil {
			return err
		}
	}
	return nil
}

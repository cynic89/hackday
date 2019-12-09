package hobbybuddy

import (
	"fmt"
	"google.golang.org/api/calendar/v3"
	"log"
	"time"
)

func CreateEvent(hobbyName string, attendees []string, srv *calendar.Service) {
	fmt.Println("Creating Event")
	event := &calendar.Event{
		Summary:     fmt.Sprintf("%s buddies", hobbyName),
		Location:    "3495 Deer Creek Road, Palo Alto",
		Description: fmt.Sprintf("A chance to meet some your colleagues who like %s", hobbyName),
		Start: &calendar.EventDateTime{
			DateTime: "2019-12-09T15:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2019-12-09T15:30:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		Recurrence: []string{"RRULE:FREQ=WEEKLY;COUNT=1"},
		Attendees:  getAttendeesList(attendees),
	}

	calendarId := "primary"
	event, err := srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}

	fmt.Printf("Event created: %s\n", event.HtmlLink)
}

func ListHobbies(srv *calendar.Service) ([]*calendar.Event, error) {
	fmt.Println("Listing Events")
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		return nil, fmt.Errorf("No upcoming events found.")
	} else {
		return events.Items, nil
	}
}

func getAttendeesList(attendees []string) (eventAttendees []*calendar.EventAttendee) {
	for _, attendee := range attendees {
		eventAttendees = append(eventAttendees, &calendar.EventAttendee{Email: attendee})
	}
	return eventAttendees
}

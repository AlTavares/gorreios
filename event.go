package gorreios

import (
	"io"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/AlTavares/gopretty"

	"github.com/andrewstuart/goq"
)

type Tracking struct {
	ID     string `goquery:".codSro"`
	Events Events `goquery:"tr"`
}

type Events []Event

type Event struct {
	Description string `goquery:"td.sroLbEvent"`
	DateRaw     string `goquery:"td.sroDtEvent"`
	DateTime    time.Time
	City        string
}

func (t Tracking) LastEvent() Event {
	return t.Events[0]
}

func parseTracking(body io.Reader) (tracking Tracking, err error) {
	err = goq.NewDecoder(body).Decode(&tracking)
	if err != nil {
		log.Println(err)
	}
	tracking.Events = cleanEvents(tracking.Events)
	gopretty.Print(tracking)
	return
}

func cleanEvents(events Events) Events {
	for i, event := range events {
		events[i] = cleanEvent(event)
	}
	return events
}

var timeLocation, _ = time.LoadLocation("America/Sao_Paulo")

func cleanEvent(event Event) Event {
	event.Description = cleanString(event.Description)
	event.DateRaw = cleanString(event.DateRaw)
	dateInfo := strings.SplitN(event.DateRaw, " ", 3)
	event.DateTime, _ = time.ParseInLocation("02/01/2006 15:04", dateInfo[0]+" "+dateInfo[1], timeLocation)
	event.City = strings.ToUpper(dateInfo[2])

	return event
}

func cleanString(str string) string {
	var re = regexp.MustCompile(`\n|\t`)
	str = re.ReplaceAllString(str, "")
	re = regexp.MustCompile(` +`)
	str = re.ReplaceAllString(str, " ")
	re = regexp.MustCompile(`\s/\s|\xA0/\xA0`)
	str = re.ReplaceAllString(str, "/")
	return str
}

package main

import (
	"github.com/snamber/date-clipper/clipboard"
	"runtime"
	"time"

	"github.com/integrii/flaggy"
)

var style string

func main() {
	flaggy.String(&style, "s", "style", "the formatting style ('iso' / 'long-date' / 'time')")
	flaggy.Parse()

	now := time.Now()
	datestring := getDateString(now, style)

	clipboard.Copy([]byte(datestring), runtime.GOOS)
}

func getDateString(now time.Time, style string) string {
	var datestring string
	switch style {
	case "iso":
		datestring = now.Format("2006-01-02")
	case "long-date":
		datestring = now.Format("January 2, 2006")
	case "time":
		datestring = now.Format("3:04 PM")
	case "utc":
		datestring = now.In(time.UTC).Format("15:04 UTC")
	default:
		datestring = now.String()
	}
	return datestring
}


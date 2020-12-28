package templates

import (
	"strconv"
	"time"
)

// date
//
// Format date by `interface{}` input, a date can be
// a `time.Time` or an `int, int32, int64`.
//
// Example: {{ date "02/01/2006" now }}
func (t *TemplateManager) date(fmt string, date interface{}) string {
	return t.dateInZone(fmt, date, "Local")
}

// dateInZone
//
// Takes a format, the date and zone. Casts the
// date to the correct format.
//
// Example: {{ dateInZone "02/01/2006" now "Europe/London" }}
func (t *TemplateManager) dateInZone(format string, date interface{}, zone string) string {
	var tm time.Time

	switch date := date.(type) {
	default:
		tm = time.Now()
	case time.Time:
		tm = date
	case *time.Time:
		tm = *date
	case int64:
		tm = time.Unix(date, 0)
	case int:
		tm = time.Unix(int64(date), 0)
	case int32:
		tm = time.Unix(int64(date), 0)
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}

	return tm.In(loc).Format(format)
}

// ago
//
// Returns a duration from the given time input
// in seconds. a date can be a `time.Time` or
// an `int, int64`.
//
// Example: {{ ago .UpdatedAt }}
func (t *TemplateManager) ago(i interface{}) string {
	var tm time.Time

	switch date := i.(type) {
	default:
		tm = time.Now()
	case time.Time:
		tm = date
	case int64:
		tm = time.Unix(date, 0)
	case int:
		tm = time.Unix(int64(date), 0)
	}

	return time.Since(tm).Round(time.Second).String()
}

// duration
//
// Formats a given amount of seconds as a `time.Duration`
// For example `duration 85` will return `1m25s`.
//
// Example: {{ duration 85 }}
func (t *TemplateManager) duration(sec interface{}) string {
	var n int64
	switch value := sec.(type) {
	default:
		n = 0
	case string:
		n, _ = strconv.ParseInt(value, 10, 64)
	case int64:
		n = value
	}
	return (time.Duration(n) * time.Second).String()
}

// htmlDate
//
// Format's a date for inserting into a HTML date
// picker input field.
//
// Example: {{ htmlDate now }}
func (t *TemplateManager) htmlDate(date interface{}) string {
	return t.dateInZone("2006-01-02", date, "Local")
}

// htmlDateInZone
//
// Returns HTML date with a timezone
//
// Example: {{ htmlDateInZone now "Europe/London" }}
func (t *TemplateManager) htmlDateInZone(date interface{}, zone string) string {
	return t.dateInZone("2006-01-02", date, zone)
}

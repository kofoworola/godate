package godate

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type GoDate struct {
	Time     time.Time
	TimeZone *time.Location
}

//IsBefore checks if the GoDate is before the passed GoDate
func (d *GoDate) IsBefore(compare *GoDate) bool {
	return d.Time.Before(compare.Time)
}

//IsAfter checks if the GoDate is before the passed GoDate
func (d *GoDate) IsAfter(compare *GoDate) bool {
	return d.Time.After(compare.Time)
}

//Sub subtracts the 'count' from the GoDate using the unit passed
func (d GoDate) Sub(count int, unit Unit) *GoDate {
	return d.Add(-count, unit)
}

//Add adds the 'count' from the GoDate using the unit passed
func (d GoDate) Add(count int, unit Unit) *GoDate {
	d.Time = d.Time.Add(time.Duration(unit * Unit(count)))
	return &d
}

//Get the difference as a duration type
func (d *GoDate) DifferenceAsDuration(compare *GoDate) time.Duration {
	return d.Time.Sub(compare.Time)
}

//Difference Returns the difference between the Godate and another in the specified unit
//If the difference is negative then the 'compare' date occurs after the date
//Else it occurs before the the date
func (d GoDate) Difference(compare *GoDate, unit Unit) int {
	difference := d.DifferenceAsFloat(compare, unit)
	return int(difference)
}

//Get the difference as a float
func (d GoDate) DifferenceAsFloat(compare *GoDate, unit Unit) float64 {
	duration := d.DifferenceAsDuration(compare)
	return float64(duration) / float64(time.Duration(unit))
}

//Gets the difference between the relative to the date value in the form of
//1 month before
//1 month after
func (d GoDate) DifferenceForHumans(compare *GoDate) string {
	differenceString, differenceInt := d.AbsDifferenceForHumans(compare)
	if differenceInt > 0 {
		return differenceString + " before"
	} else {
		return differenceString + " after"
	}
}

//Gets the difference between the relative to current time value in the form of
//1 month ago
//1 month from now
func (d GoDate) DifferenceFromNowForHumans() string {
	now := Now(d.TimeZone)
	differenceString, differenceInt := now.AbsDifferenceForHumans(&d)
	if differenceInt > 0 {
		return differenceString + " ago"
	} else {
		return differenceString + " from now"
	}
}

//Get the abs difference relative to compare time in the form
//1 month
//2 days
func (d GoDate) AbsDifferenceForHumans(compare *GoDate) (string, int) {
	sentence := make([]string, 2, 2)
	duration := Unit(math.Abs(float64(d.DifferenceAsDuration(compare))))
	var unit Unit
	if duration >= YEAR {
		unit = YEAR
	} else if duration < YEAR && duration >= MONTH {
		unit = MONTH
	} else if duration < MONTH && duration >= WEEK {
		unit = WEEK
	} else if duration < WEEK && duration >= DAY {
		unit = DAY
	} else if duration < DAY && duration >= HOUR {
		unit = HOUR
	} else if duration < HOUR && duration >= MINUTE {
		unit = MINUTE
	} else {
		unit = SECOND
	}
	difference := d.Difference(compare, unit)
	sentence[0] = strconv.Itoa(int(math.Abs(float64(difference))))
	sentence[1] = unit.String()
	if difference == 1 || difference == -1 {
		sentence[1] = strings.TrimSuffix(sentence[1], "s")
	}
	return strings.Join(sentence, " "), difference
}

func (d *GoDate) StartOfHour() *GoDate {
	y, m, day := d.Time.Date()
	return &GoDate{time.Date(y, m, day, d.Time.Hour(), 0, 0, 0, d.TimeZone), d.TimeZone}
}

func (d *GoDate) StartOfDay() *GoDate {
	y, m, day := d.Time.Date()
	return &GoDate{time.Date(y, m, day, 0, 0, 0, 0, d.TimeZone), d.TimeZone}
}

func (d *GoDate) StartOfWeek() *GoDate {
	day := d.StartOfDay().Time.Weekday()
	if day != FirstDayOfWeek {
		return d.Sub(int(day-FirstDayOfWeek), DAY).StartOfDay()
	} else{
		return d.StartOfDay()
	}
}

func (d *GoDate) StartOfMonth() *GoDate {
	y, m, _ := d.Time.Date()
	return &GoDate{time.Date(y, m, 1, 0, 0, 0, 0, d.TimeZone), d.TimeZone}
}

func (d *GoDate) StartOfQuarter() *GoDate {
	startMonth := d.StartOfMonth()
	off := (startMonth.Time.Month() - 1) % 3
	return startMonth.Sub(int(off), MONTH)
}

func (d *GoDate) StartOfYear() *GoDate {
	y, _, _ := d.Time.Date()
	return &GoDate{time.Date(y, 1, 1, 0, 0, 0, 0, d.TimeZone), d.TimeZone}
}

func (d *GoDate) EndOfHour() *GoDate {
	nextHour := d.StartOfHour().Add(1, HOUR)
	return &GoDate{nextHour.Time.Add(-time.Millisecond), d.TimeZone}
}

func (d *GoDate) EndOfDay() *GoDate {
	nextDay := d.StartOfDay().Add(1, DAY)
	return &GoDate{nextDay.Time.Add(-time.Millisecond), d.TimeZone}
}

func (d *GoDate) EndOfWeek() *GoDate {
	nextWeek := d.StartOfWeek().Add(1, WEEK)
	return &GoDate{nextWeek.Time.Add(-time.Millisecond), d.TimeZone}
}

func (d *GoDate) EndOfMonth() *GoDate {
	nextWeek := d.StartOfMonth().Add(1, MONTH)
	return &GoDate{nextWeek.Time.Add(-time.Millisecond), d.TimeZone}
}

func (d *GoDate) EndOfQuarter() *GoDate {
	nextWeek := d.StartOfQuarter().Add(3, MONTH)
	return &GoDate{nextWeek.Time.Add(-time.Millisecond), d.TimeZone}
}

func (d *GoDate) EndOfYear() *GoDate {
	nextWeek := d.StartOfYear().Add(1, MONTH)
	return &GoDate{nextWeek.Time.Add(-time.Millisecond), d.TimeZone}
}

//MidDay gets the midday time usually 12:00 PM of the current day
func (d *GoDate) MidDay() *GoDate{
	y, m, day := d.Time.Date()
	return &GoDate{time.Date(y, m, day, 12, 0, 0, 0, d.TimeZone), d.TimeZone}
}

//ToDateTimeString Formats and returns the GoDate in the form 2006-01-02 15:04:05
func (d *GoDate) ToDateTimeString() string{
	return d.Format("2006-01-02 15:04:05")
}

//ToDateString Formats and returns the GoDate in the form 2006-01-02
func (d *GoDate) ToDateString() string{
	return d.Format("2006-01-02")
}

//ToFormattedDateString Formats and returns the GoDate in the form Jan 02, 2006
func (d *GoDate) ToFormattedDateString() string{
	return d.Format("Jan 02, 2006")
}

//ToTimeString Formats and returns the GoDate in the form 15:04:05
func (d *GoDate) ToTimeString() string{
	return d.Format("15:04:05")
}

//ToDayTimeString Formats and returns the GoDate in the form Mon, Jan 2, 2006 03:04 PM
func (d *GoDate) ToDayTimeString() string{
	return d.Format("Mon, Jan 2, 2006 03:04 PM")
}

//Check if this is the weekend
func (d *GoDate) IsWeekend() bool {
	day := d.Time.Weekday()
	return day == time.Saturday || day == time.Sunday
}

func (d *GoDate) Format(format string) string{
	return d.Time.Format(format)
}

func (d GoDate) String() string{
	return d.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
}

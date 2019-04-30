package godate

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type goDate struct {
	Time     time.Time
	TimeZone *time.Location
	FirstDayOfWeek time.Weekday
}

//IsBefore checks if the goDate is before the passed goDate
func (d *goDate) IsBefore(compare *goDate) bool {
	return d.Time.Before(compare.Time)
}

//IsAfter checks if the goDate is before the passed goDate
func (d *goDate) IsAfter(compare *goDate) bool {
	return d.Time.After(compare.Time)
}

//Sub subtracts the 'count' from the goDate using the unit passed
func (d goDate) Sub(count int, unit Unit) *goDate {
	return d.Add(-count, unit)
}

//Add adds the 'count' from the goDate using the unit passed
func (d goDate) Add(count int, unit Unit) *goDate {
	d.Time = d.Time.Add(time.Duration(unit * Unit(count)))
	return &d
}

//Get the difference as a duration type
func (d *goDate) DifferenceAsDuration(compare *goDate) time.Duration {
	return d.Time.Sub(compare.Time)
}

//Difference Returns the difference between the Godate and another in the specified unit
//If the difference is negative then the 'compare' date occurs after the date
//Else it occurs before the the date
func (d goDate) Difference(compare *goDate, unit Unit) int {
	difference := d.DifferenceAsFloat(compare, unit)
	return int(difference)
}

//Get the difference as a float
func (d goDate) DifferenceAsFloat(compare *goDate, unit Unit) float64 {
	duration := d.DifferenceAsDuration(compare)
	return float64(duration) / float64(time.Duration(unit))
}

//Gets the difference between the relative to the date value in the form of
//1 month before
//1 month after
func (d goDate) DifferenceForHumans(compare *goDate) string {
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
func (d goDate) DifferenceFromNowForHumans() string {
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
func (d goDate) AbsDifferenceForHumans(compare *goDate) (string, int) {
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

func (d *goDate) StartOfHour() *goDate {
	y, m, day := d.Time.Date()
	return &goDate{time.Date(y, m, day, d.Time.Hour(), 0, 0, 0, d.TimeZone), d.TimeZone,0}
}

func (d *goDate) StartOfDay() *goDate {
	y, m, day := d.Time.Date()
	return &goDate{time.Date(y, m, day, 0, 0, 0, 0, d.TimeZone), d.TimeZone,0}
}

func (d *goDate) StartOfWeek() *goDate {
	day := d.StartOfDay().Time.Weekday()
	if firstDay := d.FirstDayOfWeek; day != firstDay {
		return d.Sub(int(day-firstDay), DAY).StartOfDay()
	} else{
		return d.StartOfDay()
	}
}

func (d *goDate) StartOfMonth() *goDate {
	y, m, _ := d.Time.Date()
	return &goDate{time.Date(y, m, 1, 0, 0, 0, 0, d.TimeZone), d.TimeZone,0}
}

func (d *goDate) StartOfQuarter() *goDate {
	startMonth := d.StartOfMonth()
	off := (startMonth.Time.Month() - 1) % 3
	return startMonth.Sub(int(off), MONTH)
}

func (d *goDate) StartOfYear() *goDate {
	y, _, _ := d.Time.Date()
	return &goDate{time.Date(y, 1, 1, 0, 0, 0, 0, d.TimeZone), d.TimeZone,0}
}

func (d *goDate) EndOfHour() *goDate {
	nextHour := d.StartOfHour().Add(1, HOUR)
	return &goDate{nextHour.Time.Add(-time.Millisecond), d.TimeZone,0}
}

func (d *goDate) EndOfDay() *goDate {
	nextDay := d.StartOfDay().Add(1, DAY)
	return &goDate{nextDay.Time.Add(-time.Millisecond), d.TimeZone,0}
}

func (d *goDate) EndOfWeek() *goDate {
	nextWeek := d.StartOfWeek().Add(1, WEEK)
	return &goDate{nextWeek.Time.Add(-time.Millisecond), d.TimeZone,0}
}

func (d *goDate) EndOfMonth() *goDate {
	nextWeek := d.StartOfMonth().Add(1, MONTH)
	return &goDate{nextWeek.Time.Add(-time.Millisecond), d.TimeZone,0}
}

func (d *goDate) EndOfQuarter() *goDate {
	nextWeek := d.StartOfQuarter().Add(3, MONTH)
	return &goDate{nextWeek.Time.Add(-time.Millisecond), d.TimeZone,0}
}

func (d *goDate) EndOfYear() *goDate {
	nextWeek := d.StartOfYear().Add(1, MONTH)
	return &goDate{nextWeek.Time.Add(-time.Millisecond), d.TimeZone,0}
}

//MidDay gets the midday time usually 12:00 PM of the current day
func (d *goDate) MidDay() *goDate {
	y, m, day := d.Time.Date()
	return &goDate{time.Date(y, m, day, 12, 0, 0, 0, d.TimeZone), d.TimeZone,0}
}

//ToDateTimeString Formats and returns the goDate in the form 2006-01-02 15:04:05
func (d *goDate) ToDateTimeString() string{
	return d.Format("2006-01-02 15:04:05")
}

//ToDateString Formats and returns the goDate in the form 2006-01-02
func (d *goDate) ToDateString() string{
	return d.Format("2006-01-02")
}

//ToFormattedDateString Formats and returns the goDate in the form Jan 02, 2006
func (d *goDate) ToFormattedDateString() string{
	return d.Format("Jan 02, 2006")
}

//ToTimeString Formats and returns the goDate in the form 15:04:05
func (d *goDate) ToTimeString() string{
	return d.Format("15:04:05")
}

//ToDayTimeString Formats and returns the goDate in the form Mon, Jan 2, 2006 03:04 PM
func (d *goDate) ToDayTimeString() string{
	return d.Format("Mon, Jan 2, 2006 03:04 PM")
}

//Check if this is the weekend
func (d *goDate) IsWeekend() bool {
	day := d.Time.Weekday()
	return day == time.Saturday || day == time.Sunday
}

func (d *goDate) Format(format string) string{
	return d.Time.Format(format)
}

func (d *goDate) SetFirstDay(weekday time.Weekday){
	d.FirstDayOfWeek = weekday
}

func (d goDate) String() string{
	return d.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
}

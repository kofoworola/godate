package godate

import (
	"math"
	"strconv"
	"strings"
	"time"
)

//TODO add timezone support
type GoDate struct {
	Time time.Time
}

//IsBefore checks if the GoDate is before the passed GoDate
func (d GoDate) IsBefore(compare *GoDate) bool {
	return d.Time.Before(compare.Time)
}

//IsAfter checks if the GoDate is before the passed GoDate
func (d GoDate) IsAfter(compare *GoDate) bool {
	return d.Time.After(compare.Time)
}

//Sub subtracts the 'count' from the GoDate using the unit passed
func (d GoDate) Sub(count int, unit int) *GoDate {
	return d.Add(-count, unit)
}

//Add adds the 'count' from the GoDate using the unit passed
func (d GoDate) Add(count int, unit int) *GoDate {
	//milliSecondOffset := time.Millisecond * time.Duration(count/int(math.Abs(float64(count))))
	switch unit {
	case MINUTES:
		duration := time.Minute
		d.Time = d.Time.Add(duration * time.Duration(count))
	case HOURS:
		duration := time.Hour
		d.Time = d.Time.Add(duration * time.Duration(count))
	case DAYS:
		d.Time = d.Time.AddDate(0, 0, count)
	case WEEKS:
		d.Time = d.Time.AddDate(0, 0, 7*count)
	case MONTHS:
		d.Time = d.Time.AddDate(0, count, 0)
	case YEARS:
		d.Time = d.Time.AddDate(count, 0, 0)
	}
	return &d
}

//Get the difference as a duration type
func (d *GoDate) DifferenceAsDuration(compare *GoDate) time.Duration {
	return d.Time.Sub(compare.Time)
}

//Difference Returns the difference between the Godate and another in the specified unit
//If the difference is negative then the 'compare' date occurs after the date
//Else it occurs before the the date
func (d *GoDate) Difference(compare *GoDate, unit int) int {
	difference := d.DifferenceAsFloat(compare, unit)
	return int(difference)
}

//Get the difference as a float
func (d *GoDate) DifferenceAsFloat(compare *GoDate, unit int) float64 {
	duration := d.DifferenceAsDuration(compare)
	switch unit {
	case MINUTES:
		return duration.Minutes()
	case HOURS:
		return duration.Hours()
	case DAYS:
		return float64(duration / DAY)
	case WEEKS:
		return float64(duration / WEEK)
	case MONTHS:
		return float64(duration / MONTH)
	default:
		return float64(duration.Hours() / 24)
	}
}

//Gets the difference between the relative to the date value in the form of
//1 month before
//1 month after
func (d *GoDate) DifferenceForHumans(compare *GoDate, ) string {
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
func (d *GoDate) DifferenceFromNowForHumans(unit int) string {
	now := Now()
	differenceString, differenceInt := now.AbsDifferenceForHumans(d)
	if differenceInt > 0 {
		return differenceString + " ago"
	} else {
		return differenceString + " from now"
	}
}

//Get the abs difference relative to compare time in the form
//1 month
//2 days
func (d *GoDate) AbsDifferenceForHumans(compare *GoDate) (string, int) {
	sentence := make([]string, 2, 2)
	duration := time.Duration(math.Abs(float64(d.DifferenceAsDuration(compare))))
	unit := 0
	if duration >= YEAR {
		unit = YEARS
	} else if duration < YEAR && duration >= MONTH {
		unit = MONTHS
	} else if duration < MONTH && duration >= WEEK {
		unit = WEEKS
	} else if duration < WEEK && duration >= DAY {
		unit = DAYS
	} else if duration < DAY && duration >= time.Hour {
		unit = HOURS
	} else if duration < time.Hour && duration >= time.Minute {
		unit = MINUTES
	} else {
		unit = SECONDS
	}
	difference := d.Difference(compare, unit)
	sentence[0] = strconv.Itoa(int(math.Abs(float64(difference))))
	if difference == 1 || difference == -1 {
		sentence[1] = strings.TrimSuffix(UnitStrings[unit], "s")
	} else {
		sentence[1] = UnitStrings[unit]
	}
	return strings.Join(sentence, " "), difference
}

func (date *GoDate) StartOfMinute() *GoDate{
	y, m, d := date.Time.Date()
	return &GoDate{time.Date(y, m, d, date.Time.Hour(), date.Time.Minute(), 0, 0, date.Time.Location())}
}

func (date *GoDate) StartOfHour() *GoDate{
	y, m, d := date.Time.Date()
	return &GoDate{time.Date(y, m, d, date.Time.Hour(), 0, 0, 0, date.Time.Location())}
}

func (date *GoDate) StartOfDay() *GoDate {
	y, m, d := date.Time.Date()
	return &GoDate{time.Date(y, m, d, 0, 0, 0, 0, date.Time.Location())}
}

func (date *GoDate) StartOfWeek() *GoDate{
	day := date.StartOfDay().Time.Weekday()
	if day != FirstDayOfWeek{
		return date.Sub(int(day - FirstDayOfWeek),DAYS)
	}
	return nil
}

func (date *GoDate) StartOfMonth() *GoDate{
	y, m, _ := date.Time.Date()
	return &GoDate{time.Date(y, m, 1, 0, 0, 0, 0, date.Time.Location())}
}

func (date *GoDate) StartOfQuarter() *GoDate{
	startMonth := date.StartOfMonth()
	off := (startMonth.Time.Month() - 1) % 3
	return startMonth.Sub(int(off),MONTHS)
	return nil;
}

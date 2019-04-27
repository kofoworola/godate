package godate

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type GoDate struct {
	StartOfWeek time.Weekday
	Time        time.Time
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
	milliSecondOffset := time.Millisecond * time.Duration(count/int(math.Abs(float64(count))))
	switch unit {
	case MINUTES:
		duration := time.Minute
		d.Time = d.Time.Add(duration * time.Duration(count))
	case HOURS:
		duration := time.Hour
		d.Time = d.Time.Add(duration * time.Duration(count))
	case DAYS:
		d.Time = d.Time.AddDate(0, 0, count).Add(milliSecondOffset)
	case WEEKS:
		d.Time = d.Time.AddDate(0, 0, 7*count).Add(milliSecondOffset)
	case MONTHS:
		d.Time = d.Time.AddDate(0, count, 0).Add(milliSecondOffset)
	case YEARS:
		d.Time = d.Time.AddDate(count, 0, 0).Add(milliSecondOffset)
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
//1 month ago
//1 month after
func (d *GoDate) DifferenceForHumans(compare *GoDate, unit int) string {
	sentence := make([]string, 3, 3)
	difference := d.Difference(compare, unit)
	sentence[0] = strconv.Itoa(int(math.Abs(float64(difference))))
	if difference == 1 || difference == -1 {
		sentence[1] = strings.TrimSuffix(UnitStrings[unit], "s")
	} else {
		sentence[1] = UnitStrings[unit]
	}
	if difference > 0 {
		sentence[2] = "before"
	} else {
		sentence[2] = "after"
	}
	return strings.Join(sentence, " ")
}

func (d *GoDate) DifferenceFromNowForHumans(unit int) string {
	now := Now()
	sentence := make([]string, 3, 3)
	difference := now.Difference(d, unit)
	sentence[0] = strconv.Itoa(int(math.Abs(float64(difference))))
	if difference == 1 || difference == -1 {
		sentence[1] = strings.TrimSuffix(UnitStrings[unit], "s")
	} else {
		sentence[1] = UnitStrings[unit]
	}
	if difference > 0 {
		sentence[2] = "ago"
	} else {
		sentence[2] = "from now"
	}
	return strings.Join(sentence, " ")
}

func (d *GoDate) AbsDifferenceForHumans(compare *GoDate) string {
	sentence := make([]string, 2, 2)
	duration := time.Duration(math.Abs(float64(d.DifferenceAsDuration(compare))))
	unit := 0
	if duration > YEAR {
		unit = YEARS
	} else if duration < YEAR && duration > MONTH {
		unit = MONTHS
	} else if duration < MONTH && duration > WEEK{
		unit = WEEKS
	} else if duration < WEEK && duration > DAY{
		unit = DAYS
	} else if duration < DAY && duration > time.Hour{
		unit = HOURS
	} else if duration < time.Hour && duration > time.Minute {
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
	return strings.Join(sentence, " ")
}

package godate

import (
	"time"
)

var FirstDayOfWeek = time.Monday

type Unit time.Duration

//UNIT CONSTANTS
const (
	SECOND Unit = Unit(time.Second)
	MINUTE = 60 * SECOND
	HOUR = 60 * MINUTE
	DAY   = HOUR * 24
	WEEK  = DAY * 7
	MONTH = DAY * 30
	YEAR  = DAY * 365
)

var UnitStrings = map[Unit]string{
	SECOND: "seconds",
	MINUTE: "minutes",
	HOUR:   "hours",
	DAY:    "days",
	WEEK:   "weeks",
	MONTH:  "months",
	YEAR:   "years",
}

func (u Unit) String() string{
	return UnitStrings[u]
}

func Create(time time.Time) *GoDate{
	return &GoDate{time,time.Location()}
}

func Now(location *time.Location) *GoDate {
	return &GoDate{time.Now().In(location),location}
}

func Tomorrow(location *time.Location) *GoDate {
	tomorrow := Now(location).Add(1,DAYS)
	return tomorrow
}

func Yesterday(location *time.Location) *GoDate {
	yesterday := Now(location).Sub(1,DAYS)
	return yesterday
}

package godate

import (
	"time"
)

var FirstDayOfWeek = time.Monday

//TYPE CONSTANTS
//TODO Improve this by using the String method
const (
	SECONDS = 0x00001
	MINUTES = 0x00002
	HOURS   = 0x00003
	DAYS    = 0x00004
	WEEKS   = 0x00005
	MONTHS  = 0x00006
	YEARS   = 0x00007
)

var UnitStrings = map[int]string{
	SECONDS: "seconds",
	MINUTES: "minutes",
	HOURS:   "hours",
	DAYS:    "days",
	WEEKS:   "weeks",
	MONTHS:  "months",
	YEARS:   "years",
}

//UNIT Value CONSTANTS
const (
	DAY   = time.Hour * 24
	WEEK  = DAY * 7
	MONTH = DAY * 30
	YEAR  = DAY * 365
)

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

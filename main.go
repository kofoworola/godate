package godate

import "time"

var StartOfWeek = time.Monday


//TYPE CONSTANTS
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
	HOURS: "hours",
	DAYS: "days",
	WEEKS: "weeks",
	MONTHS: "months",
	YEARS: "years",
}

//UNIT Value CONSTANTS
const(
	DAY = time.Hour * 24
	WEEK = DAY * 7
	MONTH = DAY * 30
	YEAR = DAY * 365
)

func Now() *GoDate {
	return &GoDate{StartOfWeek, time.Now()}
}

func FromString(stringToParse string) *GoDate {
	switch stringToParse {
	case "tomorrow":
		return Tomorrow()
	case "yesterday":
		return Yesterday()
	default:
		return Now()
	}
}

func Tomorrow() *GoDate {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return &GoDate{StartOfWeek, tomorrow}
}

func Yesterday() *GoDate {
	yesterday := time.Now().AddDate(0, 0, -1)
	return &GoDate{StartOfWeek, yesterday}
}
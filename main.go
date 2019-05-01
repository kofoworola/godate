package godate

import (
	"time"
)

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

func Create(time time.Time) *goDate {
	return &goDate{time,time.Location(),0}
}

func Now(location *time.Location) *goDate {
	return &goDate{time.Now().In(location),location,0}
}

func Tomorrow(location *time.Location) *goDate {
	tomorrow := Now(location).Add(1,DAY)
	return tomorrow
}

func Yesterday(location *time.Location) *goDate {
	yesterday := Now(location).Sub(1,DAY)
	return yesterday
}

func Parse(layout, value string) (*goDate,error){
	parsedTime, err := time.Parse(layout,value)
	if err != nil{
		return nil, err
	}
	return &goDate{Time: parsedTime, TimeZone: parsedTime.Location()},nil
}

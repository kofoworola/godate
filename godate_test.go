package godate

import (
	"strconv"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	now := parse(Now().Time)
	if now != parse(time.Now()){
		t.Error("Expected " + parse(time.Now()) + " got "+ now)
	}
}

func TestGoDate_Difference(t *testing.T) {
	today := Now()
	yesterday := today.Sub(1, HOURS)
	if difference := today.Difference(yesterday, HOURS); difference != 1{
		t.Error("Expected 1 got " + strconv.Itoa(difference))
	}
	lastWeek := today.Sub(2, WEEKS)
	if difference := today.Difference(lastWeek, MONTHS); difference != 0{
		t.Error("Expected 0 got " + strconv.Itoa(difference))
	}
}

func parse(time time.Time) string{
	return time.Format("2006-01-02")
}


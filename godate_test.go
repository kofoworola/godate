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
	yesterday := today.Sub(1, DAYS)
	tomorrow := today.Add(1,DAYS)
	if difference := today.Difference(yesterday, DAYS); difference != 1{
		t.Error("Expected 1 got " + strconv.Itoa(difference))
	}
	if difference := today.Difference(tomorrow, DAYS); difference != -1{
		t.Error("Expected -1 got " + strconv.Itoa(difference))
	}
	lastWeek := today.Sub(2, WEEKS)
	if difference := today.Difference(lastWeek, MONTHS); difference != 0{
		t.Error("Expected 0 got " + strconv.Itoa(difference))
	}
}

func TestGoDate_DifferenceForHumans(t *testing.T) {
	today := Now()
	yesterday := today.Sub(1, DAYS)
	tomorrow := today.Add(2,DAYS)
	if difference := today.DifferenceForHumans(yesterday); difference != "1 day before"{
		t.Error("got " + difference)
	}
	if difference := today.DifferenceForHumans(tomorrow); difference != "2 days after"{
		t.Error("got " + difference)
	}
}

func TestGoDate_DifferenceFromNowForHumans(t *testing.T) {
	yesterday := Yesterday()
	if difference := yesterday.DifferenceFromNowForHumans(DAYS); difference != "1 day ago"{
		t.Error("got " + difference)
	}
	now := Now()
	nextWeek := now.Add(1,WEEKS)
	if difference := nextWeek.DifferenceFromNowForHumans(WEEKS); difference != "1 week from now"{
		t.Error("got " + difference)
	}
}

func TestGoDate_AbsDifferenceForHumans(t *testing.T) {
	today := Now()
	yesterday := today.Sub(1, DAYS)
	tomorrow := today.Add(7,DAYS)
	if difference,_ := today.AbsDifferenceForHumans(yesterday); difference != "1 day"{
		t.Error("got " + difference)
	}
	if difference,_ := today.AbsDifferenceForHumans(tomorrow); difference != "1 week"{
		t.Error("got " + difference)
	}
}

func parse(time time.Time) string{
	return time.Format("2006-01-02 15:04:05")
}

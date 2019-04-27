package godate

import (
	"strconv"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	now := parse(Now(time.UTC).Time)
	if now != parse(time.Now().In(time.UTC)){
		t.Error("Expected " + parse(time.Now().In(time.UTC)) + " got "+ now)
	}
}

func TestGoDate_Difference(t *testing.T) {
	today := Now(time.UTC)
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
	today := Now(time.UTC)
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
	yesterday := Yesterday(time.UTC)
	if difference := yesterday.DifferenceFromNowForHumans(DAYS,time.UTC); difference != "1 day ago"{
		t.Error("got " + difference)
	}
	now := Now(time.UTC)
	//Add minute for offsetting due to time lag
	nextWeek := now.Add(1,WEEKS).Add(1,MINUTES)
	if difference := nextWeek.DifferenceFromNowForHumans(WEEKS,time.UTC); difference != "1 week from now"{
		t.Error("got " + difference)
	}
}

func TestGoDate_AbsDifferenceForHumans(t *testing.T) {
	today := Now(time.UTC)
	yesterday := today.Sub(1, DAYS)
	tomorrow := today.Add(7,DAYS)
	if difference,_ := today.AbsDifferenceForHumans(yesterday); difference != "1 day"{
		t.Error("got " + difference)
	}
	if difference,_ := today.AbsDifferenceForHumans(tomorrow); difference != "1 week"{
		t.Error("got " + difference)
	}
}

func TestGoDate_StartOfDay(t *testing.T) {
	today := Now(time.UTC).StartOfDay()
	if today.Time.Hour() != 0 || today.Time.Second() != 0{
		t.Error("Got "+ parse(today.Time))
	}
}

func TestGoDate_StartOfWeek(t *testing.T) {
	date := GoDate{time.Date(2019,4,27,0,0,0,0,time.UTC),time.UTC}
	FirstDayOfWeek = time.Sunday
	if date.StartOfWeek().Time.Day() != 21{
		t.Error("Got " + parse(date.StartOfWeek().Time))
	}
}

func parse(time time.Time) string{
	return time.Format("2006-01-02 15:04:05")
}

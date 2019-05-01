package godate

import "time"

//Provide useful modifications to the current godate instance

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

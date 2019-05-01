package godate

import "time"

func (d *goDate) Year(year int) *goDate{
	old := d.Time
	d.Time = time.Date(year,old.Month(),old.Day(),old.Hour(),old.Minute(),old.Second(),old.Nanosecond(),d.TimeZone)
	return d
}

func (d *goDate) Moth(month int){
	old := d.Time
	d.Time = time.Date(old.Year(), time.Month(month),old.Day(),old.Hour(),old.Minute(),old.Second(),old.Nanosecond(),d.TimeZone)
}

func (d *goDate) Day(day int){
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),day,old.Hour(),old.Minute(),old.Second(),old.Nanosecond(),d.TimeZone)
}

func (d *goDate) Hour(hour int){
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),old.Day(),hour,old.Minute(),old.Second(),old.Nanosecond(),d.TimeZone)
}

func (d *goDate) Minute(minute int){
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),old.Day(),old.Hour(),minute,old.Second(),old.Nanosecond(),d.TimeZone)
}

func (d *goDate) Second(second int){
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),old.Day(),old.Hour(),old.Minute(),second,old.Nanosecond(),d.TimeZone)
}

func (d *goDate) Nanosecond(nanosecond int){
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),old.Day(),old.Hour(),old.Minute(),old.Second(),nanosecond,d.TimeZone)
}


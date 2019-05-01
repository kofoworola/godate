package godate

import "time"

//Year sets the year of the godate variable
func (d *goDate) Year(year int) *goDate{
	old := d.Time
	d.Time = time.Date(year,old.Month(),old.Day(),old.Hour(),old.Minute(),old.Second(),old.Nanosecond(),d.TimeZone)
	return d
}

//Month sets the month of the year of the godate variable
func (d *goDate) Month(month int) *goDate{
	old := d.Time
	d.Time = time.Date(old.Year(), time.Month(month),old.Day(),old.Hour(),old.Minute(),old.Second(),old.Nanosecond(),d.TimeZone)
	return d
}

//Day sets the day of the week of the godate variable
func (d *goDate) Day(day int) *goDate{
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),day,old.Hour(),old.Minute(),old.Second(),old.Nanosecond(),d.TimeZone)
	return d
}

//Hour sets the hour of the godate variable. Range of 0-59
func (d *goDate) Hour(hour int) *goDate{
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),old.Day(),hour,old.Minute(),old.Second(),old.Nanosecond(),d.TimeZone)
	return d
}

//Minute sets the minute of the godate variable
func (d *goDate) Minute(minute int) *goDate{
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),old.Day(),old.Hour(),minute,old.Second(),old.Nanosecond(),d.TimeZone)
	return d
}

//Second sets the second of the godate variable
func (d *goDate) Second(second int) *goDate{
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),old.Day(),old.Hour(),old.Minute(),second,old.Nanosecond(),d.TimeZone)
	return d
}

//Second sets the current nanosecond of the godate variable
func (d *goDate) Nanosecond(nanosecond int) *goDate{
	old := d.Time
	d.Time = time.Date(old.Year(), old.Month(),old.Day(),old.Hour(),old.Minute(),old.Second(),nanosecond,d.TimeZone)
	return d
}


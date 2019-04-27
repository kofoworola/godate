package godate

import "time"

type GoDate struct{
	StartOfWeek time.Weekday
	Time time.Time
}

func (d GoDate) IsBefore(compare *GoDate) bool{
	return d.Time.Before(compare.Time)
}

func (d GoDate) IsAfter(compare *GoDate) bool{
	return d.Time.After(compare.Time)
}

func (d GoDate) Sub(count int, unit int) *GoDate{
	switch unit {
	case MINUTES:
		duration := time.Minute
		d.Time = d.Time.Add(-duration * time.Duration(count))
	case HOURS:
		duration := time.Hour
		d.Time = d.Time.Add(-duration * time.Duration(count))
	case DAYS:
		d.Time = d.Time.AddDate(0,0,-count);
	case WEEKS:
		d.Time = d.Time.AddDate(0,0,-7 * count)
	case MONTHS:
		d.Time = d.Time.AddDate(0,-count,0)
	case YEARS:
		d.Time = d.Time.AddDate(-count,0,0)
	}
	return &d
}

func (d *GoDate) Difference(compare *GoDate,unit int) int{
	duration := d.Time.Sub(compare.Time)
	switch unit {
	case MINUTES:
		return int(duration.Minutes())
	case HOURS:
		return int(duration.Hours())
	case DAYS:
		return int(duration/DAY)
	case WEEKS:
		return int(duration/WEEK)
	case MONTHS:
		return int(duration/MONTH)
	default:
		return int(duration.Hours()/24)
	}
}



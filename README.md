[![Build Status](https://travis-ci.org/kofoworola/godate.svg?branch=master)](https://travis-ci.org/kofoworola/godate)

# GODATE
Godate is a date toolkit for golang for easy date manupulation highly inspired by php's carbon [Carbon](https://carbon.nesbot.com/)
# Installation
```
go get github.com/kofoworola/godate
```
# Usage
The `GoDate` struct is simply a wrapper around a `time.TIme` struct and a `time.Location` object
most of the methods available to the  `GoDate` return a corresponding `GoDate` (Except the Difference and boolean based methods) that
allows you to chain methods together 
## Initialization
```
godate.Create(time.Now())
godate.Now(time.UTC)
godate.Tomorrow(time.UTC)
```
outputs
```
Sat Apr 27 18:53:11 +0100 WAT 2019
Sat Apr 27 17:53:11 +0000 UTC 2019
Sun Apr 28 17:53:11 +0000 UTC 2019
```

## Methods
```
now := godate.Now(time.UTC)
now.IsAfter(now.Add(1,godate.DAYS)) //false
now.IsAfter(now.Sub(1,godate.DAYS)) //true
now.IsBefore(now.Add(1,godate.DAYS)) //true
now.IsBefore(now.Sub(1,godate.DAYS)) //false
```

The `Add()` and `Sub()` methods takes two parameters, the `count` and the
`unit` to add which is any of the constants
```
SECONDS 
MINUTES
HOURS
DAYS 
WEEKS
MONTHS
YEARS
```


### Difference
```
now := godate.Now(time.UTC)
now.Difference(now.Sub(1,godate.DAYS),godate.DAYS) //1 
now.Difference(now.Add(7,godate.DAYS),godate.WEEKS) //-1
```

The `Difference()` method returns the difference of the passed `godate`
relative to the `godate` that the method is being called on. It returns the difference
in the specified `unit` (the second parameter).
###### +ve means the passed `godate` occurs after the method owner while -ve is the opposite

### Difference for humans
```
fmt.Println(now.DifferenceForHumans(now.Add(1,godate.WEEKS))) //1 week after
fmt.Println(now.Add(8,godate.DAYS).DifferenceFromNowForHumans()) //1 week from 
fmt.Println(now.Sub(21,godate.DAYS).DifferenceFromNowForHumans()) //3 weeks ago
```
The `DifferenceForHumans` methods works similar to the corresponding
[carbon method](https://carbon.nesbot.com/docs/#api-humandiff)
It compares of the passed `GoDate` relative to the one that the method is being called on

The `DifferenceFromNowForHumans` methods gets the difference of the date relative to the current 
time in the same timezone

### Helper methods
Below are the available self explanatory helper methods
```
fmt.Println(now) //Sun Apr 28 07:14:01 +0000 UTC 2019
fmt.Println(now.StartOfHour()) //Sun Apr 28 07:00:00 +0000 UTC 2019
fmt.Println(now.StartOfDay()) //Sun Apr 28 00:00:00 +0000 UTC 2019
fmt.Println(now.StartOfQuarter()) //Mon Apr 1 00:00:00 +0000 UTC 2019
fmt.Println(now.StartOfMonth()) //Mon Apr 1 00:00:00 +0000 UTC 2019
fmt.Println(now.StartOfYear()) //Tue Jan 1 00:00:00 +0000 UTC 2019

fmt.Println(now.EndOfHour()) //Sun Apr 28 07:59:59 +0000 UTC 2019
fmt.Println(now.EndOfDay()) //Sun Apr 28 23:59:59 +0000 UTC 2019
fmt.Println(now.EndOfQuarter()) //Sun Jun 30 23:59:59 +0000 UTC 2019
fmt.Println(now.EndOfMonth()) //Tue Apr 30 23:59:59 +0000 UTC 2019
fmt.Println(now.EndOfYear()) //Thu Jan 31 23:59:59 +0000 UTC 2019
fmt.Println(now.IsWeekend()) //true
```
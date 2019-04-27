[![Build Status](https://travis-ci.org/kofoworola/godate.svg?branch=master)](https://travis-ci.org/kofoworola/godate)

# GODATE
Godate is a date toolkit for golang for easy date manupulation highly inspired by php's carbon [Carbon](https://carbon.nesbot.com/)
# Installation
```
go get github.com/kofoworola/godate
```
# Usage
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

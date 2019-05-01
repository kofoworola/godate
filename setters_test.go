package godate

import (
	"fmt"
	"testing"
)

func TestGoDate_Setter(t *testing.T){
	date,_ := Parse("2006-01-02","2019-05-01")
	date.Year(2008).Month(10).Day(30).Hour(11).Minute(	3).Second(12)
	if formatted := date.Format("2006-01-02 15:04:05"); formatted != "2008-10-30 11:03:12"{
		fmt.Println("Got "+formatted)
	}
}
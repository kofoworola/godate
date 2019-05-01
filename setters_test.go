package godate

import (
	"fmt"
	"testing"
)

func TestGoDate_Setter(t *testing.T){
	date,_ := Parse("2006-01-02","2019-05-01")
	date.Year(2008).Year(2019)
	fmt.Println(date)
}
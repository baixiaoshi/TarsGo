package endpoint

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	e := Parse("tcp -h 10.219.139.142 -p 19386 -t 60000")
	fmt.Println(e)
	e2 := Parse("udp -h 10.219.139.142 -p 19386 -t 60000")
	fmt.Println(e2)
	tars := Endpoint2tars(e2)
	fmt.Println(tars)
	fmt.Println(Tars2endpoint(tars))
}

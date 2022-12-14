package director

import (
	"fmt"
	"strings"
	"time"
)

const layout = "2006-01-02"

type Director struct {
	Id          int        `json:"-"`
	Name        string     `json:"name"`
	DateOfBirth CustomDate `json:"date_of_birth"`
	Bio         string     `json:"bio"`
}

type CustomDate struct {
	time.Time
}

func (cus *CustomDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return
	}
	cus.Time, err = time.Parse(layout, s)
	return
}

func (cus *CustomDate) Scan(src interface{}) (err error) {
	value, ok := src.(string)
	if ok {
		cus.Time, err = time.Parse(layout, value)
	} else {
		err = fmt.Errorf("Unexpected type: %T", src)
	}
	return
}

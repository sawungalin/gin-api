package app

import (
	"encoding/json"
	"time"
)

type MyTime time.Time

var _ json.Unmarshaler = &MyTime{}

func (mt *MyTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.Local)
	if err != nil {
		return err
	}
	*mt = MyTime(t)
	return nil
}

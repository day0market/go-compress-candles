package main

import (
	"reflect"
	"time"
)

type CandleSimple struct {
	DateTime time.Time
	Volume   int64
	Open     float64
	High     float64
	Low      float64
	Close    float64
}

type CandleEnriched struct {
	DateTime  time.Time
	OpenTime  time.Time
	CloseTime time.Time

	DayOpen  float64
	DayClose float64
	DayHigh  float64
	DayLow   float64

	DayLowWithoutFirstMin  float64
	DayHighWithoutFirstMin float64

	DayVolume int64

	DayHighTime time.Time
	DayLowTime  time.Time

	DayHighTimeWithoutFirstMin time.Time
	DayLowTimeWithoutFirstMin  time.Time

	High9  float64
	High10 float64
	High11 float64
	High12 float64
	High13 float64
	High14 float64
	High15 float64

	Low9  float64
	Low10 float64
	Low11 float64
	Low12 float64
	Low13 float64
	Low14 float64
	Low15 float64
}

func (c *CandleEnriched) setPriceValue(field string, value float64) {
	v := reflect.ValueOf(c).Elem().FieldByName(field)
	if v.IsValid() {
		v.SetFloat(value)
	}
}

func (c *CandleEnriched) getPriceValue(field string) float64 {
	v := reflect.ValueOf(c).Elem().FieldByName(field).Float()
	return v
}

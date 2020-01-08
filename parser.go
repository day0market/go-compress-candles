package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func parseRowToCandle(row string) (candle *CandleSimple, err error) {
	// 20180815093000,11.840000,11.840000,11.840000,11.840000,2800
	s := strings.Split(row, ",")
	if len(s) != 6 {
		err = errors.New("Failed to parse row to candle. Expected len == 6, Found line with len = " + strconv.Itoa(len(s)))
		return candle, err
	}

	datetime, err := time.Parse("20060102150405", s[0])
	if err != nil {
		return candle, err
	}

	o, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		return candle, err
	}
	h, err := strconv.ParseFloat(s[2], 64)
	if err != nil {
		return candle, err
	}
	l, err := strconv.ParseFloat(s[3], 64)
	if err != nil {
		return candle, err
	}
	c, err := strconv.ParseFloat(s[4], 64)
	if err != nil {
		return candle, err
	}
	v, err := strconv.ParseInt(s[5], 10, 64)
	if err != nil {
		return candle, err
	}

	candle = &CandleSimple{
		DateTime: datetime,
		Volume:   v,
		Open:     o,
		High:     h,
		Low:      l,
		Close:    c,
	}
	return candle, err
}

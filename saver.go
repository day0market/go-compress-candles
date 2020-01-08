package main

import (
	"fmt"
	"os"
)

const headerRow = "DateTime,OpenTime,CloseTime,DayOpen,DayClose,DayHigh,DayLow,DayLowWithoutFirstMin," +
	"DayHighWithoutFirstMin,DayVolume,DayHighTime,DayLowTime,DayHighTimeWithoutFirstMin,DayLowTimeWithoutFirstMin," +
	"High9,High10,High11,High12,High13,High14,High15,Low9,Low10,Low11,Low12,Low13,Low14,Low15"

const dateTimeFormat = "2006-01-02 15:04:05"

func saveCandlesEnriched(compressed []*CandleEnriched, path string) (err error) {
	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s DONE!\n", path)
	}()

	writeRow := func(str string) {
		_, err = fmt.Fprintln(f, str)
		if err != nil {
			panic(err)
		}
	}

	writeRow(headerRow)

	for _, v := range compressed {
		writeRow(candleEnrichedToString(v))
	}

	return err
}

func candleEnrichedToString(c *CandleEnriched) string {
	//
	date := c.DateTime.Format("2006-01-02")
	openTime := c.OpenTime.Format(dateTimeFormat)
	closeTime := c.CloseTime.Format(dateTimeFormat)
	DayHighTime := c.DayHighTime.Format(dateTimeFormat)
	DayLowTime := c.DayLowTime.Format(dateTimeFormat)

	DayHighTimeWithoutFirstMin := c.DayHighTimeWithoutFirstMin.Format(dateTimeFormat)
	DayLowTimeWithoutFirstMin := c.DayLowTimeWithoutFirstMin.Format(dateTimeFormat)

	return fmt.Sprintf(
		"%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v%v,%v,%v,%v%v,%v,%v,%v,%v",
		date, openTime, closeTime, c.DayOpen, c.DayClose, c.DayHigh, c.DayLow,
		c.DayLowWithoutFirstMin, c.DayHighWithoutFirstMin,
		c.DayVolume, DayHighTime, DayLowTime, DayHighTimeWithoutFirstMin, DayLowTimeWithoutFirstMin,
		c.High9, c.High10, c.High11, c.High12, c.High13, c.High14, c.High15,
		c.Low9, c.Low10, c.Low11, c.Low12, c.Low13, c.Low14, c.Low15)
}

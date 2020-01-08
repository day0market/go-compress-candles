package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"
)

func processFile(path string) (compressed []*CandleEnriched, err error) {
	f, err := os.Open(path)

	if err != nil {
		return compressed, err
	}

	defer f.Close()

	compressed = compressFromFile(f)
	return compressed, err
}

func compressFromFile(f io.Reader) []*CandleEnriched {
	scanner := bufio.NewScanner(f)

	var compressed []*CandleEnriched
	var currentCandle *CandleEnriched
	var lastDate time.Time

	for scanner.Scan() {

		candle, err := parseRowToCandle(scanner.Text())

		if err != nil {
			log.Println(err)
			continue
		}

		if candle.DateTime.Weekday() != lastDate.Weekday() || lastDate.Year() != candle.DateTime.Year() {
			lastDate = candle.DateTime

			if currentCandle != nil {
				finalizeCandleEnrichedValues(currentCandle)
				compressed = append(compressed, currentCandle)
			}

			currentCandle = createNewCandleEnriched(candle)
			continue
		}

		updateCandleEnrichedPrices(candle, currentCandle)
		lastDate = candle.DateTime
	}

	if currentCandle != nil {
		finalizeCandleEnrichedValues(currentCandle)
		compressed = append(compressed, currentCandle)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return compressed
}

func updateCandleEnrichedPrices(cs *CandleSimple, ce *CandleEnriched) {
	if cs.High > ce.DayHigh {
		ce.DayHigh = cs.High
		ce.DayHighTime = cs.DateTime
	}

	if cs.High > ce.DayHighWithoutFirstMin {
		ce.DayHighWithoutFirstMin = cs.High
		ce.DayHighTimeWithoutFirstMin = cs.DateTime
	}

	if cs.Low < ce.DayLow {
		ce.DayLow = cs.Low
		ce.DayLowTime = cs.DateTime
	}

	if cs.Low < ce.DayLowWithoutFirstMin {
		ce.DayLowWithoutFirstMin = cs.Low
		ce.DayLowTimeWithoutFirstMin = cs.DateTime
	}

	ce.CloseTime = cs.DateTime
	ce.DayClose = cs.Close
	ce.DayVolume += cs.Volume

	updateCandleEnrichedHourPrices(cs, ce)

}

func updateCandleEnrichedHourPrices(candle *CandleSimple, lastCandle *CandleEnriched) {
	hour := candle.DateTime.Hour()
	if hour > 15 {
		return
	}

	highField := fmt.Sprintf("High%v", hour)
	lowField := fmt.Sprintf("Low%v", hour)

	curHigh := lastCandle.getPriceValue(highField)
	if candle.High > curHigh {
		lastCandle.setPriceValue(highField, candle.High)
	}

	curLow := lastCandle.getPriceValue(lowField)
	if candle.Low < curLow {
		lastCandle.setPriceValue(lowField, candle.Low)
	}

}

func createNewCandleEnriched(candleSimple *CandleSimple) *CandleEnriched {
	t := candleSimple.DateTime
	newCandle := CandleEnriched{
		DayOpen:                    candleSimple.Open,
		DayHighTime:                candleSimple.DateTime,
		DayLowTime:                 candleSimple.DateTime,
		CloseTime:                  candleSimple.DateTime,
		OpenTime:                   candleSimple.DateTime,
		DateTime:                   time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC),
		DayHigh:                    candleSimple.High,
		DayLow:                     candleSimple.Low,
		DayClose:                   candleSimple.Close,
		DayVolume:                  candleSimple.Volume,
		DayHighWithoutFirstMin:     math.Inf(-1),
		DayLowWithoutFirstMin:      math.Inf(1),
		DayHighTimeWithoutFirstMin: candleSimple.DateTime,
		DayLowTimeWithoutFirstMin:  candleSimple.DateTime,
	}

	setInfForPriceFields(&newCandle, "High", -1)
	setInfForPriceFields(&newCandle, "Low", 1)

	updateCandleEnrichedHourPrices(candleSimple, &newCandle)

	return &newCandle

}

func setInfForPriceFields(candleEnriched *CandleEnriched, baseFieldName string, infSign int) {
	for i := 9; i < 16; i++ {
		fName := fmt.Sprintf("%s%v", baseFieldName, i)
		candleEnriched.setPriceValue(fName, math.Inf(infSign))
	}
}

func finalizeCandleEnrichedValues(candleEnriched *CandleEnriched) {
	if candleEnriched == nil {
		return
	}

	_updField := func(fName string) {
		curVal := candleEnriched.getPriceValue(fName)

		if !math.IsInf(curVal, 1) && !math.IsInf(curVal, -1) {
			return
		}
		candleEnriched.setPriceValue(fName, math.NaN())

	}

	for _, baseFieldName := range []string{"High", "Low"} {
		for i := 9; i < 16; i++ {
			fName := fmt.Sprintf("%s%v", baseFieldName, i)
			_updField(fName)
		}

	}

	for _, fName := range []string{"DayLowWithoutFirstMin", "DayHighWithoutFirstMin"} {
		_updField(fName)
	}

}

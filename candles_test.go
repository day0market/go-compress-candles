package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math"
	"path"
	"testing"
)

func checkSingleFile(t *testing.T, pth string) {
	compressed, err := processFile(pth)

	assert.Nil(t, err)
	assert.True(t, len(compressed) > 0)

	for _, c := range compressed {

		assert.True(t, c.DayHigh >= c.DayOpen)
		assert.True(t, c.DayHigh >= c.DayClose)
		assert.True(t, c.DayLow <= c.DayOpen)
		assert.True(t, c.DayLow <= c.DayClose)

		assert.True(t, c.OpenTime.After(c.DateTime))

		assert.False(t, c.CloseTime.Before(c.OpenTime))
		assert.False(t, c.DayHighTime.Before(c.OpenTime))
		assert.False(t, c.DayHighTimeWithoutFirstMin.Before(c.OpenTime))
		assert.False(t, c.DayLowTimeWithoutFirstMin.Before(c.OpenTime))

		for _, fName := range []string{"DayLowWithoutFirstMin", "DayHighWithoutFirstMin"} {
			val := c.getPriceValue(fName)
			assert.False(t, math.IsInf(val, 1))
			assert.False(t, math.IsInf(val, -1))
		}

		assert.True(t, c.DayVolume > 0)

		for _, fNameB := range []string{"High", "Low"} {
			for i := 9; i < 16; i++ {
				fName := fmt.Sprintf("%s%v", fNameB, i)
				val := c.getPriceValue(fName)
				assert.False(t, math.IsInf(val, 1))
				assert.False(t, math.IsInf(val, -1))
			}
		}

		for i := 9; i < 16; i++ {
			highF := fmt.Sprintf("High%v", i)
			lowF := fmt.Sprintf("Low%v", i)

			valH := c.getPriceValue(highF)
			valL := c.getPriceValue(lowF)

			if !math.IsNaN(valH) {
				assert.False(t, math.IsNaN(valL))
				assert.True(t, valH >= valL)
			}

			if !math.IsNaN(valL) {
				assert.False(t, math.IsNaN(valH))
				assert.True(t, valH >= valL)
			}

		}

	}

}

func TestCandlesEnriched(t *testing.T) {

	files, err := ioutil.ReadDir("./test_files")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		checkSingleFile(t, path.Join("./test_files", f.Name()))
	}

}

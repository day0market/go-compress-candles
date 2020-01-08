package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

const SOURCE_FOLDER = "D:\\MarketData\\US\\Candles\\1Min"
const DEST_FOLDER = "D:\\MarketData\\US\\EnrichedDays"
const FILE_MASK = ".txt"

func main() {
	files, err := ioutil.ReadDir(SOURCE_FOLDER)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), FILE_MASK) {
			continue
		}

		compressed, err := processFile(path.Join(SOURCE_FOLDER, f.Name()))

		if err != nil {
			log.Printf("Failed to process file: %s", f.Name())
			continue
		}

		if compressed == nil || len(compressed) == 0 {
			fmt.Printf("%s no candles found", f.Name())
			continue
		}

		err = saveCandlesEnriched(compressed, path.Join(DEST_FOLDER, f.Name()))
		if err != nil {
			log.Printf("Failed to save processed file: %s", f.Name())
		}
	}
}

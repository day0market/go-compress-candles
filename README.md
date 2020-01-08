# GO tools to compress 1 min OHLC to Daily OHLC

You can use it to compress 1 min OHLC to **enriched** daily bars (With high time, low time, highs and lows by hours, 
high and low with excluded first candle). Works only for US stocks. 

## Input (replace constants in main.go):

1. `SOURCE_FOLDER`: path to folder with 1 minutes bars. It should have 1 file for each stock. 
File format: 'Datetime (YYYYmmddHHMMSS)', 'Open', 'High', 'Low', 'Close', 'Volume' - comma separated, one ohlc per line.
Check example on `test_files` folder
2. `DEST_FOLDER`: path where you want store created daily quotes
3. `FILE_MASK` - use only this files extension from 

## How to run

Compile as folder and run :)
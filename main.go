package main

import (
	"mathias/binance-ctc-converter/binance"
	"mathias/binance-ctc-converter/ctc"
)

const (
	binanceFilePath       = "assets/binance.csv"
	binanceOutputFilePath = "assets/binance_output.csv"
)

func main() {
	ctcTransactions := binance.ConvertBinanceCSV(binanceFilePath, binanceOutputFilePath)
	ctc.ConvertTransactions(binanceOutputFilePath, ctcTransactions)
}

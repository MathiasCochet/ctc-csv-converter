package binance

import (
	"fmt"
	"mathias/binance-ctc-converter/ctc"
)

func ConvertBinanceCSV(inputFilePath, outputFilePath string) []ctc.Transaction {
	binanceTransactions, err := parseCSV(inputFilePath)
	if err != nil {
		fmt.Println(err)
	}

	return mapBinanceTransactionsToCTCTransactions(binanceTransactions)
}

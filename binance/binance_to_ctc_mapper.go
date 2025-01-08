package binance

import (
	"fmt"
	"mathias/binance-ctc-converter/ctc"
	"regexp"
	"strconv"
)

func mapBinanceTransactionsToCTCTransactions(binanceTransactions []Transaction) []ctc.Transaction {
	var ctcTransactions = []ctc.Transaction{}
	for i, transaction := range binanceTransactions {
		ctcTransactions = append(ctcTransactions, ctc.Transaction{
			TimestampUTC:  transaction.Date.String(),
			Type:          mapTransactionType(transaction.Side),
			BaseCurrency:  extractCryptoSymbol(transaction.Executed),
			BaseAmount:    extractAmount(transaction.Executed),
			QuoteCurrency: extractCryptoSymbol(transaction.Amount),
			QuoteAmount:   extractAmount(transaction.Amount),
			FeeCurrency:   extractCryptoSymbol(transaction.Fee),
			FeeAmount:     extractAmount(transaction.Fee),
			ID:            fmt.Sprintf("%d", i),
			Description:   fmt.Sprintf("%s %s for %s", transaction.Side, transaction.Executed, transaction.Amount),
		})
	}
	return ctcTransactions
}

func mapTransactionType(transactionType string) string {
	switch transactionType {
	case "BUY":
		return "Buy"
	case "SELL":
		return "Sell"
	}
	panic("unkown TYPE detected....")
}

func extractCryptoSymbol(value string) string {
	// Define a regex to match only letters (A-Z, a-z)
	re := regexp.MustCompile(`[A-Za-z]+$`)
	match := re.FindString(value)
	return match
}

func extractAmount(value string) float64 {
	re := regexp.MustCompile(`^[0-9.]+`)
	amountString := re.FindString(value)
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		fmt.Printf("Error converting number: %s", amountString)
	}
	return amount
}

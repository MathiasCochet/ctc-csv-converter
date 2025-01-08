package ctc

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ConvertTransactions(filePath string, transactions []Transaction) {
	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"Timestamp (UTC)", "Type", "Base Currency", "Base Amount",
		"Quote Currency (Optional)", "Quote Amount (Optional)",
		"Fee Currency (Optional)", "Fee Amount (Optional)",
		"From (Optional)", "To (Optional)", "Blockchain (Optional)",
		"ID (Optional)", "Description (Optional)",
		"Reference Price Per Unit (Optional)",
		"Reference Price Currency (Optional)",
	}
	writer.Write(header)

	for _, t := range transactions {
		row := []string{
			t.TimestampUTC,
			t.Type,
			t.BaseCurrency,
			formatFloat(t.BaseAmount),
			t.QuoteCurrency,
			formatFloat(t.QuoteAmount),
			t.FeeCurrency,
			formatFloat(t.FeeAmount),
			t.From,
			t.To,
			t.Blockchain,
			t.ID,
			t.Description,
			formatFloat(t.ReferencePricePerUnit),
			t.ReferencePriceCurrency,
		}
		writer.Write(row)
	}

	fmt.Println("CSV written successfully")
}

func formatFloat(f float64) string {
	if f == 0 {
		return ""
	}
	return strconv.FormatFloat(f, 'f', -1, 64)

}

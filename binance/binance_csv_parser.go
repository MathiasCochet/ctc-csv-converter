package binance

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func parseCSV(filePath string) ([]Transaction, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.Comma = ','

	// Read all rows
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	// Check if the CSV has a header and data
	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file is empty or has no data rows")
	}

	// Parse the header to ensure correctness (optional)
	header := records[0]
	expectedHeader := []string{"Date(UTC)", "Pair", "Side", "Price", "Executed", "Amount", "Fee"}
	if len(header) != len(expectedHeader) {
		return nil, fmt.Errorf("unexpected header format")
	}

	// Parse rows into Transactions
	var transactions []Transaction
	for i, row := range records[1:] { // Skip header row
		if len(row) != len(expectedHeader) {
			return nil, fmt.Errorf("unexpected number of columns in row %d", i+2)
		}

		// Parse Date
		date, err := time.Parse("2006-01-02 15:04:05", row[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing date in row %d: %v", i+2, err)
		}

		// Parse Price
		price, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing price in row %d: %v", i+2, err)
		}

		// Create Transaction
		transaction := Transaction{
			Date:     date,
			Pair:     row[1],
			Side:     row[2],
			Price:    price,
			Executed: row[4],
			Amount:   row[5],
			Fee:      row[6],
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

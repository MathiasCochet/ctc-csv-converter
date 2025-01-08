package ctc

type Transaction struct {
	TimestampUTC           string  // Timestamp of the transaction (UTC)
	Type                   string  // Type of transaction
	BaseCurrency           string  // Base currency involved in the transaction
	BaseAmount             float64 // Amount of the base currency
	QuoteCurrency          string  // Quote currency (optional)
	QuoteAmount            float64 // Amount of the quote currency (optional)
	FeeCurrency            string  // Fee currency (optional)
	FeeAmount              float64 // Fee amount (optional)
	From                   string  // Source of the transaction (optional)
	To                     string  // Destination of the transaction (optional)
	Blockchain             string  // Blockchain used (optional)
	ID                     string  // Transaction ID (optional)
	Description            string  // Description of the transaction (optional)
	ReferencePricePerUnit  float64 // Reference price per unit (optional)
	ReferencePriceCurrency string  // Currency of the reference price (optional)
}

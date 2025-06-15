package models

type Master struct {
	ID              uint    `gorm:"primaryKey"`
	Segment         string  `json:"segment"`
	Token           string  `json:"token"`
	Symbol          string  `json:"symbol"`
	TradingSymbol   string  `json:"tradingSymbol"`
	IntrtumentType  string  `json:"instrumentType"`
	ExpiryDate      int64   `json:"expiryDate"`
	TickSize        int     `json:"tickSize"`
	LotSize         int     `json:"lotSize"`
	OptionType      string  `json:"optionType"`
	Strike          int     `json:"strike"`
	PricePrec       int     `json:"pricePrecision"`
	Multiplier      int     `json:"multiplier"`
	Isin            *string `json:"isin"`
	PriceMultiplier float64 `json:"priceMultiplier"`
	CompanyName     string  `json:"companyName"`
}

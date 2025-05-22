package models

type ChartData struct {
	ID			uint		`gorm:"primaryKey"`
	Name		string
	RateDate 	int64
	Open		float32
	High		float32
	Low			float32
	Close		float32
}
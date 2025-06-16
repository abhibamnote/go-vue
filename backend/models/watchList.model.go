package models

import "gorm.io/gorm"

type WatchList struct {
	gorm.Model
	UserId   int    `gorm:"uniqueIndex:idx_userId_masterId" json:"userId"`
	MasterId int    `gorm:"uniqueIndex:idx_userId_masterId" json:"masterId"`
	Master   Master `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
} 

type PostWatchListSchema struct {
	UserId   int `json:"userId"`
	MasterId int `json:"masterId"`
}

package models

import "time"

type ConversionType string

const (
	TypeASCII    ConversionType = "ascii"
	TypePixelArt ConversionType = "pixelart"
)

type ConversionRecord struct {
	ID        string         `bson:"_id,omitempty" json:"id"`
	FileName  string         `bson:"filename" json:"filename"`
	Type      ConversionType `bson:"type" json:"type"`
	Result    string         `bson:"result" json:"result"`
	CreatedAt time.Time      `bson:"created_at" json:"createdAt"`
}

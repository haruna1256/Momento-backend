package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Photo struct {
    PhotoID   uuid.UUID `gorm:"type:char(36);primaryKey" json:"photo_id"`
    AlbumID   string    `gorm:"type:varchar(255);not null" json:"album_id"`
    ImageURL  string    `gorm:"type:text;not null" json:"image_url"`
    Caption   string    `gorm:"type:text" json:"caption"`
    Place     string    `gorm:"type:text" json:"place"`
    Latitude  float64   `gorm:"type:double;default:0" json:"latitude"`
    Longitude float64   `gorm:"type:double;default:0" json:"longitude"`
    TakenAt   time.Time `gorm:"type:datetime" json:"taken_at"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// GORMのBeforeCreateフックでUUID自動生成
func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
    if p.PhotoID == uuid.Nil {
        p.PhotoID = uuid.New()
    }
    return
}

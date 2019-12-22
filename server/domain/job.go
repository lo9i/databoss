package domain

import (
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

type Job struct {
	Id        uint64     `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name      string     `json:"name,omitempty" gorm:"size:255;not null"`
	From      date.Date  `json:"from,omitempty"`
	To        *date.Date `json:"to,omitempty"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

package domain

import (
	"time"
)

type Candidate struct {
	Id          uint64 `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	UserId      string `gorm:"type:varchar(20);not null" json:"userId,omitempty"`
	FirstName   string `gorm:"size:255;not null" json:"firstName,omitempty"`
	LastName    string `gorm:"size:255;not null" json:"lastName,omitempty"`
	Male        bool   `gorm:"not null" json:"male"`
	PhoneNumber string `gorm:"size:255;null" json:"phoneNumber,omitempty"`

	Street     string `gorm:"size:255;null" json:"street,omitempty"`
	Number     int    `gorm:"null" json:"number,omitempty"`
	Floor      int    `gorm:"null" json:"floor,omitempty"`
	Department string `gorm:"size:10;null" json:"department,omitempty"`
	City       string `gorm:"size:255;null" json:"city,omitempty"`
	ZipCode    string `gorm:"size:10;null" json:"zipCode,omitempty"`
	State      string `gorm:"size:255;null" json:"state,omitempty"`

	Email string `gorm:"size:255;null" json:"state,omitempty"`

	Jobs            []*Job    `gorm:"foreignkey:CandidateId" json:"jobs,omitempty"`
	BirthDate       time.Time `gorm:"not null" json:"birthDate"`
	LastRetrievedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"lastRetrievedAt"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CandidateService interface {

	/*
	 * Get user
	 */
	Get(id uint64) *Candidate

	GetByUserId(userId string) (*Candidate, error)

	/*
	 * Get list of users by some criteria
	 */
	Find(where ...interface{}) *[]Candidate
}

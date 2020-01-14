package domain

import (
	"time"
)

type Job struct {
	Id          uint64    `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	CandidateId uint64    `json:"candidate_id,omitempty"`
	Name        string    `gorm:"size:255;not null" json:"name,omitempty"`
	From        Date      `json:"from,omitempty"`
	To          *Date     `json:"to,omitempty"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type JobService interface {

	/*
	 * Create user
	 */
	Save(job *Job) *Job

	/*
	 * Get user
	 */
	Get(id uint64) *Job

	FindByUserId(id uint64) *[]Job
	/*
	 * Get list of users by some criteria
	 */
	Find(where ...interface{}) *[]Job
}

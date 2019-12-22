package domain

type Candidate struct {
	Id        uint64 `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	FirstName string `json:"firstName,omitempty" gorm:"size:255;not null"`
	LastName  string `json:"lastName,omitempty" gorm:"size:255;not null"`
	Jobs      []*Job `json:"jobs,omitempty" gorm:"foreignkey:CandidateId;association_foreignkey:Id"`
}

type CandidateService interface {

	/*
	 * Create user
	 */
	Save(user *Candidate) (*Candidate, error)

	/*
	 * Get user
	 */
	Get(id uint64) *Candidate

	/*
	 * Get list of users by some criteria
	 */
	Find(where ...interface{}) *[]Candidate
}

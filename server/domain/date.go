package domain

type Date struct {
	// Year of date. Must be from 1 to 9999, or 0 if specifying a date without
	// a year.
	Year int32 `json:"year,omitempty"`
	// Month of year. Must be from 1 to 12, or 0 if specifying a year without a
	// month and day.
	Month int32 `json:"month,omitempty"`
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0
	// if specifying a year by itself or a year and month where the day is not
	// significant.
	Day int32 `json:"day,omitempty"`
}

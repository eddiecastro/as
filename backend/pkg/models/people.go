package models

type People struct {
	ID                    int32  `json:"id"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	DisplayName           string `json:"display_name"`
	EmailAddress          string `json:"email_address"`
	SecondaryEmailAddress string `json:"secondary_email_address"`
	PersonalEmailAddress  string `json:"personal_email_address"`
	Title                 string `json:"title"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	Country               string `json:"country"`
}

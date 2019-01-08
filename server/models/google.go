package models

// GoogleProfile is our google profile model
type GoogleProfile struct {
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Email     string `json:"email"`
}

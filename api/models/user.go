package models

// User represents a user in the system
type User struct {
    ID    interface{} `json:"id"` 
    Name  string `json:"name"`
    Email string `json:"email"`
		BirthDate string `json:"birthDate"` 
}
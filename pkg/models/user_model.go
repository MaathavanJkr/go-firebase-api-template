// models/user.go
package models

type User struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Add other user properties as needed
}

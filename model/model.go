package model

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

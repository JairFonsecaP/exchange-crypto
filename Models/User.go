package models

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"` ///corregir
	// Email string `json:"email"`
}

func (u *User) Valid(userInput *User) bool {
	return u.Username == userInput.Username && u.Password == userInput.Password
}

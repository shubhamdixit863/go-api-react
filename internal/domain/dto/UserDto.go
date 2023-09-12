package dto

type UserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Location  string `json:"location"`
	Schedule  string `json:"schedule"`
	Password  string `json:"password"`
	Degree    string `json:"degree"`
}

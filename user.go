package todoapp

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

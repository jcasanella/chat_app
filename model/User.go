package model

type User struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

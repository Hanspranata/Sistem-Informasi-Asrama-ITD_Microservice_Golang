package models

type Admin struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Role struct {
	ID   int
	Name string
}

type Permission struct {
	ID          int
	Name        string
	Description string
}

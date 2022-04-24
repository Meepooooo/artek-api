package data

type State struct {
	Sessions []Session
}

type Session struct {
	ID    int
	Users []User
}

type User struct {
	ID   int
	Name string
	Role Role
}

type Role int

const (
	Member Role = iota
	Owner
)

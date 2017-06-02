package common

// Badge is actual raw markdown badge
type Badge struct {
	Name     string
	Markdown string
}

// Badge provides conventions for checking, generating badges
type BadgeContract struct {
	Name        string
	Description string
}

type Badger interface {
	Badges() ([]Badge, error)
}

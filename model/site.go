package model

type Site struct {
	Id string
	Active bool
	User User
	Category []byte
}
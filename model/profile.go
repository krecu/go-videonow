package model

type Profile struct {
	Id string
	Site Site
	Active bool
	Test bool
	Bad bool
	Category []byte
}
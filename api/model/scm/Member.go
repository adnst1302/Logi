package scm

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	UserId   string
	Email    string
	Role     string
	Password string
}

type MemberProfile struct {
	gorm.Model
	UserId         string
	FullName       string
	PlaceOfBirth   string
	BirthDate      string
	PhoneNumber    string
	Address        string
	EducationLevel string
	Division       string
	SignInDate     string
	StatusEmployee string
}

package models

import (
	"regexp"
	"unicode"
)
type User struct {
  Id       	int64
  Username 	string	`sql:"size:255; not null; unique;"`
  Email		string 	`sql:"size:255; not null; unique;"` 
  FirstName	string	`sql:"size:30; not null;"`
  Surname 	string 	`sql:"size:30; not null;"`   
  Password 	string 	`sql:"size:255; not null;"`
  Role 		string 	`sql:"size:30"` 
  IsEnabled	bool 	`sql:"default:true"`
  	ConfirmPassword 	string 	`sql:"-"`
  	Message				string 	`sql:"-"`
}

func (u *User) ValidateRegister() {


		errorstring := "";


		if VerifyName(u.Username) {
			errorstring = errorstring + "Username is invalid. A-Za-z0-9 only. -- "
		}
		if VerifyName(u.FirstName) {
			errorstring = errorstring + "First Name is invalid. A-Za-z0-9 only. -- "
		}
		if VerifyName(u.Surname) {
			errorstring = errorstring + "Last Name is invalid. A-Za-z0-9 only. -- "
		}
	


		if !VerifyEmail(u.Email) {
			errorstring = errorstring + "Email is invalid. --"
		}

		sevenOrMore, upper, lower, special := VerifyPassword(u.Password)
		if (!(sevenOrMore && upper && lower && special)) {
			errorstring = errorstring + "Password must be at least 7 characters, one lower case, one upper case, one digit and one special character --"
		}

		if (u.Password != u.ConfirmPassword) {
			errorstring = errorstring + "Passwords don't match. -- "
		}

		u.Message = errorstring

}

func (u *User) ValidateLogin() {
	errorstring := "";

	if VerifyName(u.Username) {
		errorstring = errorstring + "Username is invalid. A-Za-z0-9 only. -- "
	}

	sevenOrMore, upper, lower, special := VerifyPassword(u.Password)
	if (!(sevenOrMore && upper && lower && special)) {
		errorstring = errorstring + "Password must be at least 7 characters, one lower case, one upper case, one digit and one special character --"
	}

	u.Message = errorstring


}

func (u *User) ValidateNewPassword() {
	errorstring := "";

	sevenOrMore, upper, lower, special := VerifyPassword(u.Password)
	if (!(sevenOrMore && upper && lower && special)) {
		errorstring = errorstring + "Password must be at least 7 characters, one lower case, one upper case, one digit and one special character --"
	}

	if (u.Password != u.ConfirmPassword) {
		errorstring = errorstring + "Passwords don't match. -- "
	}
	
	u.Message = errorstring


}

//returns true if invalid
func VerifyName(s string) bool {
	reg, err := regexp.Compile(`\W`)

	if err != nil {
		panic(err)
	}

	return reg.MatchString(string(s))

}

func VerifyEmail(s string) bool {

	reg , err := regexp.Compile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`)

	if err != nil {
		panic(err)
	}

	return reg.MatchString(string(s))
}

func VerifyPassword(s string) (sevenOrMore, number, upper, special bool) {
    letters := 0
    for _, s := range s {
        switch {
        case unicode.IsNumber(s):
            number = true
        case unicode.IsUpper(s):
            upper = true
            letters++
        case unicode.IsPunct(s) || unicode.IsSymbol(s):
            special = true
        case unicode.IsLetter(s) || s == ' ':
            letters++
        default:
            //return false, false, false, false
        }
    }
    sevenOrMore = letters >= 7
    return
}
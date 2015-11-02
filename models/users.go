package models

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
}

